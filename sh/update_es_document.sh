#!/bin/bash

# ES文档更新脚本
# 更新ID为1948747010836271105的文档

# ES配置
ES_HOST="http://localhost:9200"
ES_INDEX="addons_customer_pro_clues"
DOC_ID="1948747010836271105"

echo "=== ES文档更新 ==="
echo "索引: $ES_INDEX"
echo "文档ID: $DOC_ID"
echo "更新字段: services_id=84, services_ids=84,85"
echo ""

# 方法1: 使用POST请求更新
echo "🔧 方法1: 使用POST请求更新"
echo "curl -X POST \"$ES_HOST/$ES_INDEX/_update/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d @es_update_command.json"
echo ""

# 执行更新
RESULT=$(curl -s -X POST "$ES_HOST/$ES_INDEX/_update/$DOC_ID" \
  -H "Content-Type: application/json" \
  -d @es_update_command.json)

# 检查更新结果
if [ $? -eq 0 ]; then
    echo "✅ 更新请求成功"
    echo "响应内容:"
    echo "$RESULT" | jq '.' 2>/dev/null || echo "$RESULT"
    
    # 检查更新是否成功
    if echo "$RESULT" | grep -q '"result":"updated"'; then
        echo ""
        echo "🎉 文档更新成功！"
    else
        echo ""
        echo "⚠️  更新可能未生效，请检查响应"
    fi
else
    echo "❌ 更新请求失败"
    echo "错误信息:"
    echo "$RESULT"
fi

echo ""
echo "=== 验证更新结果 ==="

# 查询更新后的文档
echo "🔍 查询更新后的文档..."
QUERY_RESULT=$(curl -s -X GET "$ES_HOST/$ES_INDEX/_doc/$DOC_ID")

if [ $? -eq 0 ]; then
    echo "✅ 查询成功"
    echo "更新后的文档内容:"
    echo "$QUERY_RESULT" | jq '._source | {id, name, services_id, services_ids, status}' 2>/dev/null || echo "$QUERY_RESULT"
else
    echo "❌ 查询失败"
    echo "$QUERY_RESULT"
fi

echo ""
echo "=== 其他更新方法 ==="

# 方法2: 使用PUT请求（完整替换）
echo "🔧 方法2: 使用PUT请求完整替换文档"
echo "curl -X PUT \"$ES_HOST/$ES_INDEX/_doc/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"services_id\":\"84\",\"services_ids\":\"84,85\",...其他字段}'"
echo ""

# 方法3: 使用脚本更新
echo "🔧 方法3: 使用脚本更新（条件更新）"
echo "curl -X POST \"$ES_HOST/$ES_INDEX/_update/$DOC_ID\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"script\":{\"source\":\"ctx._source.services_id=params.sid;ctx._source.services_ids=params.sids\",\"lang\":\"painless\",\"params\":{\"sid\":\"84\",\"sids\":\"84,85\"}}}'"
echo ""

echo "=== 使用说明 ==="
echo "1. 确保ES服务正在运行"
echo "2. 确保文档ID存在"
echo "3. 更新后会自动刷新索引"
echo "4. 可以使用查询命令验证更新结果" 