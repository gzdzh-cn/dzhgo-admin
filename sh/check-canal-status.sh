#!/bin/bash

# Canal系统状态检查脚本
echo "=========================================="
echo "        Canal系统状态检查"
echo "=========================================="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 检查函数
check_service() {
    local service_name=$1
    local container_name=$2
    
    echo -e "${BLUE}检查 $service_name 状态...${NC}"
    if docker ps --format "table {{.Names}}\t{{.Status}}" | grep -q "$container_name.*Up"; then
        echo -e "${GREEN}✅ $service_name 运行正常${NC}"
        return 0
    else
        echo -e "${RED}❌ $service_name 未运行或异常${NC}"
        return 1
    fi
}

check_port() {
    local port=$1
    local service_name=$2
    
    echo -e "${BLUE}检查端口 $port ($service_name)...${NC}"
    if lsof -i :$port > /dev/null 2>&1; then
        echo -e "${GREEN}✅ 端口 $port 正常监听${NC}"
        return 0
    else
        echo -e "${RED}❌ 端口 $port 未监听${NC}"
        return 1
    fi
}

check_es_health() {
    echo -e "${BLUE}检查Elasticsearch健康状态...${NC}"
    local health=$(curl -s http://localhost:9200/_cluster/health 2>/dev/null)
    if [ $? -eq 0 ]; then
        local status=$(echo "$health" | grep -o '"status":"[^"]*"' | cut -d'"' -f4)
        echo -e "${GREEN}✅ Elasticsearch状态: $status${NC}"
        return 0
    else
        echo -e "${RED}❌ Elasticsearch连接失败${NC}"
        return 1
    fi
}

check_es_index() {
    echo -e "${BLUE}检查ES索引...${NC}"
    local index_count=$(curl -s "http://localhost:9200/_cat/indices/addons_customer_pro_clues?v" 2>/dev/null | wc -l)
    if [ "$index_count" -gt 1 ]; then
        echo -e "${GREEN}✅ addons_customer_pro_clues 索引存在${NC}"
        return 0
    else
        echo -e "${RED}❌ addons_customer_pro_clues 索引不存在${NC}"
        return 1
    fi
}

check_mysql_binlog() {
    echo -e "${BLUE}检查MySQL Binlog状态...${NC}"
    local binlog_status=$(docker exec mysql-canal mysql -u canal -pcanal123 -e "SHOW MASTER STATUS\G" 2>/dev/null)
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ MySQL Binlog正常${NC}"
        return 0
    else
        echo -e "${RED}❌ MySQL Binlog异常${NC}"
        return 1
    fi
}

check_canal_admin() {
    echo -e "${BLUE}检查Canal Admin...${NC}"
    if curl -s http://localhost:8089/api/v1/canal/config/server/list > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Canal Admin正常${NC}"
        return 0
    else
        echo -e "${RED}❌ Canal Admin异常${NC}"
        return 1
    fi
}



# 主检查流程
echo "开始系统状态检查..."
echo ""

# 检查Docker服务
check_service "MySQL" "mysql-canal"
check_service "Canal Server" "canal-server"
check_service "Canal Admin" "canal-admin"
check_service "Canal Adapter" "canal-adapter"
check_service "Elasticsearch" "elasticsearch"
check_service "Kibana" "kibana"

echo ""

# 检查端口
check_port "3306" "MySQL"
check_port "11110" "Canal Server"
check_port "8089" "Canal Admin"
check_port "8081" "Canal Adapter"
check_port "9200" "Elasticsearch"
check_port "5601" "Kibana"

echo ""

# 检查服务健康状态
check_es_health
check_es_index
check_mysql_binlog
check_canal_admin

echo ""
echo "=========================================="
echo "状态检查完成！"
echo "=========================================="

# 显示访问地址
echo -e "${BLUE}📋 服务访问地址：${NC}"
echo "  • Canal Admin: http://localhost:8089 (admin/123456)"
echo "  • Elasticsearch: http://localhost:9200"
echo "  • Kibana: http://localhost:5601"
echo ""

# 显示常用命令
echo -e "${BLUE}📝 常用命令：${NC}"
echo "  • 查看日志: docker-compose logs -f [service_name]"
echo "  • 重启服务: docker-compose restart [service_name]"
echo "  • 停止服务: docker-compose down"
echo "  • 启动服务: docker-compose up -d"
echo ""

# 显示测试命令
echo -e "${BLUE}🧪 测试命令：${NC}"
echo "  • 运行测试: ./test-canal-kafka-sync.sh"
echo "  • 检查ES数据: curl http://localhost:9200/addons_customer_pro_clues/_search?pretty"
echo "  • 检查MySQL数据: docker exec mysql-canal mysql -u canal -pcanal123 -e 'SELECT COUNT(*) FROM dzhgo_go.addons_customer_pro_clues;'" 