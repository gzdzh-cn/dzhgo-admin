#!/bin/bash

# Canal + Kafka + ES 部署脚本
echo "=========================================="
echo "        Canal + Kafka + ES 部署脚本"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 停止并删除现有容器
echo "停止现有容器..."
docker-compose down

# 启动服务
echo "启动Canal + Kafka + ES服务..."
docker-compose up -d

# 等待服务启动
echo "等待服务启动..."
sleep 30

# 检查服务状态
echo "检查服务状态..."
docker-compose ps

# 等待ES服务完全启动
echo "等待Elasticsearch服务启动..."
for i in {1..60}; do
    if curl -s http://localhost:9200/_cluster/health > /dev/null 2>&1; then
        echo "✅ Elasticsearch服务已启动"
        break
    fi
    echo "等待Elasticsearch启动... ($i/60)"
    sleep 5
done

# 初始化ES索引
echo "初始化ES索引..."
docker exec canal-adapter /home/admin/canal-adapter/conf/es7/init_es_index.sh

# 检查Canal Admin状态
echo "检查Canal Admin状态..."
for i in {1..30}; do
    if curl -s http://localhost:8089/api/v1/canal/config/server/list > /dev/null 2>&1; then
        echo "✅ Canal Admin服务已启动"
        break
    fi
    echo "等待Canal Admin启动... ($i/30)"
    sleep 2
done

echo "=========================================="
echo "部署完成！"
echo "=========================================="
echo "服务访问地址："
echo "  - Canal Admin: http://localhost:8089 (admin/123456)"
echo "  - Elasticsearch: http://localhost:9200"
echo "  - Kibana: http://localhost:5601"
echo "  - Kafka: localhost:9092"
echo "=========================================="
echo "下一步操作："
echo "1. 访问Canal Admin配置数据源和实例"
echo "2. 启动Canal Adapter同步任务"
echo "3. 测试数据同步"
echo "==========================================" 