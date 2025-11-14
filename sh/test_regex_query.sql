-- 测试正则表达式查询，验证85能匹配到85但不会匹配到185
-- 测试数据
-- services_ids = '85'        -- 应该匹配
-- services_ids = '85,648'    -- 应该匹配  
-- services_ids = '648,85'    -- 应该匹配
-- services_ids = '648,85,8576' -- 应该匹配
-- services_ids = '185'       -- 不应该匹配
-- services_ids = '185,648'   -- 不应该匹配
-- services_ids = '648,185'   -- 不应该匹配
-- services_ids = '285'       -- 不应该匹配

-- 测试查询
SELECT 
    services_ids,
    CASE 
        WHEN services_ids REGEXP '^85$|^85,|,85$|,85,' THEN '匹配85'
        ELSE '不匹配85'
    END as match_result
FROM (
    SELECT '85' as services_ids
    UNION ALL SELECT '85,648'
    UNION ALL SELECT '648,85'
    UNION ALL SELECT '648,85,8576'
    UNION ALL SELECT '185'
    UNION ALL SELECT '185,648'
    UNION ALL SELECT '648,185'
    UNION ALL SELECT '285'
) test_data
ORDER BY services_ids;

-- 实际查询示例（替换原来的FIND_IN_SET）
-- 原来的查询：
-- SELECT COUNT(1) FROM `addons_customer_pro_clues` AS clues WHERE
-- `clues`.`status`='0' AND FIND_IN_SET(85,clues.services_ids) AND `deleted_at` IS NULL

-- 修改后的查询：
-- SELECT COUNT(1) FROM `addons_customer_pro_clues` AS clues WHERE
-- `clues`.`status`='0' AND clues.services_ids REGEXP '^85$|^85,|,85$|,85,' AND `deleted_at` IS NULL 