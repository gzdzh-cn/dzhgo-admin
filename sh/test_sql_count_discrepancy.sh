#!/bin/bash

# 测试SQL结果数量不一致的问题
echo "=== SQL结果数量不一致分析 ==="
echo "时间: $(date)"
echo ""

# 数据库连接信息（请根据实际情况修改）
DB_HOST="localhost"
DB_PORT="3306"
DB_NAME="dzhgo_go"
DB_USER="root"
DB_PASS="your_password"

# 用户ID
USER_ID="1152921504606846976"

echo "=== 执行第一条SQL（COUNT查询）==="
COUNT_SQL="SELECT COUNT(1) FROM base_sys_notice AS n RIGHT JOIN base_sys_notice_user_read AS nr ON (nr.notice_id = n.id) WHERE (nr.user_id = '$USER_ID') AND n.deleted_at IS NULL AND nr.deleted_at IS NULL"
echo "SQL: $COUNT_SQL"
echo ""

echo "=== 执行第二条SQL（详细查询）==="
DETAIL_SQL="SELECT n.id,n.createTime,n.title,n.remark,n.noType,nr.readTime,nr.status FROM base_sys_notice AS n RIGHT JOIN base_sys_notice_user_read AS nr ON (nr.notice_id = n.id) WHERE (nr.user_id = '$USER_ID') AND n.deleted_at IS NULL AND nr.deleted_at IS NULL ORDER BY n.createTime DESC LIMIT 15 OFFSET 0"
echo "SQL: $DETAIL_SQL"
echo ""

echo "=== 分析查询 ==="
echo "1. 检查base_sys_notice_user_read表中的记录数："
CHECK_USER_READ="SELECT COUNT(*) as user_read_count FROM base_sys_notice_user_read WHERE user_id = '$USER_ID' AND deleted_at IS NULL"
echo "SQL: $CHECK_USER_READ"
echo ""

echo "2. 检查base_sys_notice表中的记录数："
CHECK_NOTICE="SELECT COUNT(*) as notice_count FROM base_sys_notice WHERE deleted_at IS NULL"
echo "SQL: $CHECK_NOTICE"
echo ""

echo "3. 检查RIGHT JOIN的详细结果："
DETAIL_JOIN="SELECT 
    nr.id as read_id,
    nr.notice_id,
    nr.user_id,
    nr.readTime,
    nr.status,
    n.id as notice_id_from_notice,
    n.title,
    n.deleted_at as notice_deleted_at,
    nr.deleted_at as read_deleted_at
FROM base_sys_notice_user_read AS nr 
LEFT JOIN base_sys_notice AS n ON (nr.notice_id = n.id)
WHERE nr.user_id = '$USER_ID' 
ORDER BY nr.id"
echo "SQL: $DETAIL_JOIN"
echo ""

echo "=== 建议的修复方案 ==="
echo "1. 使用LEFT JOIN替代RIGHT JOIN："
FIXED_COUNT="SELECT COUNT(1) FROM base_sys_notice_user_read AS nr LEFT JOIN base_sys_notice AS n ON (nr.notice_id = n.id) WHERE (nr.user_id = '$USER_ID') AND n.deleted_at IS NULL AND nr.deleted_at IS NULL"
echo "SQL: $FIXED_COUNT"
echo ""

echo "2. 或者使用INNER JOIN："
INNER_COUNT="SELECT COUNT(1) FROM base_sys_notice AS n INNER JOIN base_sys_notice_user_read AS nr ON (nr.notice_id = n.id) WHERE (nr.user_id = '$USER_ID') AND n.deleted_at IS NULL AND nr.deleted_at IS NULL"
echo "SQL: $INNER_COUNT"
echo ""

echo "=== 执行建议 ==="
echo "请根据您的数据库连接信息修改脚本中的DB_HOST、DB_PORT、DB_NAME、DB_USER、DB_PASS变量"
echo "然后运行以下命令来执行测试："
echo "mysql -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS $DB_NAME -e \"$COUNT_SQL\""
echo "mysql -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS $DB_NAME -e \"$DETAIL_SQL\""
