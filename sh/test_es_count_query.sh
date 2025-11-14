#!/bin/bash

# ES查询脚本 - 对比MySQL数据总数
# 对应MySQL查询: SELECT COUNT(1) FROM `addons_customer_pro_clues` AS clues WHERE `clues`.`status`='0' AND (clues.services_ids REGEXP '^85$|^85,|,85$|,85,') AND `deleted_at` IS NULL

# ES配置
ES_HOST="http://localhost:9200"  # 根据实际ES地址修改
ES_INDEX="addons_customer_pro_clues"

echo "=== Elasticsearch 数据总数查询 ==="
echo "索引: $ES_INDEX"
echo "查询条件: status='0' AND services_ids包含'85' AND deleted_at为null"
echo ""

# 构建ES查询
QUERY='{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "status": "0"
          }
        },
        {
          "bool": {
            "should": [
              {
                "term": {
                  "services_ids": "85"
                }
              },
              {
                "wildcard": {
                  "services_ids": "85,*"
                }
              },
              {
                "wildcard": {
                  "services_ids": "*,85"
                }
              },
              {
                "wildcard": {
                  "services_ids": "*,85,*"
                }
              }
            ],
            "minimum_should_match": 1
          }
        }
      ],
      "must_not": [
        {
          "exists": {
            "field": "deleted_at"
          }
        }
      ]
    }
  },
  "size": 0
}'

echo "执行查询..."
echo ""

# 执行ES查询
RESULT=$(curl -s -X GET "$ES_HOST/$ES_INDEX/_count" \
  -H "Content-Type: application/json" \
  -d "$QUERY")

# 检查查询是否成功
if [ $? -eq 0 ]; then
    # 提取总数
    COUNT=$(echo "$RESULT" | grep -o '"count":[0-9]*' | cut -d':' -f2)
    
    if [ -n "$COUNT" ]; then
        echo "✅ ES查询成功"
        echo "📊 ES数据总数: $COUNT"
        echo ""
        echo "🔍 查询详情:"
        echo "$RESULT" | python3 -m json.tool 2>/dev/null || echo "$RESULT"
    else
        echo "❌ 无法解析ES响应中的count字段"
        echo "响应内容:"
        echo "$RESULT"
    fi
else
    echo "❌ ES查询失败"
    echo "请检查ES服务是否正常运行，以及索引是否存在"
fi

echo ""
echo "=== 对比说明 ==="
echo "MySQL查询: SELECT COUNT(1) FROM addons_customer_pro_clues WHERE status='0' AND (services_ids REGEXP '^85$|^85,|,85$|,85,') AND deleted_at IS NULL"
echo ""
echo "如果ES总数与MySQL总数不一致，可能的原因："
echo "1. ES数据同步不完整"
echo "2. 数据格式转换问题"
echo "3. 字段映射问题"
echo "4. 软删除字段处理不一致" 