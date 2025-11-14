#!/bin/bash

# Docker部署测试脚本

echo "=== Docker部署测试 ==="
echo ""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 测试函数
test_service() {
    local service_name=$1
    local test_url=$2
    local description=$3
    
    echo -e "${BLUE}测试 $description...${NC}"
    
    if curl -s "$test_url" > /dev/null 2>&1; then
        echo -e "${GREEN}✅ $description 正常${NC}"
        return 0
    else
        echo -e "${RED}❌ $description 失败${NC}"
        return 1
    fi
}

# 检查Docker服务状态
check_docker_service() {
    local service_name=$1
    local description=$2
    
    echo -e "${BLUE}检查 $description 状态...${NC}"
    
    if docker-compose ps | grep -q "$service_name.*Up"; then
        echo -e "${GREEN}✅ $description 运行正常${NC}"
        return 0
    else
        echo -e "${RED}❌ $description 未运行${NC}"
        return 1
    fi
}

# 测试MySQL连接
test_mysql() {
    echo -e "${BLUE}测试MySQL连接...${NC}"
    
    if docker-compose exec mysql mysql -u canal -pcanal -e "SELECT 1;" > /dev/null 2>&1; then
        echo -e "${GREEN}✅ MySQL连接正常${NC}"
        return 0
    else
        echo -e "${RED}❌ MySQL连接失败${NC}"
        return 1
    fi
}

# 测试ES连接
test_elasticsearch() {
    echo -e "${BLUE}测试Elasticsearch连接...${NC}"
    
    if curl -s "http://localhost:9200/_cluster/health" > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Elasticsearch连接正常${NC}"
        return 0
    else
        echo -e "${RED}❌ Elasticsearch连接失败${NC}"
        return 1
    fi
}

# 测试数据同步
test_data_sync() {
    echo -e "${BLUE}测试数据同步...${NC}"
    
    # 插入测试数据
    echo "插入测试数据..."
    docker-compose exec mysql mysql -u canal -pcanal dzhgo -e "
    INSERT INTO addons_customer_pro_clues (id, name, mobile, status, services_ids, create_time) 
    VALUES ('test_$(date +%s)', '测试用户$(date +%s)', '13800138000', 0, '85', NOW())
    ON DUPLICATE KEY UPDATE name = VALUES(name);" > /dev/null 2>&1
    
    # 等待同步
    echo "等待数据同步..."
    sleep 10
    
    # 检查ES中的数据
    local count=$(curl -s "http://localhost:9200/addons_customer_pro_clues/_count" | jq -r '.count' 2>/dev/null)
    
    if [ "$count" != "null" ] && [ "$count" -gt 0 ]; then
        echo -e "${GREEN}✅ 数据同步正常，ES中共有 $count 条数据${NC}"
        return 0
    else
        echo -e "${RED}❌ 数据同步失败${NC}"
        return 1
    fi
}

# 显示服务信息
show_service_info() {
    echo ""
    echo -e "${YELLOW}=== 服务信息 ===${NC}"
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
    echo "Canal Admin:"
    echo "  - 管理界面: http://localhost:8089"
    echo "  - 用户名: admin"
    echo "  - 密码: admin"
    echo ""
}

# 主测试流程
main() {
    echo "开始测试Docker部署..."
    echo ""
    
    local all_tests_passed=true
    
    # 1. 检查Docker服务状态
    echo -e "${YELLOW}=== 检查服务状态 ===${NC}"
    echo ""
    
    check_docker_service "mysql" "MySQL服务" || all_tests_passed=false
    check_docker_service "elasticsearch" "Elasticsearch服务" || all_tests_passed=false
    check_docker_service "canal-server" "Canal Server服务" || all_tests_passed=false
    check_docker_service "canal-adapter" "Canal Adapter服务" || all_tests_passed=false
    check_docker_service "canal-admin" "Canal Admin服务" || all_tests_passed=false
    
    echo ""
    
    # 2. 测试服务连接
    echo -e "${YELLOW}=== 测试服务连接 ===${NC}"
    echo ""
    
    test_mysql || all_tests_passed=false
    test_elasticsearch || all_tests_passed=false
    test_service "canal-server" "http://localhost:11110/actuator/health" "Canal Server健康检查" || all_tests_passed=false
    test_service "canal-adapter" "http://localhost:8081/actuator/health" "Canal Adapter健康检查" || all_tests_passed=false
    test_service "canal-admin" "http://localhost:8089" "Canal Admin管理界面" || all_tests_passed=false
    
    echo ""
    
    # 3. 测试数据同步
    echo -e "${YELLOW}=== 测试数据同步 ===${NC}"
    echo ""
    
    test_data_sync || all_tests_passed=false
    
    echo ""
    
    # 4. 显示结果
    if [ "$all_tests_passed" = true ]; then
        echo -e "${GREEN}🎉 所有测试通过！Docker部署成功！${NC}"
        echo ""
        show_service_info
    else
        echo -e "${RED}❌ 部分测试失败，请检查服务状态和日志${NC}"
        echo ""
        echo "查看日志命令:"
        echo "  docker-compose logs -f"
        echo "  docker-compose logs -f canal-adapter"
        echo "  docker-compose logs -f canal-server"
    fi
}

# 运行测试
main 