#!/bin/bash

# 简单的ES查询命令
# 请根据实际ES地址修改 ES_HOST

ES_HOST="http://localhost:9200"
ES_INDEX="addons_customer_pro_clues"

echo "=== ES查询命令 ==="
echo ""

# 方法1: 使用JSON文件
echo "方法1: 使用JSON文件"
echo "curl -X GET \"$ES_HOST/$ES_INDEX/_count\" -H \"Content-Type: application/json\" -d @es_count_query.json"
echo ""

# 方法2: 直接使用内联JSON
echo "方法2: 直接使用内联JSON"
echo "curl -X GET \"$ES_HOST/$ES_INDEX/_count\" -H \"Content-Type: application/json\" -d '"
echo '{
  "query": {
    "bool": {
      "must": [
        {"term": {"status": "0"}},
        {
          "bool": {
            "should": [
              {"term": {"services_ids": "85"}},
              {"wildcard": {"services_ids": "85,*"}},
              {"wildcard": {"services_ids": "*,85"}},
              {"wildcard": {"services_ids": "*,85,*"}}
            ],
            "minimum_should_match": 1
          }
        }
      ],
      "must_not": [
        {"exists": {"field": "deleted_at"}}
      ]
    }
  },
  "size": 0
}'
echo "'"
echo ""

# 方法3: 使用jq解析结果
echo "方法3: 使用jq解析结果（需要安装jq）"
echo "curl -s -X GET \"$ES_HOST/$ES_INDEX/_count\" -H \"Content-Type: application/json\" -d @es_count_query.json | jq '.count'"
echo ""

echo "=== 使用说明 ==="
echo "1. 确保ES服务正在运行"
echo "2. 根据实际ES地址修改 ES_HOST 变量"
echo "3. 确保索引 $ES_INDEX 存在"
echo "4. 运行命令后对比MySQL查询结果" 