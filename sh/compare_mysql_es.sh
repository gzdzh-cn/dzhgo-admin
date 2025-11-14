#!/bin/bash

# MySQL和ES数据对比脚本
# 对应MySQL查询: SELECT COUNT(1) FROM addons_customer_pro_clues WHERE status='0' AND (services_ids REGEXP '^85$|^85,|,85$|,85,') AND deleted_at IS NULL

echo "=== MySQL vs Elasticsearch 数据对比 ==="
echo ""

# ES配置
ES_HOST="http://localhost:9200"
ES_INDEX="addons_customer_pro_clues"

# 获取ES总数
echo "🔍 查询ES数据..."
ES_COUNT=$(curl -s -X GET "$ES_HOST/$ES_INDEX/_count" -H "Content-Type: application/json" -d @es_count_query.json | jq -r '.count')

if [ "$ES_COUNT" = "null" ] || [ -z "$ES_COUNT" ]; then
    echo "❌ ES查询失败"
    exit 1
fi

echo "✅ ES查询成功: $ES_COUNT 条记录"
echo ""

# 显示查询条件
echo "📋 查询条件:"
echo "  - status = 0"
echo "  - services_ids 包含 '85' (支持: 85, 85,*, *,85, *,85,*)"
echo "  - deleted_at 为 null"
echo ""

# 显示ES查询详情
echo "🔍 ES查询详情:"
cat es_count_query.json | jq '.'
echo ""

# 显示索引信息
echo "📊 索引信息:"
TOTAL_COUNT=$(curl -s -X GET "$ES_HOST/$ES_INDEX/_count" | jq -r '.count')
echo "  总记录数: $TOTAL_COUNT"
echo "  符合条件记录数: $ES_COUNT"
echo "  占比: $(echo "scale=2; $ES_COUNT * 100 / $TOTAL_COUNT" | bc)%"
echo ""

echo "=== 对比结果 ==="
echo "ES数据总数: $ES_COUNT"
echo ""
echo "请运行对应的MySQL查询进行对比:"
echo "SELECT COUNT(1) FROM addons_customer_pro_clues WHERE status='0' AND (services_ids REGEXP '^85$|^85,|,85$|,85,') AND deleted_at IS NULL;"
echo ""

# 如果安装了mysql客户端，可以自动执行MySQL查询
if command -v mysql &> /dev/null; then
    echo "💡 提示: 如果您配置了MySQL连接，可以运行以下命令自动对比:"
    echo "mysql -u用户名 -p密码 -h主机 数据库名 -e \"SELECT COUNT(1) FROM addons_customer_pro_clues WHERE status='0' AND (services_ids REGEXP '^85$|^85,|,85$|,85,') AND deleted_at IS NULL;\""
fi 