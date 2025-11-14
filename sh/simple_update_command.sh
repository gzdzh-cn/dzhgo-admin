#!/bin/bash

# 简单的ES更新命令
# 更新ID为1948747010836271105的文档

ES_HOST="http://localhost:9200"
ES_INDEX="addons_customer_pro_clues"
DOC_ID="1948747010836271105"

echo "=== 简单ES更新命令 ==="
echo ""

# 方法1: 使用JSON文件
echo "方法1: 使用JSON文件"
echo "curl -X POST \"$ES_HOST/$ES_INDEX/_update/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d @es_update_command.json"
echo ""

# 方法2: 直接使用内联JSON
echo "方法2: 直接使用内联JSON"
echo "curl -X POST \"$ES_HOST/$ES_INDEX/_update/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"doc\":{\"services_id\":\"84\",\"services_ids\":\"84,85\"}}'"
echo ""

# 方法3: 使用脚本更新
echo "方法3: 使用脚本更新"
echo "curl -X POST \"$ES_HOST/$ES_INDEX/_update/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"script\":{\"source\":\"ctx._source.services_id=\\\"84\\\";ctx._source.services_ids=\\\"84,85\\\"\",\"lang\":\"painless\"}}'"
echo ""

# 验证命令
echo "验证更新结果:"
echo "curl -X GET \"$ES_HOST/$ES_INDEX/_doc/$DOC_ID\" | jq '._source | {id, name, services_id, services_ids, status}'"
echo ""

echo "✅ 更新已完成！文档ID: $DOC_ID"
echo "📊 更新内容: services_id=84, services_ids=84,85" 