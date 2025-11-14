#!/bin/bash

# Canal + Kafka 本地构建部署脚本
echo "=== Canal + Kafka 本地构建部署 ==="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 检查Docker和Docker Compose
check_dependencies() {
    echo -e "${BLUE}检查依赖...${NC}"
    
    if ! command -v docker &> /dev/null; then
        echo -e "${RED}Docker未安装，请先安装Docker${NC}"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        echo -e "${RED}Docker Compose未安装，请先安装Docker Compose${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}依赖检查通过${NC}"
}

# 创建必要的目录
create_directories() {
    echo -e "${BLUE}创建目录结构...${NC}"
    
    mkdir -p docker/mysql/{conf,data,init}
    mkdir -p docker/canal-server/{conf,logs}
    mkdir -p docker/canal-server/conf/example
    mkdir -p docker/canal-admin/{conf,logs}
    mkdir -p docker/canal-adapter/{conf,logs}
    mkdir -p docker/canal-adapter/conf/es7
    
    echo -e "${GREEN}目录创建完成${NC}"
}

# 构建Canal镜像
build_canal_images() {
    echo -e "${BLUE}构建Canal镜像...${NC}"
    
    # 构建Canal Server
    echo "构建Canal Server镜像..."
    docker build -t canal-server:latest ./docker/canal-server/
    
    # 构建Canal Admin
    echo "构建Canal Admin镜像..."
    docker build -t canal-admin:latest ./docker/canal-admin/
    
    # 构建Canal Adapter
    echo "构建Canal Adapter镜像..."
    docker build -t canal-adapter:latest ./docker/canal-adapter/
    
    echo -e "${GREEN}Canal镜像构建完成${NC}"
}

# 初始化数据库
init_database() {
    echo -e "${BLUE}初始化数据库...${NC}"
    
    # 创建初始化SQL脚本
    cat > docker/mysql/init/01-init.sql << 'EOF'
-- 创建数据库
CREATE DATABASE IF NOT EXISTS dzhgo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE dzhgo;

-- 创建线索表
CREATE TABLE IF NOT EXISTS addons_customer_pro_clues (
    id VARCHAR(50) PRIMARY KEY,
    serialId INT,
    guest_id VARCHAR(50),
    name VARCHAR(100),
    mobile VARCHAR(20),
    wechat VARCHAR(50),
    keywords TEXT,
    status INT DEFAULT 0,
    account_id VARCHAR(50),
    project_id VARCHAR(50),
    services_id VARCHAR(50),
    services_ids VARCHAR(200),
    source_from VARCHAR(50),
    followup_type VARCHAR(50),
    level VARCHAR(20),
    filter_remark TEXT,
    created_id VARCHAR(50),
    filter_group_ids VARCHAR(200),
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    INDEX idx_update_time (updateTime),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建属性表
CREATE TABLE IF NOT EXISTS addons_customer_pro_clues_attr (
    id VARCHAR(50) PRIMARY KEY,
    clues_id VARCHAR(50),
    guest_ip_info TEXT,
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (clues_id) REFERENCES addons_customer_pro_clues(id),
    INDEX idx_clues_id (clues_id),
    INDEX idx_update_time (updateTime)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建Canal用户
CREATE USER IF NOT EXISTS 'canal'@'%' IDENTIFIED BY 'canal123';
GRANT ALL PRIVILEGES ON dzhgo.* TO 'canal'@'%';
FLUSH PRIVILEGES;

-- 插入测试数据
INSERT INTO addons_customer_pro_clues (id, serialId, name, mobile, status, services_ids) VALUES
('test_001', 1, '测试用户1', '13800138001', 0, '85,86'),
('test_002', 2, '测试用户2', '13800138002', 0, '85'),
('test_003', 3, '测试用户3', '13800138003', 1, '86,87');

INSERT INTO addons_customer_pro_clues_attr (id, clues_id, guest_ip_info) VALUES
('attr_001', 'test_001', '192.168.1.100'),
('attr_002', 'test_002', '192.168.1.101'),
('attr_003', 'test_003', '192.168.1.102');
EOF
    
    echo -e "${GREEN}数据库初始化脚本创建完成${NC}"
}

# 启动服务
start_services() {
    echo -e "${BLUE}启动Canal + Kafka服务...${NC}"
    
    # 启动基础服务
    docker-compose -f docker-compose-canal-kafka-local.yml up -d mysql elasticsearch 
    
    echo -e "${YELLOW}等待基础服务启动...${NC}"
    sleep 30
    
    # 启动Canal服务
    docker-compose -f docker-compose-canal-kafka-local.yml up -d canal-server canal-admin
    
    echo -e "${YELLOW}等待Canal服务启动...${NC}"
    sleep 20
    
    # 启动Canal Adapter
    docker-compose -f docker-compose-canal-kafka-local.yml up -d canal-adapter
    
    echo -e "${GREEN}所有服务启动完成${NC}"
}

# 检查服务状态
check_services() {
    echo -e "${BLUE}检查服务状态...${NC}"
    
    # 检查MySQL
    echo "MySQL状态:"
    docker-compose -f docker-compose-canal-kafka-local.yml ps mysql
    

    
    # 检查Canal Server
    echo "Canal Server状态:"
    docker-compose -f docker-compose-canal-kafka-local.yml ps canal-server
    
    # 检查Canal Adapter
    echo "Canal Adapter状态:"
    docker-compose -f docker-compose-canal-kafka-local.yml ps canal-adapter
    
    # 检查Elasticsearch
    echo "Elasticsearch状态:"
    docker-compose -f docker-compose-canal-kafka-local.yml ps elasticsearch
}

# 测试数据同步
test_sync() {
    echo -e "${BLUE}测试数据同步...${NC}"
    
    # 等待服务完全启动
    sleep 10
    
    # 检查ES索引
    echo "检查ES索引:"
    curl -s "http://localhost:9200/_cat/indices?v"
    
    # 检查Canal Admin
    echo "Canal Admin访问地址: http://localhost:8089"
    echo "用户名: admin, 密码: admin"
    

}

# 主函数
main() {
    check_dependencies
    create_directories
    build_canal_images
    init_database
    start_services
    check_services
    test_sync
    
    echo -e "${GREEN}=== 部署完成 ===${NC}"
    echo -e "${YELLOW}访问地址:${NC}"
    echo "Canal Admin: http://localhost:8089"
    echo "Kibana: http://localhost:5601"
    echo "Elasticsearch: http://localhost:9200"
}

# 执行主函数
main 