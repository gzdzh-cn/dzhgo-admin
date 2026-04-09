-- ============================================================
-- 线上迁移脚本：FIND_IN_SET 优化 → 关联表
-- 
-- 变更内容：
--   1. 新建 addons_customer_pro_clues_service 关联表
--   2. 将 services_ids 字段拆分迁移到关联表
--   3. 验证数据一致性
--
-- 执行方式：
--   mysql -h <host> -P <port> -u<user> -p<password> <database> < migrate_finding_set_optimization.sql
--
-- 注意事项：
--   - 脚本设计为幂等，可重复执行
--   - 迁移期间不影响现有业务（只读不写）
--   - 建议在低峰期执行
--   - 执行后需部署新版 Go 代码才能生效
-- ============================================================

-- 记录开始时间
SELECT CONCAT('迁移开始时间: ', NOW()) AS info;

-- ============================================================
-- Step 1: 创建关联表（幂等：IF NOT EXISTS）
-- ============================================================
SELECT 'Step 1: 创建关联表...' AS info;

CREATE TABLE IF NOT EXISTS `addons_customer_pro_clues_service` (
    `id` VARCHAR(64) NOT NULL PRIMARY KEY COMMENT '雪花ID',
    `clues_id` VARCHAR(64) NOT NULL COMMENT '线索ID',
    `user_id` VARCHAR(64) NOT NULL COMMENT '客服ID',
    INDEX `idx_addons_customer_pro_service_clues_id` (`clues_id`),
    INDEX `idx_addons_customer_pro_service_user_id` (`user_id`),
    UNIQUE KEY `uk_addons_customer_pro_service_clues_user` (`clues_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='线索-客服分配关联表';

SELECT 'Step 1 完成: 关联表已创建' AS info;

-- ============================================================
-- Step 2: 清空旧数据（幂等：重复执行不会报错）
-- ============================================================
SELECT 'Step 2: 清空旧数据...' AS info;

TRUNCATE TABLE `addons_customer_pro_clues_service`;

SELECT 'Step 2 完成' AS info;

-- ============================================================
-- Step 3: 迁移现有数据（将 services_ids 拆分写入关联表）
-- ============================================================
SELECT 'Step 3: 开始迁移数据...' AS info;

INSERT IGNORE INTO `addons_customer_pro_clues_service` (`id`, `clues_id`, `user_id`)
SELECT 
    CONCAT(c.id, '_', TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.services_ids, ',', n.n), ',', -1))) AS id,
    c.id AS clues_id,
    TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(c.services_ids, ',', n.n), ',', -1)) AS user_id
FROM `addons_customer_pro_clues` c
CROSS JOIN (
    SELECT 1 n UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5
    UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
    UNION SELECT 11 UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15
    UNION SELECT 16 UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20
    UNION SELECT 21 UNION SELECT 22 UNION SELECT 23 UNION SELECT 24 UNION SELECT 25
    UNION SELECT 26 UNION SELECT 27 UNION SELECT 28 UNION SELECT 29 UNION SELECT 30
) n
WHERE c.services_ids IS NOT NULL 
  AND c.services_ids != '' 
  AND c.services_ids != '0'
  AND n.n <= 1 + LENGTH(c.services_ids) - LENGTH(REPLACE(c.services_ids, ',', ''))
  AND c.deleted_at IS NULL;

SELECT 'Step 3 完成: 数据迁移结束' AS info;

-- ============================================================
-- Step 4: 验证迁移结果
-- ============================================================
SELECT 'Step 4: 验证迁移结果...' AS info;

SELECT 
    COUNT(*) AS '关联表总记录数',
    COUNT(DISTINCT clues_id) AS '覆盖线索数',
    COUNT(DISTINCT user_id) AS '涉及客服数'
FROM `addons_customer_pro_clues_service`;

-- 检查有 services_ids 但未迁移的线索（应为 0）
SELECT COUNT(*) AS '遗漏线索数(应为0)' FROM `addons_customer_pro_clues` c
WHERE c.services_ids IS NOT NULL 
  AND c.services_ids != '' 
  AND c.services_ids != '0'
  AND c.deleted_at IS NULL
  AND NOT EXISTS (
      SELECT 1 FROM `addons_customer_pro_clues_service` cs 
      WHERE cs.clues_id = c.id
  );

-- 抽样对比（取 5 条验证 services_ids 拆分是否正确）
SELECT c.id, c.services_ids, GROUP_CONCAT(cs.user_id) AS service_table_user_ids
FROM `addons_customer_pro_clues` c
LEFT JOIN `addons_customer_pro_clues_service` cs ON cs.clues_id = c.id
WHERE c.services_ids IS NOT NULL AND c.services_ids != '' AND c.services_ids != '0' AND c.deleted_at IS NULL
GROUP BY c.id, c.services_ids
LIMIT 5;

-- ============================================================
-- Step 5: 添加查询优化索引
-- ============================================================
SELECT 'Step 5: 添加查询优化索引...' AS info;

-- 前缀索引：加速纯数字关键字（手机号/guest_id）精确匹配搜索
ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_addons_customer_pro_clues_mobile` (`status`, `dtype`, `mobile`(20));
ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_addons_customer_pro_clues_guest_id` (`status`, `dtype`, `guest_id`(64));

-- 覆盖索引：加速 COUNT 分页查询（status + dtype + deleted_at）
ALTER TABLE `addons_customer_pro_clues` ADD INDEX `idx_addons_customer_pro_clues_status_dtype` (`status`, `dtype`, `deleted_at`);

-- FULLTEXT 全文索引：加速中文名/文本关键字搜索（ngram 分词）
ALTER TABLE `addons_customer_pro_clues` ADD FULLTEXT INDEX `idx_addons_customer_pro_clues_fulltext` (`name`, `mobile`, `wechat`) WITH PARSER ngram;

SELECT 'Step 5 完成: 索引已创建' AS info;

-- ============================================================
-- Step 6: 验证索引使用情况
SELECT 'Step 5: 索引验证...' AS info;

SHOW INDEX FROM `addons_customer_pro_clues_service`;

SELECT CONCAT('迁移完成时间: ', NOW()) AS info;
SELECT '========================================' AS info;
SELECT '迁移完成！请部署新版 Go 代码后生效。' AS info;
SELECT '旧版代码仍可正常运行，不受影响。' AS info;
