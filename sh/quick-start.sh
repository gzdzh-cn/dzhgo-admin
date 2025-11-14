#!/bin/bash

# Canal + Kafka + ES 快速启动脚本
echo "=========================================="
echo "        Canal + Kafka + ES 快速启动"
echo "=========================================="

# 检查Docker状态
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 检查jq是否安装
if ! command -v jq &> /dev/null; then
    echo "⚠️  jq未安装，将使用curl替代"
    JQ_AVAILABLE=false
else
    JQ_AVAILABLE=true
fi

# 停止现有服务
echo "停止现有服务..."
docker-compose down

# 启动服务
echo "启动服务..."
docker-compose up -d

# 等待服务启动
echo "等待服务启动..."
sleep 30

# 检查服务状态
echo "检查服务状态..."
docker-compose ps

# 等待ES启动
echo "等待Elasticsearch启动..."
for i in {1..60}; do
    if curl -s http://localhost:9200/_cluster/health > /dev/null 2>&1; then
        echo "✅ Elasticsearch已启动"
        break
    fi
    if [ $i -eq 60 ]; then
        echo "❌ Elasticsearch启动超时"
        exit 1
    fi
    echo "等待Elasticsearch启动... ($i/60)"
    sleep 5
done

# 初始化ES索引
echo "初始化ES索引..."
docker exec canal-adapter /home/admin/canal-adapter/conf/es7/init_es_index.sh

# 检查索引创建结果
echo "检查ES索引..."
if [ "$JQ_AVAILABLE" = true ]; then
    curl -s "http://localhost:9200/_cat/indices?v" | jq .
else
    curl -s "http://localhost:9200/_cat/indices?v"
fi

# 等待Canal Admin启动
echo "等待Canal Admin启动..."
for i in {1..30}; do
    if curl -s http://localhost:8089/api/v1/canal/config/server/list > /dev/null 2>&1; then
        echo "✅ Canal Admin已启动"
        break
    fi
    if [ $i -eq 30 ]; then
        echo "⚠️  Canal Admin启动可能较慢，请稍后手动检查"
    fi
    echo "等待Canal Admin启动... ($i/30)"
    sleep 2
done

echo "=========================================="
echo "🎉 启动完成！"
echo "=========================================="
echo "📋 服务访问地址："
echo "  • Canal Admin: http://localhost:8089"
echo "    用户名: admin, 密码: 123456"
echo "  • Elasticsearch: http://localhost:9200"
echo "  • Kibana: http://localhost:5601"
echo "  • Kafka: localhost:9092"
echo "=========================================="
echo "📝 下一步操作："
echo "1. 访问Canal Admin配置数据源和实例"
echo "2. 启动Canal Adapter同步任务"
echo "3. 运行测试脚本: ./test-canal-kafka-sync.sh"
echo "==========================================" 