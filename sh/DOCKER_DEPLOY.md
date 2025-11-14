# Docker部署Canal Adapter指南

## 🚀 快速部署

### 1. 环境要求
- Docker 20.10+
- Docker Compose 2.0+
- 至少4GB内存

### 2. 一键部署
```bash
# 运行部署脚本
./docker-deploy.sh
```

### 3. 手动部署
```bash
# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f
```

## 📊 服务架构

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    MySQL    │    │ Canal Server│    │Canal Adapter│
│   (3306)    │◄──►│   (11111)   │◄──►│   (8081)    │
└─────────────┘    └─────────────┘    └─────────────┘
                           │                    │
                           ▼                    ▼
                   ┌─────────────┐    ┌─────────────┐
                   │Canal Admin  │    │Elasticsearch│
                   │   (8089)    │    │   (9200)    │
                   └─────────────┘    └─────────────┘
```

## 🔧 服务配置

### MySQL配置
- **端口**: 3306
- **用户名**: canal
- **密码**: canal
- **数据库**: dzhgo
- **字符集**: utf8mb4

### Elasticsearch配置
- **端口**: 9200 (HTTP), 9300 (Transport)
- **集群名**: elasticsearch
- **安全**: 禁用
- **内存**: 512MB

### Canal Server配置
- **管理端口**: 11110
- **数据端口**: 11111
- **监控端口**: 11112
- **模式**: TCP

### Canal Adapter配置
- **端口**: 8081
- **模式**: TCP
- **批量大小**: 1000
- **重试次数**: 0

## 📁 文件结构

```
.
├── docker-compose.yml              # Docker Compose配置
├── docker-deploy.sh               # 部署脚本
├── docker/
│   ├── canal-server/
│   │   └── conf/
│   │       ├── canal.conf         # Canal Server配置
│   │       └── example/
│   │           └── instance.properties
│   ├── canal-adapter/
│   │   └── conf/
│   │       ├── application.yml    # Adapter主配置
│   │       └── es7/
│   │           └── clues.yml      # ES映射配置
│   ├── mysql/
│   │   ├── init/
│   │   │   └── 01-init.sql       # 数据库初始化
│   │   └── conf.d/
│   └── canal-admin/
│       └── conf/
```

## 🔍 验证部署

### 1. 检查服务状态
```bash
# 查看所有服务状态
docker-compose ps

# 检查服务健康状态
curl http://localhost:11110/actuator/health  # Canal Server
curl http://localhost:8081/actuator/health   # Canal Adapter
curl http://localhost:9200/_cluster/health    # Elasticsearch
```

### 2. 验证数据同步
```bash
# 检查ES索引
curl http://localhost:9200/_cat/indices

# 检查同步数据量
curl http://localhost:9200/addons_customer_pro_clues/_count

# 查看同步的文档
curl http://localhost:9200/addons_customer_pro_clues/_search?size=5
```

### 3. 测试数据插入
```bash
# 连接到MySQL
docker-compose exec mysql mysql -u canal -pcanal dzhgo

# 插入测试数据
INSERT INTO addons_customer_pro_clues (id, name, mobile, status, services_ids) 
VALUES ('test_001', '测试用户', '13800138000', 0, '85');

# 检查ES是否同步
curl http://localhost:9200/addons_customer_pro_clues/_search?q=name:测试用户
```

## 📈 监控和日志

### 查看日志
```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f canal-adapter
docker-compose logs -f canal-server
docker-compose logs -f mysql
docker-compose logs -f elasticsearch
```

### 监控指标
```bash
# Canal Server指标
curl http://localhost:11110/actuator/metrics

# Canal Adapter指标
curl http://localhost:8081/actuator/metrics

# ES集群状态
curl http://localhost:9200/_cluster/stats
```

## 🛠️ 故障排除

### 常见问题

#### 1. 服务启动失败
```bash
# 查看详细错误日志
docker-compose logs [service-name]

# 重启服务
docker-compose restart [service-name]

# 重新构建并启动
docker-compose up -d --build
```

#### 2. 数据同步失败
```bash
# 检查Canal Server连接
docker-compose exec canal-server netstat -tlnp

# 检查MySQL binlog
docker-compose exec mysql mysql -u canal -pcanal -e "SHOW MASTER STATUS;"

# 检查ES连接
curl http://localhost:9200/_cluster/health
```

#### 3. 内存不足
```bash
# 调整ES内存
# 编辑 docker-compose.yml 中的 ES_JAVA_OPTS
# 例如: -Xms256m -Xmx256m
```

### 调试命令
```bash
# 进入容器调试
docker-compose exec canal-adapter bash
docker-compose exec canal-server bash
docker-compose exec mysql bash
docker-compose exec elasticsearch bash

# 检查网络连接
docker network ls
docker network inspect canal-network
```

## 🔄 数据管理

### 备份数据
```bash
# 备份MySQL数据
docker-compose exec mysql mysqldump -u canal -pcanal dzhgo > backup.sql

# 备份ES数据
curl -X GET "localhost:9200/addons_customer_pro_clues/_search?size=1000" > es_backup.json
```

### 恢复数据
```bash
# 恢复MySQL数据
docker-compose exec -T mysql mysql -u canal -pcanal dzhgo < backup.sql

# 恢复ES数据
curl -X POST "localhost:9200/_bulk" -H "Content-Type: application/json" --data-binary @es_backup.json
```

### 清理数据
```bash
# 停止并清理所有数据
docker-compose down -v

# 重新初始化
docker-compose up -d
```

## 📝 常用命令

```bash
# 服务管理
docker-compose up -d              # 启动所有服务
docker-compose down               # 停止所有服务
docker-compose restart            # 重启所有服务
docker-compose ps                 # 查看服务状态
docker-compose logs -f            # 查看实时日志

# 单个服务管理
docker-compose restart canal-adapter    # 重启Canal Adapter
docker-compose logs -f canal-server     # 查看Canal Server日志
docker-compose exec mysql bash          # 进入MySQL容器

# 数据操作
docker-compose exec mysql mysql -u canal -pcanal dzhgo  # 连接MySQL
curl http://localhost:9200/addons_customer_pro_clues/_count  # 查询ES数据量
```

## 🎯 性能优化

### 1. 内存配置
```yaml
# 在 docker-compose.yml 中调整
elasticsearch:
  environment:
    - "ES_JAVA_OPTS=-Xms1g -Xmx1g"  # 增加ES内存
```

### 2. 批量配置
```yaml
# 在 docker/canal-adapter/conf/application.yml 中调整
canal.conf:
  batchSize: 2000        # 增加批量大小
  syncBatchSize: 2000
  commitBatch: 2000
```

### 3. 网络优化
```yaml
# 在 docker-compose.yml 中添加
services:
  canal-adapter:
    ulimits:
      nofile:
        soft: 65536
        hard: 65536
```

## 🔐 安全配置

### 1. 修改默认密码
```bash
# 修改MySQL密码
docker-compose exec mysql mysql -u root -ppassword
ALTER USER 'canal'@'%' IDENTIFIED BY 'new_password';

# 更新配置文件中的密码
# 编辑 docker/canal-adapter/conf/application.yml
```

### 2. 启用ES安全
```yaml
# 在 docker-compose.yml 中修改
elasticsearch:
  environment:
    - xpack.security.enabled=true
    - ELASTIC_PASSWORD=your_password
```

## 📞 技术支持

如果遇到问题，可以：

1. 查看详细日志: `docker-compose logs -f`
2. 检查服务状态: `docker-compose ps`
3. 验证网络连接: `docker network inspect canal-network`
4. 查看系统资源: `docker stats` 