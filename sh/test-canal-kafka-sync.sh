#!/bin/bash

# Canal数据同步测试脚本
echo "=========================================="
echo "        Canal数据同步测试脚本"
echo "=========================================="

# 检查服务状态
echo "检查服务状态..."
docker-compose ps

# 检查ES索引
echo "检查ES索引..."
curl -s "http://localhost:9200/_cat/indices?v"

# 检查ES中的clues数据
echo "检查ES中的clues数据..."
curl -s "http://localhost:9200/addons_customer_pro_clues/_search?pretty" | jq '.hits.total.value'

# 检查MySQL中的数据
echo "检查MySQL中的clues数据..."
docker exec mysql-canal mysql -u canal -pcanal123 -e "SELECT COUNT(*) as clues_count FROM dzhgo_go.addons_customer_pro_clues;"

echo "检查MySQL中的attr数据..."
docker exec mysql-canal mysql -u canal -pcanal123 -e "SELECT COUNT(*) as attr_count FROM dzhgo_go.addons_customer_pro_clues_attr;"

# 插入测试数据
echo "插入测试数据..."
docker exec mysql-canal mysql -u canal -pcanal123 -e "
INSERT INTO dzhgo_go.addons_customer_pro_clues (id, serialId, name, mobile, status, services_ids, createTime, updateTime) VALUES
('test_canal_001', 1001, 'Canal测试用户1', '13900139001', 0, '85,86', NOW(), NOW()),
('test_canal_002', 1002, 'Canal测试用户2', '13900139002', 0, '85', NOW(), NOW())
ON DUPLICATE KEY UPDATE updateTime = NOW();
"

docker exec mysql-canal mysql -u canal -pcanal123 -e "
INSERT INTO dzhgo_go.addons_customer_pro_clues_attr (clues_id, guest_ip_info) VALUES
('test_canal_001', '192.168.100.100'),
('test_canal_002', '192.168.100.101')
ON DUPLICATE KEY UPDATE guest_ip_info = VALUES(guest_ip_info);
"

echo "等待数据同步..."
sleep 10

# 再次检查ES中的数据
echo "检查ES中的测试数据..."
curl -s "http://localhost:9200/addons_customer_pro_clues/_search?q=id:test_canal_001&pretty"

echo "=========================================="
echo "测试完成！"
echo "==========================================" 