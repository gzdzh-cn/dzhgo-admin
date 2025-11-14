# Canal Adapter 配置完成总结

## 🎯 配置目标

实现MySQL中`addons_customer_pro_clues`和`addons_customer_pro_clues_attr`表的增量同步到Elasticsearch的`addons_customer_pro_clues`索引，按照`customer_pro_clues_es.go`中的逻辑进行字段聚合。

## 📁 已创建的配置文件

### 1. Canal Adapter配置
- **`docker/canal-adapter/conf/es7/clues_aggregated.yml`** - 主要的聚合同步配置
- **`docker/canal-adapter/conf/es7/clues_mapping.json`** - ES索引映射定义
- **`docker/canal-adapter/conf/es7/init_es_index.sh`** - ES索引初始化脚本

### 2. Canal Server配置
- **`docker/canal-server/conf/example/instance.properties`** - 已更新表过滤配置

### 3. 部署和测试脚本
- **`deploy-canal-kafka.sh`** - 完整部署脚本
- **`quick-start.sh`** - 快速启动脚本
- **`test-canal-kafka-sync.sh`** - 数据同步测试脚本
- **`check-canal-status.sh`** - 系统状态检查脚本

### 4. 文档
- **`canal_adapter_guide.md`** - 详细配置指南
- **`canal_adapter_setup.md`** - 本总结文档

## 🔧 核心配置说明

### 1. 聚合查询逻辑

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

### 2. ES索引映射

按照`customer_pro_clues_es.go`中的`buildESDocument`方法，定义了完整的字段映射：
- **文本字段**: name, keywords, filter_remark, guest_ip_info
- **关键词字段**: id, mobile, wechat, account_id等
- **日期字段**: createTime, updateTime, ocean_time, deleted_at
- **数字字段**: serialId, status

### 3. 增量同步配置

- **监听表**: `addons_customer_pro_clues` 和 `addons_customer_pro_clues_attr`
- **批量大小**: 1000条/批
- **同步模式**: 实时增量同步
- **错误重试**: 3次重试机制

## 🚀 快速开始

### 1. 启动系统
```bash
./quick-start.sh
```

### 2. 检查状态
```bash
./check-canal-status.sh
```

### 3. 测试同步
```bash
./test-canal-kafka-sync.sh
```

## 📊 监控和验证

### 1. 服务状态监控
- Canal Admin: http://localhost:8089 (admin/123456)
- Elasticsearch: http://localhost:9200
- Kibana: http://localhost:5601

### 2. 数据验证命令
```bash
# 检查ES索引
curl http://localhost:9200/_cat/indices?v

# 检查ES数据
curl http://localhost:9200/addons_customer_pro_clues/_search?pretty

# 检查MySQL数据
docker exec mysql-canal mysql -u canal -pcanal123 -e "SELECT COUNT(*) FROM dzhgo_go.addons_customer_pro_clues;"
```

### 3. 日志监控
```bash
# Canal Adapter日志
docker logs -f canal-adapter

# Canal Server日志
docker logs -f canal-server

# Kafka消息
docker exec -it kafka kafka-console-consumer --bootstrap-server localhost:9092 --topic canal_topic --from-beginning
```

## 🔄 数据同步流程

1. **MySQL变更** → 触发binlog
2. **Canal Server** → 监听binlog，发送到Kafka
3. **Canal Adapter** → 消费Kafka消息，执行聚合查询
4. **Elasticsearch** → 更新索引数据

## ⚠️ 注意事项

### 1. 性能考虑
- 聚合查询可能影响性能，建议在attr表上建立clues_id索引
- 大批量数据同步时，适当调整批量大小
- 监控ES集群资源使用情况

### 2. 数据一致性
- 确保MySQL的binlog配置正确
- 定期验证ES和MySQL数据一致性
- 配置适当的错误处理和重试机制

### 3. 监控告警
- 监控同步延迟
- 监控错误率
- 监控系统资源使用

## 🛠️ 故障排除

### 1. 常见问题
- **同步延迟**: 检查Kafka消息积压和Canal Adapter处理速度
- **连接失败**: 检查网络连接和服务状态
- **字段映射错误**: 检查ES索引映射和SQL查询字段

### 2. 重启服务
```bash
# 重启特定服务
docker-compose restart canal-adapter

# 重启所有服务
docker-compose restart
```

### 3. 清理重建
```bash
# 删除ES索引
curl -X DELETE http://localhost:9200/addons_customer_pro_clues

# 重新创建索引
docker exec canal-adapter /home/admin/canal-adapter/conf/es7/init_es_index.sh
```

## 📈 优化建议

### 1. 性能优化
- 根据数据量调整批量大小
- 优化ES索引分片和副本配置
- 使用合适的字段类型和映射

### 2. 监控优化
- 设置同步状态监控
- 配置错误告警
- 监控系统资源使用

### 3. 数据优化
- 定期清理过期数据
- 优化查询条件
- 建立合适的索引

## ✅ 配置完成状态

- [x] Canal Server配置
- [x] Canal Adapter配置
- [x] ES索引映射
- [x] 聚合查询逻辑
- [x] 部署脚本
- [x] 测试脚本
- [x] 监控脚本
- [x] 文档说明

## 🎉 配置完成！

现在你可以使用上述脚本和配置来启动Canal系统，实现clues和attr表的增量同步到ES。所有配置都按照`customer_pro_clues_es.go`中的逻辑进行了优化，确保数据同步的准确性和性能。 