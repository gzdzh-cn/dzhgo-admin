# Canal Adapter 配置指南

## 概述

本配置用于通过Canal Adapter实现MySQL中`addons_customer_pro_clues`和`addons_customer_pro_clues_attr`表的增量同步到Elasticsearch的`addons_customer_pro_clues`索引。

## 架构说明

```
MySQL (dzhgo_go) 
    ↓ (Canal Server监听binlog)
Kafka (消息队列)
    ↓ (Canal Adapter消费)
Elasticsearch (addons_customer_pro_clues索引)
```

## 配置说明

### 1. Canal Server配置

**文件**: `docker/canal-server/conf/example/instance.properties`

- 监听MySQL的binlog
- 过滤表：`addons_customer_pro_clues` 和 `addons_customer_pro_clues_attr`
- 发送消息到Kafka topic

### 2. Canal Adapter配置

**文件**: `docker/canal-adapter/conf/es7/clues_aggregated.yml`

- 数据源：MySQL (dzhgo_go)
- 目标：Elasticsearch (addons_customer_pro_clues)
- SQL查询：LEFT JOIN聚合clues和attr表
- 字段映射：按照`customer_pro_clues_es.go`中的逻辑

### 3. ES索引映射

**文件**: `docker/canal-adapter/conf/es7/clues_mapping.json`

定义了ES索引的字段类型映射，包括：
- 文本字段：name, keywords, filter_remark, guest_ip_info
- 关键词字段：id, mobile, wechat, account_id等
- 日期字段：createTime, updateTime, ocean_time, deleted_at
- 数字字段：serialId, status

## 部署步骤

### 1. 启动服务

```bash
# 启动所有服务
./deploy-canal-kafka.sh
```

### 2. 初始化ES索引

```bash
# 自动执行，或手动执行
docker exec canal-adapter /home/admin/canal-adapter/conf/es7/init_es_index.sh
```

### 3. 配置Canal Admin

1. 访问 http://localhost:8089
2. 登录：admin/123456
3. 配置数据源和实例
4. 启动同步任务

### 4. 测试数据同步

```bash
# 运行测试脚本
./test-canal-kafka-sync.sh
```

## 数据同步逻辑

### 字段映射

根据`customer_pro_clues_es.go`中的`buildESDocument`方法，同步以下字段：

**clues表字段**:
- id, serialId, guest_id, name, mobile, wechat
- keywords, status, account_id, project_id
- services_id, services_ids, source_from, followup_type
- level, filter_remark, created_id, filter_group_ids
- createTime, updateTime, deleted_at, ocean_time

**attr表字段**:
- guest_ip_info (通过LEFT JOIN关联)

### 聚合逻辑

```sql
SELECT 
  c.id, c.serialId, c.guest_id, c.name, c.mobile, c.wechat,
  c.keywords, c.status, c.account_id, c.project_id,
  c.services_id, c.services_ids, c.source_from, c.followup_type,
  c.level, c.filter_remark, c.created_id, c.filter_group_ids,
  c.createTime, c.updateTime, c.deleted_at, c.ocean_time,
  COALESCE(ca.guest_ip_info, '') as guest_ip_info
FROM addons_customer_pro_clues c
LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
```

## 监控和调试

### 1. 服务状态检查

```bash
# 检查所有服务状态
docker-compose ps

# 检查ES索引
curl http://localhost:9200/_cat/indices?v

# 检查ES数据
curl http://localhost:9200/addons_customer_pro_clues/_search?pretty
```

### 2. 日志监控

```bash
# Canal Server日志
docker logs -f canal-server

# Canal Adapter日志
docker logs -f canal-adapter

# Kafka消息
docker exec -it kafka kafka-console-consumer --bootstrap-server localhost:9092 --topic canal_topic --from-beginning
```

### 3. 数据验证

```bash
# 检查MySQL数据
docker exec mysql-canal mysql -u canal -pcanal123 -e "SELECT COUNT(*) FROM dzhgo_go.addons_customer_pro_clues;"

# 检查ES数据
curl -s "http://localhost:9200/addons_customer_pro_clues/_count" | jq '.'
```

## 故障排除

### 1. 常见问题

**问题**: Canal Adapter无法连接ES
**解决**: 检查ES服务状态和网络连接

**问题**: 数据同步延迟
**解决**: 检查Kafka消息积压和Canal Adapter处理速度

**问题**: 字段映射错误
**解决**: 检查ES索引映射和SQL查询字段

### 2. 重启服务

```bash
# 重启特定服务
docker-compose restart canal-adapter

# 重启所有服务
docker-compose restart
```

### 3. 清理数据

```bash
# 删除ES索引
curl -X DELETE http://localhost:9200/addons_customer_pro_clues

# 重新创建索引
docker exec canal-adapter /home/admin/canal-adapter/conf/es7/init_es_index.sh
```

## 性能优化

### 1. 批量配置

- `commitBatch`: 1000 (每批处理1000条)
- `syncBatchSize`: 1000 (同步批次大小)

### 2. 索引优化

- 使用keyword类型存储精确匹配字段
- 使用text类型存储全文搜索字段
- 合理设置分片和副本数

### 3. 监控指标

- 同步延迟时间
- 数据处理速度
- 错误率统计

## 访问地址

- **Canal Admin**: http://localhost:8089 (admin/123456)
- **Elasticsearch**: http://localhost:9200
- **Kibana**: http://localhost:5601
- **Kafka**: localhost:9092 