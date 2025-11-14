#!/bin/bash

# ELK栈启动脚本
echo "=========================================="
echo "        ELK栈启动脚本"
echo "=========================================="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 检查Docker状态
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker未运行，请先启动Docker${NC}"
    exit 1
fi

# 创建必要的目录
echo -e "${BLUE}创建必要的目录...${NC}"
mkdir -p logs
mkdir -p docker/logstash/config docker/logstash/pipeline docker/logstash/logs
mkdir -p docker/filebeat

# 停止现有服务
echo -e "${BLUE}停止现有服务...${NC}"
docker-compose down

# 启动服务
echo -e "${BLUE}启动ELK栈服务...${NC}"
docker-compose up -d

# 等待服务启动
echo -e "${BLUE}等待服务启动...${NC}"
sleep 30

# 检查服务状态
echo -e "${BLUE}检查服务状态...${NC}"
docker-compose ps

# 等待Elasticsearch启动
echo -e "${BLUE}等待Elasticsearch启动...${NC}"
for i in {1..60}; do
    if curl -s http://localhost:9200/_cluster/health > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Elasticsearch已启动${NC}"
        break
    fi
    if [ $i -eq 60 ]; then
        echo -e "${RED}❌ Elasticsearch启动超时${NC}"
        exit 1
    fi
    echo "等待Elasticsearch启动... ($i/60)"
    sleep 5
done

# 等待Logstash启动
echo -e "${BLUE}等待Logstash启动...${NC}"
for i in {1..30}; do
    if curl -s http://localhost:9600 > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Logstash已启动${NC}"
        break
    fi
    if [ $i -eq 30 ]; then
        echo -e "${YELLOW}⚠️  Logstash启动可能较慢，请稍后手动检查${NC}"
    fi
    echo "等待Logstash启动... ($i/30)"
    sleep 2
done

# 检查ELK栈健康状态
echo -e "${BLUE}检查ELK栈健康状态...${NC}"

# 检查Elasticsearch
if curl -s http://localhost:9200/_cluster/health > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Elasticsearch健康${NC}"
else
    echo -e "${RED}❌ Elasticsearch异常${NC}"
fi

# 检查Logstash
if curl -s http://localhost:9600 > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Logstash健康${NC}"
else
    echo -e "${RED}❌ Logstash异常${NC}"
fi

# 检查Kibana
if curl -s http://localhost:5601 > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Kibana健康${NC}"
else
    echo -e "${RED}❌ Kibana异常${NC}"
fi

echo ""
echo "=========================================="
echo -e "${GREEN}🎉 ELK栈启动完成！${NC}"
echo "=========================================="
echo -e "${BLUE}📋 服务访问地址：${NC}"
echo "  • Kibana: http://localhost:5601"
echo "  • Elasticsearch: http://localhost:9200"
echo "  • Logstash API: http://localhost:9600"
echo "  • Filebeat: 收集日志中..."
echo ""
echo -e "${BLUE}📝 常用命令：${NC}"
echo "  • 查看日志: docker-compose logs -f [service_name]"
echo "  • 查看所有日志: docker-compose logs -f"
echo "  • 重启服务: docker-compose restart [service_name]"
echo "  • 停止服务: docker-compose down"
echo ""
echo -e "${BLUE}🔍 查看实时日志：${NC}"
echo "  • Logstash日志: docker-compose logs -f logstash"
echo "  • Filebeat日志: docker-compose logs -f filebeat"
echo "  • Elasticsearch日志: docker-compose logs -f elasticsearch"
echo ""
echo -e "${BLUE}📊 Kibana操作：${NC}"
echo "1. 访问 http://localhost:5601"
echo "2. 创建索引模式: logs-*"
echo "3. 创建仪表板查看日志"
echo "=========================================="
