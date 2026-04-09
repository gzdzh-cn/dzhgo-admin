-- ============================================================
-- 线索查询优化 - 生产环境部署脚本
-- 
-- 执行前请确认：
--   1. 在低峰期执行
--   2. 备份 addons_customer_pro_clues 表数据
--   3. 全部执行完成后再部署 Go 代码
--
-- 执行方式：
--   mysql -h <host> -P <port> -u<user> -p<password> <database> < production_deploy.sql
-- ============================================================

SELECT CONCAT('=== 线索查询优化部署开始: ', NOW(), ' ===') AS info;

-- ============================================================
-- Step 1: 添加覆盖索引（COUNT + ORDER BY 只读索引不回表）
-- ============================================================
SELECT 'Step 1: 添加覆盖索引...' AS info;

SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4) 
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_cover_count_sort') > 0,
    'SELECT "索引 idx_cover_count_sort 已存在，跳过" AS info',
    'ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_cover_count_sort` (`status`, `dtype`, `deleted_at`, `createTime` DESC)'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

SELECT 'Step 1 完成' AS info;

-- ============================================================
-- Step 2: 删除冗余索引（与覆盖索引前缀冲突，导致优化器选错）
-- ============================================================
SELECT 'Step 2: 删除冗余索引...' AS info;

-- idx_status_dtype_deleted: status,dtype,deleted_at — 被 idx_cover_count_sort 完全覆盖
SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_status_dtype_deleted') > 0,
    'ALTER TABLE `addons_customer_pro_clues` DROP INDEX `idx_status_dtype_deleted`',
    'SELECT "idx_status_dtype_deleted 不存在，跳过" AS info'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- idx_addons_customer_pro_clues_status_dtype: status,dtype,deleted_at — 同上
SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_addons_customer_pro_clues_status_dtype') > 0,
    'ALTER TABLE `addons_customer_pro_clues` DROP INDEX `idx_addons_customer_pro_clues_status_dtype`',
    'SELECT "idx_addons_customer_pro_clues_status_dtype 不存在，跳过" AS info'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- idx_addons_customer_pro_clues_mobile: status,dtype,mobile — 前缀冲突，后续重建为单列索引
SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_addons_customer_pro_clues_mobile') > 0,
    'ALTER TABLE `addons_customer_pro_clues` DROP INDEX `idx_addons_customer_pro_clues_mobile`',
    'SELECT "idx_addons_customer_pro_clues_mobile 不存在，跳过" AS info'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- idx_addons_customer_pro_clues_guest_id: status,dtype,guest_id — 同上
SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_addons_customer_pro_clues_guest_id') > 0,
    'ALTER TABLE `addons_customer_pro_clues` DROP INDEX `idx_addons_customer_pro_clues_guest_id`',
    'SELECT "idx_addons_customer_pro_clues_guest_id 不存在，跳过" AS info'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

SELECT 'Step 2 完成' AS info;

-- ============================================================
-- Step 3: 修复字段类型（longtext → 合理的 varchar/text）
-- ============================================================
SELECT 'Step 3: 修复字段类型...' AS info;

ALTER TABLE `addons_customer_pro_clues` 
    MODIFY COLUMN `mobile` varchar(32) DEFAULT NULL COMMENT '手机号',
    MODIFY COLUMN `wechat` varchar(128) DEFAULT NULL COMMENT '微信号',
    MODIFY COLUMN `weixin` varchar(128) DEFAULT NULL COMMENT 'weixin',
    MODIFY COLUMN `guest_id` varchar(64) DEFAULT NULL COMMENT '53标识',
    MODIFY COLUMN `account_id` varchar(64) DEFAULT NULL COMMENT '账户ID',
    MODIFY COLUMN `project_id` varchar(64) DEFAULT NULL COMMENT '项目ID',
    MODIFY COLUMN `name` varchar(200) DEFAULT NULL COMMENT '标题',
    MODIFY COLUMN `created_name` varchar(64) DEFAULT NULL COMMENT '创建者',
    MODIFY COLUMN `services_ids` varchar(500) DEFAULT NULL COMMENT '分配过的客服ID',
    MODIFY COLUMN `source_from` varchar(32) DEFAULT NULL COMMENT '来源',
    MODIFY COLUMN `keywords` varchar(500) DEFAULT NULL COMMENT '关键字',
    MODIFY COLUMN `level` varchar(100) DEFAULT NULL COMMENT '线索等级',
    MODIFY COLUMN `remark` text DEFAULT NULL COMMENT '备注',
    MODIFY COLUMN `filterRemark` text DEFAULT NULL COMMENT '过滤原因',
    MODIFY COLUMN `filter_group_ids` varchar(500) DEFAULT NULL COMMENT '组ID';

