#!/bin/bash

# Canal Adapter 安装和配置脚本

echo "=== Canal Adapter 安装配置 ==="
echo ""

# 1. 下载Canal Adapter
echo "🔧 步骤1: 下载Canal Adapter"
CANAL_VERSION="1.1.6"
CANAL_ADAPTER_URL="https://github.com/alibaba/canal/releases/download/canal-$CANAL_VERSION/canal.adapter-$CANAL_VERSION.tar.gz"

echo "下载地址: $CANAL_ADAPTER_URL"
echo "wget $CANAL_ADAPTER_URL"
echo "tar -xzf canal.adapter-$CANAL_VERSION.tar.gz"
echo "cd canal.adapter-$CANAL_VERSION"
echo ""

# 2. 配置Canal Adapter
echo "🔧 步骤2: 配置Canal Adapter"
echo ""

echo "编辑 conf/application.yml:"
cat > canal_adapter_application.yml << 'EOF'
server:
  port: 8081
spring:
  jackson:
    date-format: yyyy-MM-dd HH:mm:ss
    time-zone: GMT+8

canal.conf:
  mode: tcp #tcp kafka rocketMQ rabbitMQ
  flatMessage: true
  batchSize: 1000
  syncBatchSize: 1000
  retries: 0
  timeout:
  accessKey:
  secretKey:
  consumerProperties:
    # canal tcp consumer
    canal.tcp.server.host: 127.0.0.1:11111
    canal.tcp.batch.size: 500
    canal.tcp.timeout: 100
    # ...
  srcDataSources:
    defaultDS:
      url: jdbc:mysql://127.0.0.1:3306/dzhgo?useUnicode=true&characterEncoding=UTF-8&useSSL=false&serverTimezone=Asia/Shanghai
      username: root
      password: password
      driver: com.mysql.cj.jdbc.Driver
  canalAdapters:
  - instance: example # canal instance Name or mq topic name
    groups:
    - groupId: g1
      outerAdapters:
      - name: es7
        hosts: 127.0.0.1:9300
        properties:
          cluster.name: elasticsearch
          xpack.security.enabled: false
EOF

echo "配置文件已生成: canal_adapter_application.yml"
echo ""

# 3. 配置ES适配器
echo "🔧 步骤3: 配置ES适配器"
echo ""

echo "编辑 conf/es7/clues.yml:"
cat > canal_adapter_es7_clues.yml << 'EOF'
dataSourceKey: defaultDS
destination: example
groupId: g1
esMapping:
  _index: addons_customer_pro_clues
  _type: _doc
  _id: _id
  upsert: true
  pk: id
  sql: "SELECT 
          c.id,
          c.name,
          c.mobile,
          c.wechat,
          c.keywords,
          c.status,
          c.services_ids,
          c.source_from,
          c.followup_type,
          c.level,
          c.create_time,
          c.update_time,
          ca.guest_ip_info,
          ca.address,
          ca.city,
          ca.from_page
        FROM addons_customer_pro_clues c
        LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
        WHERE c.deleted_at IS NULL"
  etlCondition: "where c.update_time>={}"
  commitBatch: 1000
EOF

echo "ES配置文件已生成: canal_adapter_es7_clues.yml"
echo ""

# 4. 配置聚合表
echo "🔧 步骤4: 配置聚合表"
echo ""

echo "编辑 conf/es7/clues_aggregated.yml:"
cat > canal_adapter_es7_aggregated.yml << 'EOF'
dataSourceKey: defaultDS
destination: example
groupId: g1
esMapping:
  _index: addons_customer_pro_clues_aggregated
  _type: _doc
  _id: _id
  upsert: true
  pk: id
  sql: "SELECT 
          c.id,
          c.name,
          c.mobile,
          c.wechat,
          c.keywords,
          c.status,
          c.services_ids,
          c.source_from,
          c.followup_type,
          c.level,
          c.create_time,
          c.update_time,
          ca.guest_ip_info,
          ca.address,
          ca.city,
          ca.from_page,
          ca.talk_page,
          ca.land_page,
          ca.se,
          ca.ip,
          ca.chat_content,
          CONCAT(c.name, ' ', IFNULL(ca.address, ''), ' ', IFNULL(ca.city, '')) as search_text
        FROM addons_customer_pro_clues c
        LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
        WHERE c.deleted_at IS NULL"
  etlCondition: "where c.update_time>={}"
  commitBatch: 1000
EOF

echo "聚合表配置文件已生成: canal_adapter_es7_aggregated.yml"
echo ""

# 5. 启动命令
echo "🔧 步骤5: 启动命令"
echo ""
echo "启动Canal Server:"
echo "cd canal-server"
echo "./bin/startup.sh"
echo ""
echo "启动Canal Adapter:"
echo "cd canal.adapter-$CANAL_VERSION"
echo "./bin/startup.sh"
echo ""

# 6. 验证命令
echo "🔧 步骤6: 验证命令"
echo ""
echo "检查Canal Server状态:"
echo "curl http://localhost:11110/actuator/health"
echo ""
echo "检查Canal Adapter状态:"
echo "curl http://localhost:8081/actuator/health"
echo ""
echo "查看同步日志:"
echo "tail -f logs/adapter/adapter.log"
echo ""

echo "=== 配置完成 ==="
echo "请按照上述步骤进行配置和启动" 