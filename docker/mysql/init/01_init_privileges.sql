-- ----------------------------
-- Initialize database privileges
-- ----------------------------

-- Create canal user if not exists
CREATE USER IF NOT EXISTS 'canal'@'%' IDENTIFIED BY 'canal123';

-- Grant privileges to canal user for all databases
GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%';

-- Grant privileges to canal user for canal_manager database
GRANT ALL PRIVILEGES ON canal_manager.* TO 'canal'@'%';

-- Flush privileges to apply changes
FLUSH PRIVILEGES; 