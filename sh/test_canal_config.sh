#!/bin/bash

# Canal Adapter 配置测试脚本

echo "=== Canal Adapter 配置测试 ==="
echo ""

# 测试SQL查询
echo "🔍 测试SQL查询..."
TEST_SQL="SELECT 
          c.id,
          c.serialId,
          c.guest_id,
          c.name,
          c.mobile,
          c.wechat,
          c.keywords,
          c.status,
          c.account_id,
          c.project_id,
          c.services_id,
          c.services_ids,
          c.source_from,
          c.followup_type,
          c.level,
          c.filter_remark,
          c.created_id,
          c.filter_group_ids,
          c.deleted_at,
          DATE_FORMAT(c.create_time, '%Y-%m-%dT%H:%i:%s+08:00') as createTime,
          DATE_FORMAT(c.update_time, '%Y-%m-%dT%H:%i:%s+08:00') as updateTime,
          DATE_FORMAT(c.ocean_time, '%Y-%m-%dT%H:%i:%s+08:00') as ocean_time,
          DATE_FORMAT(c.deleted_at, '%Y-%m-%dT%H:%i:%s+08:00') as deleted_at,
          ca.guest_ip_info
        FROM addons_customer_pro_clues c
        LEFT JOIN addons_customer_pro_clues_attr ca ON c.id = ca.clues_id
        WHERE c.deleted_at IS NULL
        LIMIT 5"

echo "SQL查询:"
echo "$TEST_SQL"
echo ""

# 检查配置文件
echo "🔍 检查配置文件..."
if [ -f "canal_adapter_config/application.yml" ]; then
    echo "✅ application.yml 存在"
else
    echo "❌ application.yml 不存在"
fi

if [ -f "canal_adapter_config/es7/clues.yml" ]; then
    echo "✅ clues.yml 存在"
else
    echo "❌ clues.yml 不存在"
fi

echo ""

# 显示配置摘要
echo "📋 配置摘要:"
echo "主表: addons_customer_pro_clues"
echo "关联表: addons_customer_pro_clues_attr"
echo "关联条件: c.id = ca.clues_id"
echo "过滤条件: c.deleted_at IS NULL"
echo "ES索引: addons_customer_pro_clues"
echo ""

# 字段映射检查
echo "🔧 字段映射检查:"
echo "Clues表字段 (19个):"
echo "  - id, serialId, guest_id, name, mobile, wechat"
echo "  - keywords, status, account_id, project_id"
echo "  - services_id, services_ids, source_from, followup_type"
echo "  - level, filter_remark, created_id, filter_group_ids, deleted_at"
echo ""
echo "时间字段 (4个):"
echo "  - createTime, updateTime, ocean_time, deleted_at"
echo ""
echo "Attr表字段 (1个):"
echo "  - guest_ip_info"
echo ""

# 性能配置检查
echo "⚡ 性能配置:"
echo "批量大小: 1000"
echo "增量同步: 基于 update_time"
echo "时间格式: ISO 8601"
echo ""

# 验证命令
echo "🔍 验证命令:"
echo ""
echo "1. 测试MySQL连接:"
echo "mysql -u root -p -h localhost -e \"SELECT COUNT(*) FROM addons_customer_pro_clues WHERE deleted_at IS NULL;\""
echo ""
echo "2. 测试ES连接:"
echo "curl -X GET \"localhost:9200/_cluster/health\""
echo ""
echo "3. 检查Canal Server:"
echo "curl http://localhost:11110/actuator/health"
echo ""
echo "4. 检查Canal Adapter:"
echo "curl http://localhost:8081/actuator/health"
echo ""
echo "5. 查看同步日志:"
echo "tail -f logs/adapter/adapter.log"
echo ""

echo "=== 测试完成 ==="
echo "请根据上述信息验证配置是否正确" 