SELECT 'Step 3 完成' AS info;

-- ============================================================
-- Step 4: 重建 mobile 和 guest_id 单列索引
-- ============================================================
SELECT 'Step 4: 重建单列索引...' AS info;

SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_clues_mobile') > 0,
    'SELECT "idx_clues_mobile 已存在，跳过" AS info',
    'ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_clues_mobile` (`mobile`)'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM information_schema.statistics 
     WHERE CONVERT(table_schema USING utf8mb4) = CONVERT(DATABASE() USING utf8mb4)
       AND CONVERT(table_name USING utf8mb4) = 'addons_customer_pro_clues'
       AND CONVERT(index_name USING utf8mb4) = 'idx_clues_guest_id') > 0,
    'SELECT "idx_clues_guest_id 已存在，跳过" AS info',
    'ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_clues_guest_id` (`guest_id`)'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

SELECT 'Step 4 完成' AS info;

-- ============================================================
-- Step 5: 新建关联表（排序规则必须与主表一致）
-- ============================================================
SELECT 'Step 5: 创建关联表...' AS info;

CREATE TABLE IF NOT EXISTS `addons_customer_pro_clues_followup_type` (
    `id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'ID',
    `clues_id` VARCHAR(64) NOT NULL COMMENT '线索ID',
    `type_value` VARCHAR(32) NOT NULL COMMENT '跟进类型值',
    `createTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updateTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_clues_followup_clues_id` (`clues_id`),
    UNIQUE KEY `uk_clues_followup` (`clues_id`, `type_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='线索-跟进类型关联表';

CREATE TABLE IF NOT EXISTS `addons_customer_pro_clues_filter_group` (
    `id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'ID',
    `clues_id` VARCHAR(64) NOT NULL COMMENT '线索ID',
    `group_id` VARCHAR(64) NOT NULL COMMENT '客服组ID',
    `createTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updateTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_clues_filter_clues_id` (`clues_id`),
    UNIQUE KEY `uk_clues_filter_group` (`clues_id`, `group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='线索-客服组关联表';

CREATE TABLE IF NOT EXISTS `addons_customer_pro_clues_level` (
    `id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'ID',
    `clues_id` VARCHAR(64) NOT NULL COMMENT '线索ID',
    `level_value` VARCHAR(32) NOT NULL COMMENT '等级值',
    `createTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updateTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_clues_level_clues_id` (`clues_id`),
    UNIQUE KEY `uk_clues_level` (`clues_id`, `level_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='线索-等级关联表';

-- clues_service 表（首次部署由框架自动创建；此处仅为补齐字段用）
CREATE TABLE IF NOT EXISTS `addons_customer_pro_clues_service` (
    `id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'ID',
    `clues_id` VARCHAR(64) NOT NULL COMMENT '线索ID',
    `user_id` VARCHAR(64) NOT NULL COMMENT '客服用户ID',
    `createTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updateTime` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_clues_service_clues_id` (`clues_id`),
    INDEX `idx_clues_service_user_id` (`user_id`),
    UNIQUE KEY `uk_clues_service` (`clues_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='线索-客服分配关联表';

-- 确保排序规则与主表一致（防表已存在但排序规则不对）
ALTER TABLE `addons_customer_pro_clues_followup_type` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
ALTER TABLE `addons_customer_pro_clues_filter_group` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
ALTER TABLE `addons_customer_pro_clues_level` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
ALTER TABLE `addons_customer_pro_clues_service` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

SELECT 'Step 5 完成' AS info;

-- ============================================================
-- Step 6: 迁移 followup_type 数据
-- ============================================================
SELECT 'Step 6: 迁移 followup_type 数据...' AS info;

TRUNCATE TABLE `addons_customer_pro_clues_followup_type`;

INSERT IGNORE INTO `addons_customer_pro_clues_followup_type` (`id`, `clues_id`, `type_value`, `createTime`, `updateTime`)
SELECT 
    CONCAT(c.id, '_ft_', TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.followup_type, ',', n.n), ',', -1))) AS id,
    c.id AS clues_id,
    TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.followup_type, ',', n.n), ',', -1)) AS type_value,
    COALESCE(c.createTime, NOW(3)) AS createTime,
    NOW(3) AS updateTime
FROM `addons_customer_pro_clues` c
CROSS JOIN (
    SELECT 1 n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5
    UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
) n
WHERE c.followup_type IS NOT NULL 
  AND c.followup_type != '' 
  AND n.n <= 1 + LENGTH(c.followup_type) - LENGTH(REPLACE(c.followup_type, ',', ''))
  AND c.deleted_at IS NULL;

SELECT CONCAT('  followup_type 迁移记录数: ', COUNT(*)) AS info 
FROM `addons_customer_pro_clues_followup_type`;

SELECT 'Step 6 完成' AS info;

-- ============================================================
-- Step 7: 迁移 filter_group_ids 数据
-- ============================================================
SELECT 'Step 7: 迁移 filter_group_ids 数据...' AS info;

TRUNCATE TABLE `addons_customer_pro_clues_filter_group`;

INSERT IGNORE INTO `addons_customer_pro_clues_filter_group` (`id`, `clues_id`, `group_id`, `createTime`, `updateTime`)
SELECT 
    CONCAT(c.id, '_fg_', TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.filter_group_ids, ',', n.n), ',', -1))) AS id,
    c.id AS clues_id,
    TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.filter_group_ids, ',', n.n), ',', -1)) AS group_id,
    COALESCE(c.createTime, NOW(3)) AS createTime,
    NOW(3) AS updateTime
FROM `addons_customer_pro_clues` c
CROSS JOIN (
    SELECT 1 n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5
    UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
) n
WHERE c.filter_group_ids IS NOT NULL 
  AND c.filter_group_ids != '' 
  AND n.n <= 1 + LENGTH(c.filter_group_ids) - LENGTH(REPLACE(c.filter_group_ids, ',', ''))
  AND c.deleted_at IS NULL;

SELECT CONCAT('  filter_group_ids 迁移记录数: ', COUNT(*)) AS info 
FROM `addons_customer_pro_clues_filter_group`;

SELECT 'Step 7 完成' AS info;

-- ============================================================
-- Step 8: 迁移 level 数据（纯数字和JSON数组两种格式）
-- ============================================================
SELECT 'Step 8: 迁移 level 数据...' AS info;

TRUNCATE TABLE `addons_customer_pro_clues_level`;

-- 纯数字格式：level = '1' 或 '1,2,3'
INSERT IGNORE INTO `addons_customer_pro_clues_level` (`id`, `clues_id`, `level_value`, `createTime`, `updateTime`)
SELECT 
    CONCAT(c.id, '_lv_', TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.level, ',', n.n), ',', -1))) AS id,
    c.id AS clues_id,
    TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.level, ',', n.n), ',', -1)) AS level_value,
    COALESCE(c.createTime, NOW(3)) AS createTime,
    NOW(3) AS updateTime
FROM `addons_customer_pro_clues` c
CROSS JOIN (
    SELECT 1 n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5
) n
WHERE c.level IS NOT NULL 
  AND c.level != '' 
  AND c.level NOT LIKE '[%'           -- 排除JSON数组格式
  AND c.level NOT LIKE '%"%'          -- 排除包含引号的格式
  AND n.n <= 1 + LENGTH(c.level) - LENGTH(REPLACE(c.level, ',', ''))
  AND c.deleted_at IS NULL;

-- JSON数组格式：level = '["1","2"]' 或 '[1,2]'
INSERT IGNORE INTO `addons_customer_pro_clues_level` (`id`, `clues_id`, `level_value`, `createTime`, `updateTime`)
SELECT 
    CONCAT(c.id, '_lv_', TRIM(v.level_item)) AS id,
    c.id AS clues_id,
    TRIM(v.level_item) AS level_value,
    COALESCE(c.createTime, NOW(3)) AS createTime,
    NOW(3) AS updateTime
FROM `addons_customer_pro_clues` c,
JSON_TABLE(
    CASE 
        WHEN c.level LIKE '[%' THEN c.level 
        ELSE CONCAT('["', REPLACE(c.level, ',', '","'), '"]')
    END,
    '$[*]' COLUMNS (level_item VARCHAR(32) PATH '$')
) v
WHERE c.level IS NOT NULL 
  AND c.level != '' 
  AND c.level LIKE '[%'               -- 只处理JSON数组格式
  AND v.level_item IS NOT NULL
  AND v.level_item != ''
  AND c.deleted_at IS NULL;

SELECT CONCAT('  level 迁移记录数: ', COUNT(*)) AS info 
FROM `addons_customer_pro_clues_level`;

SELECT 'Step 8 完成' AS info;

-- ============================================================
-- Step 8.5: 迁移 services_ids 数据 → clues_service
-- ============================================================
SELECT 'Step 8.5: 迁移 services_ids 数据...' AS info;

TRUNCATE TABLE `addons_customer_pro_clues_service`;

INSERT IGNORE INTO `addons_customer_pro_clues_service` (`id`, `clues_id`, `user_id`, `createTime`, `updateTime`)
SELECT 
    CONCAT(c.id, '_sv_', TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.services_ids, ',', n.n), ',', -1))) AS id,
    c.id AS clues_id,
    TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.services_ids, ',', n.n), ',', -1)) AS user_id,
    COALESCE(c.createTime, NOW(3)) AS createTime,
    NOW(3) AS updateTime
FROM `addons_customer_pro_clues` c
CROSS JOIN (
    SELECT 1 n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5
    UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
    UNION SELECT 11 UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15
    UNION SELECT 16 UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20
) n
WHERE c.services_ids IS NOT NULL 
  AND c.services_ids != '' 
  AND c.services_ids != '0'
  AND n.n <= 1 + LENGTH(c.services_ids) - LENGTH(REPLACE(c.services_ids, ',', ''))
  AND c.deleted_at IS NULL;

SELECT CONCAT('  clues_service 迁移记录数: ', COUNT(*)) AS info 
FROM `addons_customer_pro_clues_service`;

SELECT 'Step 8.5 完成' AS info;

-- ============================================================
-- Step 9: 验证
-- ============================================================
SELECT 'Step 9: 验证...' AS info;

SELECT 
    t.`表名` AS `关联表`,
    COUNT(*) AS `记录数`,
    COUNT(DISTINCT clues_id) AS `覆盖线索数`
FROM (
    SELECT 'followup_type' AS `表名`, clues_id FROM addons_customer_pro_clues_followup_type
    UNION ALL
    SELECT 'filter_group', clues_id FROM addons_customer_pro_clues_filter_group
    UNION ALL
    SELECT 'level', clues_id FROM addons_customer_pro_clues_level
    UNION ALL
    SELECT 'clues_service', clues_id FROM addons_customer_pro_clues_service
) t
GROUP BY t.`表名`;

SELECT CONCAT('=== 线索查询优化部署完成: ', NOW(), ' ===') AS info;
SELECT '请部署新版 Go 代码后生效。旧版代码仍可正常运行。' AS info;
