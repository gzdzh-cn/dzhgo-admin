#!/bin/bash

# Docker部署Canal Adapter脚本

echo "=== Docker部署Canal Adapter ==="
echo ""

# 检查Docker和Docker Compose
echo "🔍 检查环境..."
if ! command -v docker &> /dev/null; then
    echo "❌ Docker未安装，请先安装Docker"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose未安装，请先安装Docker Compose"
    exit 1
fi

echo "✅ Docker环境检查通过"
echo ""

# 创建目录结构
echo "📁 创建目录结构..."
mkdir -p docker/canal-server/conf/example
mkdir -p docker/canal-server/logs
mkdir -p docker/canal-adapter/conf/es7
mkdir -p docker/canal-adapter/logs
mkdir -p docker/mysql/init
mkdir -p docker/mysql/conf.d
mkdir -p docker/canal-admin/conf
mkdir -p docker/canal-admin/logs

echo "✅ 目录结构创建完成"
echo ""

# 启动服务
echo "🚀 启动Docker服务..."
docker-compose up -d

echo "✅ 服务启动完成"
echo ""

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 30

# 检查服务状态
echo "🔍 检查服务状态..."
echo ""

echo "1. 检查MySQL服务..."
if docker-compose ps mysql | grep -q "Up"; then
    echo "✅ MySQL服务运行正常"
else
    echo "❌ MySQL服务启动失败"
fi

echo ""

echo "2. 检查Elasticsearch服务..."
if docker-compose ps elasticsearch | grep -q "Up"; then
    echo "✅ Elasticsearch服务运行正常"
else
    echo "❌ Elasticsearch服务启动失败"
fi

echo ""

echo "3. 检查Canal Server服务..."
if docker-compose ps canal-server | grep -q "Up"; then
    echo "✅ Canal Server服务运行正常"
else
    echo "❌ Canal Server服务启动失败"
fi

echo ""

echo "4. 检查Canal Adapter服务..."
if docker-compose ps canal-adapter | grep -q "Up"; then
    echo "✅ Canal Adapter服务运行正常"
else
    echo "❌ Canal Adapter服务启动失败"
fi

echo ""

# 验证连接
echo "🔍 验证服务连接..."
echo ""

echo "1. 测试MySQL连接..."
docker-compose exec mysql mysql -u canal -pcanal -e "SELECT COUNT(*) FROM dzhgo.addons_customer_pro_clues;" 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ MySQL连接正常"
else
    echo "❌ MySQL连接失败"
fi

echo ""

echo "2. 测试Elasticsearch连接..."
curl -s http://localhost:9200/_cluster/health > /dev/null
if [ $? -eq 0 ]; then
    echo "✅ Elasticsearch连接正常"
else
    echo "❌ Elasticsearch连接失败"
fi

echo ""

echo "3. 测试Canal Server连接..."
curl -s http://localhost:11110/actuator/health > /dev/null
if [ $? -eq 0 ]; then
    echo "✅ Canal Server连接正常"
else
    echo "❌ Canal Server连接失败"
fi

echo ""

echo "4. 测试Canal Adapter连接..."
curl -s http://localhost:8081/actuator/health > /dev/null
if [ $? -eq 0 ]; then
    echo "✅ Canal Adapter连接正常"
else
    echo "❌ Canal Adapter连接失败"
fi

echo ""

# 显示服务信息
echo "📊 服务信息:"
echo ""
echo "MySQL:"
echo "  - 地址: localhost:3306"
echo "  - 用户名: canal"
echo "  - 密码: canal"
echo "  - 数据库: dzhgo"
echo ""

echo "Elasticsearch:"
echo "  - 地址: localhost:9200"
echo "  - 集群: elasticsearch"
echo ""

echo "Canal Server:"
echo "  - 管理端口: localhost:11110"
echo "  - 数据端口: localhost:11111"
echo ""

echo "Canal Adapter:"
echo "  - 管理端口: localhost:8081"
echo ""

echo "Canal Admin (可选):"
echo "  - 管理界面: http://localhost:8089"
echo "  - 用户名: admin"
echo "  - 密码: admin"
echo ""

# 验证数据同步
echo "🔍 验证数据同步..."
echo ""

echo "1. 检查ES索引..."
curl -s http://localhost:9200/_cat/indices | grep addons_customer_pro_clues
if [ $? -eq 0 ]; then
    echo "✅ ES索引创建成功"
else
    echo "❌ ES索引创建失败"
fi

echo ""

echo "2. 检查同步数据..."
curl -s "http://localhost:9200/addons_customer_pro_clues/_count" | jq '.count' 2>/dev/null
if [ $? -eq 0 ]; then
    COUNT=$(curl -s "http://localhost:9200/addons_customer_pro_clues/_count" | jq -r '.count')
    echo "✅ 同步数据量: $COUNT 条"
else
    echo "❌ 数据同步失败"
fi

echo ""

echo "=== 部署完成 ==="
echo ""
echo "🎉 Canal Adapter Docker部署成功！"
echo ""
echo "📝 常用命令:"
echo "  查看服务状态: docker-compose ps"
echo "  查看服务日志: docker-compose logs -f"
echo "  停止服务: docker-compose down"
echo "  重启服务: docker-compose restart"
echo "  清理数据: docker-compose down -v"
echo ""
echo "🔍 监控命令:"
echo "  查看Canal Adapter日志: docker-compose logs -f canal-adapter"
echo "  查看Canal Server日志: docker-compose logs -f canal-server"
echo "  查看MySQL日志: docker-compose logs -f mysql"
echo "  查看ES日志: docker-compose logs -f elasticsearch" 