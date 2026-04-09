#!/bin/bash

# 导出 dzhgo_go 数据库全部数据 - 菜单版本
# 使用方法：./database_tool.sh

# 设置变量
DB_NAME="dzhcrm_go"
DB_USER="root"
DB_PASSWORD="dzh123456"
CONTAINER_NAME="dzh_mysql57"
BACKUP_DIR="/home/backup"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# 在容器内创建备份目录
docker exec ${CONTAINER_NAME} mkdir -p ${BACKUP_DIR}

# 显示菜单
show_menu() {
    echo "=========================================="
    echo "        dzhgo_go 数据库导出工具"
    echo "=========================================="
    echo "1. 导出完整数据库（表结构 + 数据）"
    echo "2. 仅导出表结构"
    echo "3. 仅导出数据（不包含表结构）"
    echo "4. 导出特定表（支持多表）"
    echo "5. 查看所有表"
    echo "6. 导入 SQL 文件（支持gzip、zip、sql文件）"
    echo "7. 导入目录内所有SQL文件"
    echo "0. 退出"
    echo "=========================================="
    echo -n "请选择操作 (0-7): "
}

# 导出完整数据库
export_complete() {
    BACKUP_FILE="${BACKUP_DIR}/dzhgo_go_complete_${TIMESTAMP}.sql"
    echo "正在导出完整数据库到容器内 ${BACKUP_FILE}..."
    
    # 在容器内执行导出到BACKUP_DIR
    docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysqldump -u ${DB_USER} ${DB_NAME} > ${BACKUP_FILE}"
    
    if [ $? -eq 0 ]; then
        echo "✅ 完整数据库导出成功！"
        echo "📁 备份文件位置: 容器内 ${BACKUP_FILE}"
        echo "📊 文件大小: $(docker exec ${CONTAINER_NAME} du -h ${BACKUP_FILE} | cut -f1)"
        
        # 压缩备份文件
        echo "正在压缩备份文件..."
        docker exec ${CONTAINER_NAME} gzip ${BACKUP_FILE}
        echo "✅ 压缩完成: ${BACKUP_FILE}.gz"
        echo "📊 压缩后大小: $(docker exec ${CONTAINER_NAME} du -h ${BACKUP_FILE}.gz | cut -f1)"
    else
        echo "❌ 数据库导出失败！"
    fi
}

# 仅导出表结构
export_structure() {
    BACKUP_FILE="${BACKUP_DIR}/dzhgo_go_structure_${TIMESTAMP}.sql"
    echo "正在导出表结构到容器内 ${BACKUP_FILE}..."
    
    # 在容器内执行导出到BACKUP_DIR
    docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysqldump -u ${DB_USER} --no-data ${DB_NAME} > ${BACKUP_FILE}"
    
    if [ $? -eq 0 ]; then
        echo "✅ 表结构导出成功！"
        echo "📁 备份文件位置: 容器内 ${BACKUP_FILE}"
        echo "📊 文件大小: $(docker exec ${CONTAINER_NAME} du -h ${BACKUP_FILE} | cut -f1)"
    else
        echo "❌ 表结构导出失败！"
    fi
}

# 仅导出数据
export_data() {
    BACKUP_FILE="${BACKUP_DIR}/dzhgo_go_data_${TIMESTAMP}.sql"
    echo "正在导出数据到容器内 ${BACKUP_FILE}..."
    
    # 在容器内执行导出到BACKUP_DIR
    docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysqldump -u ${DB_USER} --no-create-info ${DB_NAME} > ${BACKUP_FILE}"
    
    if [ $? -eq 0 ]; then
        echo "✅ 数据导出成功！"
        echo "📁 备份文件位置: 容器内 ${BACKUP_FILE}"
        echo "📊 文件大小: $(docker exec ${CONTAINER_NAME} du -h ${BACKUP_FILE} | cut -f1)"
    else
        echo "❌ 数据导出失败！"
    fi
}

# 导出特定表
export_table() {
    echo "请输入要导出的表名（多个表用空格或逗号分隔）: "
    echo "例如: clues api"
    echo "或者: clues,api"
    echo -n "表名: "
    read TABLE_NAMES
    
    if [ -z "$TABLE_NAMES" ]; then
        echo "❌ 表名不能为空！"
        return
    fi
    
    # 处理表名，支持空格和逗号分隔
    # 先替换逗号为空格，然后处理多个空格
    TABLE_NAMES=$(echo "$TABLE_NAMES" | sed 's/,/ /g' | tr -s ' ')
    
    # 验证表是否存在
    echo "正在验证表是否存在..."
    INVALID_TABLES=""
    VALID_TABLES=""
    
    for table in $TABLE_NAMES; do
        # 检查表是否存在
        if docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} -e \"USE ${DB_NAME}; SHOW TABLES LIKE '${table}';\" | grep -q '${table}'"; then
            VALID_TABLES="$VALID_TABLES $table"
        else
            INVALID_TABLES="$INVALID_TABLES $table"
        fi
    done
    
    # 显示验证结果
    if [ -n "$INVALID_TABLES" ]; then
        echo "❌ 以下表不存在: $INVALID_TABLES"
    fi
    
    if [ -z "$VALID_TABLES" ]; then
        echo "❌ 没有有效的表可以导出！"
        return
    fi
    
    # 生成文件名（使用第一个表名作为基础）
    FIRST_TABLE=$(echo $VALID_TABLES | awk '{print $1}')
    # 生成文件名（使用每个表的最后单词拼接）
    if [ $(echo $VALID_TABLES | wc -w) -eq 1 ]; then
        # 单个表使用原表名
        BACKUP_FILE="${BACKUP_DIR}/${FIRST_TABLE}_${TIMESTAMP}.sql"
    else
        # 多个表使用最后单词拼接
        LAST_WORDS=""
        for table in $VALID_TABLES; do
            # 获取表的最后单词（以_分隔）
            last_word=$(echo "$table" | sed 's/.*_//')
            if [ -z "$last_word" ]; then
                # 如果没有_分隔符，使用整个表名
                last_word="$table"
            fi
            LAST_WORDS="${LAST_WORDS}_${last_word}"
        done
        # 去掉开头的下划线
        LAST_WORDS=$(echo "$LAST_WORDS" | sed 's/^_//')
        BACKUP_FILE="${BACKUP_DIR}/${LAST_WORDS}_${TIMESTAMP}.sql"
    fi
    
    echo "✅ 将导出以下表: $VALID_TABLES"
    echo "正在导出表到容器内 ${BACKUP_FILE}..."
    
    # 构建mysqldump命令
    DUMP_CMD="MYSQL_PWD=${DB_PASSWORD} mysqldump -u ${DB_USER} ${DB_NAME} $VALID_TABLES"
    
    # 在容器内执行导出到BACKUP_DIR
    docker exec ${CONTAINER_NAME} bash -c "$DUMP_CMD > ${BACKUP_FILE}"
    
    if [ $? -eq 0 ]; then
        echo "✅ 表导出成功！"
        echo "📁 备份文件位置: 容器内 ${BACKUP_FILE}"
        echo "📊 文件大小: $(docker exec ${CONTAINER_NAME} du -h ${BACKUP_FILE} | cut -f1)"
        
        # 显示导出的表列表
        echo "📋 导出的表:"
        for table in $VALID_TABLES; do
            echo "   - $table"
        done
    else
        echo "❌ 表导出失败！"
    fi
}

# 查看所有表
show_tables() {
    echo "数据库 ${DB_NAME} 中的所有表："
    echo "=========================================="
    docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} -e \"SHOW TABLES FROM ${DB_NAME};\""
    echo "=========================================="
}

# 恢复数据


# 导入目录内所有SQL文件
import_directory() {
    echo "请选择导入方式："
    echo "1. 从容器内 ${BACKUP_DIR} 目录导入到默认数据库"
    echo "2. 自定义导入（指定数据库和目录）"
    echo "0. 返回主菜单"
    echo -n "请选择 (0-2): "
    read import_choice
    
    case $import_choice in
        1)
            # 从容器内导入到默认数据库
            echo -n "请输入容器内 ${BACKUP_DIR} 目录中的文件夹名: "
            read SQL_DIR
            
            # 检查容器内目录是否存在
            if ! docker exec ${CONTAINER_NAME} test -d "${BACKUP_DIR}/${SQL_DIR}"; then
                echo "❌ 目录不存在: 容器内 ${BACKUP_DIR}/${SQL_DIR}"
                return
            fi
            
            echo "正在导入目录内所有SQL文件到数据库 ${DB_NAME}..."
            
            # 查找目录内的所有SQL文件
            SQL_FILES=$(docker exec ${CONTAINER_NAME} bash -c "find ${BACKUP_DIR}/${SQL_DIR} -name '*.sql' -type f | sort")
            if [ -n "$SQL_FILES" ]; then
                echo "找到SQL文件: $SQL_FILES"
                echo "开始顺序导入..."
                
                SUCCESS_COUNT=0
                TOTAL_COUNT=$(echo "$SQL_FILES" | wc -l)
                
                for sql_file in $SQL_FILES; do
                    echo "正在导入: $(basename "$sql_file")"
                    if docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME} < ${sql_file}"; then
                        echo "✅ 导入成功: $(basename "$sql_file")"
                        SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
                    else
                        echo "❌ 导入失败: $(basename "$sql_file")"
                    fi
                done
                
                echo "导入完成！成功: $SUCCESS_COUNT/$TOTAL_COUNT"
                if [ $SUCCESS_COUNT -eq $TOTAL_COUNT ]; then
                    echo "✅ 所有SQL文件导入成功！"
                else
                    echo "⚠️  部分SQL文件导入失败"
                fi
            else
                echo "❌ 目录内未找到SQL文件"
                return 1
            fi
            ;;
        2)
            # 自定义导入
            echo "=== 自定义导入设置 ==="
            echo -n "请输入目标数据库名称: "
            read TARGET_DB
            
            if [ -z "$TARGET_DB" ]; then
                echo "❌ 数据库名称不能为空！"
                return
            fi
            
            echo -n "请输入容器内的目录路径: "
            read SQL_DIR
            
            # 检查容器内目录是否存在
            if ! docker exec ${CONTAINER_NAME} test -d "${SQL_DIR}"; then
                echo "❌ 目录不存在: 容器内 ${SQL_DIR}"
                return
            fi
            
            echo "正在导入目录内所有SQL文件到数据库 ${TARGET_DB}..."
            
            # 查找目录内的所有SQL文件
            SQL_FILES=$(docker exec ${CONTAINER_NAME} bash -c "find ${SQL_DIR} -name '*.sql' -type f | sort")
            if [ -n "$SQL_FILES" ]; then
                echo "找到SQL文件: $SQL_FILES"
                echo "开始顺序导入..."
                
                SUCCESS_COUNT=0
                TOTAL_COUNT=$(echo "$SQL_FILES" | wc -l)
                
                for sql_file in $SQL_FILES; do
                    echo "正在导入: $(basename "$sql_file")"
                    if docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB} < ${sql_file}"; then
                        echo "✅ 导入成功: $(basename "$sql_file")"
                        SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
                    else
                        echo "❌ 导入失败: $(basename "$sql_file")"
                    fi
                done
                
                echo "导入完成！成功: $SUCCESS_COUNT/$TOTAL_COUNT"
                if [ $SUCCESS_COUNT -eq $TOTAL_COUNT ]; then
                    echo "✅ 所有SQL文件导入成功！"
                else
                    echo "⚠️  部分SQL文件导入失败"
                fi
            else
                echo "❌ 目录内未找到SQL文件"
                return 1
            fi
            ;;
        0)
            return
            ;;
        *)
            echo "❌ 无效选择！"
            ;;
    esac
}

# 导入 SQL 文件
import_sql() {
    echo "请选择导入方式："
    echo "1. 从宿主机文件导入到默认数据库"
    echo "2. 从容器内 ${BACKUP_DIR} 目录导入到默认数据库"
    echo "3. 自定义导入（指定数据库和文件）"
    echo "0. 返回主菜单"
    echo -n "请选择 (0-3): "
    read import_choice
    
    case $import_choice in
        1)
            # 从宿主机导入到默认数据库
            echo -n "请输入宿主机上的 SQL 文件路径: "
            read SQL_FILE
            
            if [ ! -f "$SQL_FILE" ]; then
                echo "❌ 文件不存在: ${SQL_FILE}"
                return
            fi
            
            echo "正在导入 SQL 文件到数据库 ${DB_NAME}..."
            
            if [[ "$SQL_FILE" == *.gz ]]; then
                # 如果是gzip压缩文件，先解压再导入
                gunzip -c "$SQL_FILE" | docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME}"
                                elif [[ "$SQL_FILE" == *.zip ]]; then
                        # 如果是zip文件，先解压再导入
                        echo "检测到zip文件，正在解压..."
                        TEMP_DIR=$(mktemp -d)
                        # 检查宿主机是否有unzip
                        if ! command -v unzip &> /dev/null; then
                            echo "❌ 宿主机缺少unzip命令，请先安装: brew install unzip (macOS) 或 apt-get install unzip (Linux)"
                            rm -rf "$TEMP_DIR"
                            return 1
                        fi
                        unzip -o "$SQL_FILE" -d "$TEMP_DIR"
                
                # 从zip文件名提取基础名称，用于匹配SQL文件
                BASE_NAME=$(echo "$(basename "$SQL_FILE")" | sed 's/\.zip$//' | sed 's/_[0-9]\{4\}-[0-9]\{2\}-[0-9]\{2\}_[0-9]\{2\}-[0-9]\{2\}-[0-9]\{2\}_mysql_data$//')
                echo "查找匹配的SQL文件，基础名称: $BASE_NAME"
                
                # 查找解压后的sql文件，优先匹配基础名称
                SQL_FILES=$(find "$TEMP_DIR" -name "*.sql" -type f)
                if [ -n "$SQL_FILES" ]; then
                    echo "找到SQL文件: $SQL_FILES"
                    
                                                # 优先导入匹配基础名称的文件
                            MATCHED_FILE=""
                            for sql_file in $SQL_FILES; do
                                sql_name=$(basename "$sql_file")
                                # 尝试多种匹配方式
                                if [[ "$sql_name" == *"$BASE_NAME"* ]] || \
                                   [[ "$sql_name" == *"dzhgo_go"* ]] || \
                                   [[ "$sql_name" == *"dzh3136"* ]]; then
                                    MATCHED_FILE="$sql_file"
                                    echo "✅ 找到匹配的SQL文件: $sql_name"
                                    break
                                fi
                            done
                    
                                                # 如果找到匹配的文件，只导入它
                            if [ -n "$MATCHED_FILE" ]; then
                                echo "正在导入匹配的SQL文件: $MATCHED_FILE"
                                docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME}" < "$MATCHED_FILE"
                            else
                                echo "❌ 未找到匹配的SQL文件！"
                                echo "期望找到包含 '$ZIP_BASE_NAME' 的SQL文件"
                                echo "实际找到的SQL文件: $SQL_FILES"
                                echo "请检查zip文件内容或重命名文件"
                                rm -rf "$TEMP_DIR"
                                return 1
                            fi
                else
                    echo "❌ 解压后未找到SQL文件！"
                    rm -rf "$TEMP_DIR"
                    return 1
                fi
                rm -rf "$TEMP_DIR"
            else
                # 直接导入
                docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME}" < "$SQL_FILE"
            fi
            
            if [ $? -eq 0 ]; then
                echo "✅ SQL 文件导入成功！"
            else
                echo "❌ SQL 文件导入失败！"
            fi
            ;;
        2)
            # 从容器内导入到默认数据库
            echo -n "请输入容器内 SQL 文件路径 (绝对路径或 ${BACKUP_DIR} 下的相对路径): "
            read SQL_FILE
            
            # 如果是绝对路径直接使用，否则拼接BACKUP_DIR
            if [[ "$SQL_FILE" == /* ]]; then
                FULL_SQL_PATH="${SQL_FILE}"
            else
                FULL_SQL_PATH="${BACKUP_DIR}/${SQL_FILE}"
            fi
            
            # 检查容器内文件是否存在
            if ! docker exec ${CONTAINER_NAME} test -f "${FULL_SQL_PATH}"; then
                echo "❌ 文件不存在: 容器内 ${FULL_SQL_PATH}"
                return
            fi
            
            echo "正在导入 SQL 文件到数据库 ${DB_NAME}..."
            
            if [[ "$SQL_FILE" == *.gz ]]; then
                # 如果是gzip压缩文件，先解压再导入
                docker exec ${CONTAINER_NAME} bash -c "gunzip -c ${FULL_SQL_PATH} | MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME}"
            elif [[ "$SQL_FILE" == *.zip ]]; then
                # 如果是zip文件，先解压再导入
                echo "检测到zip文件，正在解压..."
                # 先检查并安装unzip
                echo "检查容器内unzip命令..."
                if ! docker exec ${CONTAINER_NAME} bash -c "command -v unzip 2>/dev/null || which unzip 2>/dev/null"; then
                    echo "容器内缺少unzip命令，尝试安装..."
                    # 尝试不同的包管理器
                    if docker exec ${CONTAINER_NAME} bash -c "command -v apt-get 2>/dev/null"; then
                        echo "使用apt-get安装unzip..."
                        docker exec ${CONTAINER_NAME} bash -c "apt-get update && apt-get install -y unzip"
                    elif docker exec ${CONTAINER_NAME} bash -c "command -v yum 2>/dev/null"; then
                        echo "使用yum安装unzip..."
                        docker exec ${CONTAINER_NAME} bash -c "yum install -y unzip"
                    elif docker exec ${CONTAINER_NAME} bash -c "command -v apk 2>/dev/null"; then
                        echo "使用apk安装unzip..."
                        docker exec ${CONTAINER_NAME} bash -c "apk add unzip"
                    else
                        echo "❌ 无法识别包管理器，请手动安装unzip后重试"
                        return 1
                    fi
                    
                    # 等待安装完成并验证
                    echo "等待安装完成..."
                    sleep 3
                    if ! docker exec ${CONTAINER_NAME} bash -c "command -v unzip 2>/dev/null || which unzip 2>/dev/null"; then
                        echo "❌ unzip安装失败，请手动安装后重试"
                        return 1
                    else
                        echo "✅ unzip安装成功"
                    fi
                else
                    echo "容器内已有unzip命令"
                fi
                docker exec ${CONTAINER_NAME} bash -c "cd ${BACKUP_DIR} && unzip -o ${SQL_FILE}"
                
                # 从zip文件名提取基础名称，用于匹配SQL文件
                ZIP_BASE_NAME=$(echo "$SQL_FILE" | sed 's/\.zip$//')
                echo "查找匹配的SQL文件，zip基础名称: $ZIP_BASE_NAME"
                
                # 查找解压后的sql文件，优先匹配zip文件名
                SQL_FILES=$(docker exec ${CONTAINER_NAME} bash -c "cd ${BACKUP_DIR} && find . -name '*.sql' -type f")
                if [ -n "$SQL_FILES" ]; then
                    echo "找到SQL文件: $SQL_FILES"
                    
                    # 优先导入匹配zip文件名的SQL文件
                    MATCHED_FILE=""
                    for sql_file in $SQL_FILES; do
                        sql_name=$(basename "$sql_file")
                        # 检查SQL文件名是否包含zip文件名的主要部分
                        if [[ "$sql_name" == *"$ZIP_BASE_NAME"* ]]; then
                            MATCHED_FILE="$sql_file"
                            echo "✅ 找到匹配的SQL文件: $sql_name"
                            break
                        fi
                    done
                    
                    # 如果找到匹配的文件，只导入它
                    if [ -n "$MATCHED_FILE" ]; then
                        echo "正在导入匹配的SQL文件: $MATCHED_FILE"
                        docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME} < $(dirname ${FULL_SQL_PATH})/${MATCHED_FILE#./}"
                    else
                        echo "❌ 未找到匹配的SQL文件！"
                        echo "期望找到包含 '$ZIP_BASE_NAME' 的SQL文件"
                        echo "实际找到的SQL文件: $SQL_FILES"
                        echo "请检查zip文件内容或重命名文件"
                        return 1
                    fi
                else
                    echo "❌ 解压后未找到SQL文件！"
                    return 1
                fi
            else
                # 直接导入
                docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${DB_NAME} < ${FULL_SQL_PATH}"
            fi
            
            if [ $? -eq 0 ]; then
                echo "✅ SQL 文件导入成功！"
            else
                echo "❌ SQL 文件导入失败！"
            fi
            ;;
        3)
            # 自定义导入
            echo "=== 自定义导入设置 ==="
            echo -n "请输入目标数据库名称: "
            read TARGET_DB
            
            if [ -z "$TARGET_DB" ]; then
                echo "❌ 数据库名称不能为空！"
                return
            fi
            
            echo "请选择文件来源："
            echo "1. 宿主机文件"
            echo "2. 容器内文件"
            echo -n "请选择 (1-2): "
            read file_source
            
            case $file_source in
                1)
                    # 从宿主机导入
                    echo -n "请输入宿主机上的 SQL 文件路径: "
                    read SQL_FILE
                    
                    if [ ! -f "$SQL_FILE" ]; then
                        echo "❌ 文件不存在: ${SQL_FILE}"
                        return
                    fi
                    
                    echo "正在导入 SQL 文件到数据库 ${TARGET_DB}..."
                    
                    if [[ "$SQL_FILE" == *.gz ]]; then
                        # 如果是gzip压缩文件，先解压再导入
                        gunzip -c "$SQL_FILE" | docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB}"
                    elif [[ "$SQL_FILE" == *.zip ]]; then
                        # 如果是zip文件，先解压再导入
                        echo "检测到zip文件，正在解压..."
                        TEMP_DIR=$(mktemp -d)
                        # 检查宿主机是否有unzip
                        if ! command -v unzip &> /dev/null; then
                            echo "❌ 宿主机缺少unzip命令，请先安装: brew install unzip (macOS) 或 apt-get install unzip (Linux)"
                            rm -rf "$TEMP_DIR"
                            return 1
                        fi
                        unzip -o "$SQL_FILE" -d "$TEMP_DIR"
                        
                        # 从zip文件名提取基础名称，用于匹配SQL文件
                        ZIP_BASE_NAME=$(echo "$(basename "$SQL_FILE")" | sed 's/\.zip$//')
                        echo "查找匹配的SQL文件，zip基础名称: $ZIP_BASE_NAME"
                        
                        # 查找解压后的sql文件，优先匹配zip文件名
                        SQL_FILES=$(find "$TEMP_DIR" -name "*.sql" -type f)
                        if [ -n "$SQL_FILES" ]; then
                            echo "找到SQL文件: $SQL_FILES"
                            
                            # 优先导入匹配zip文件名的SQL文件
                            MATCHED_FILE=""
                            for sql_file in $SQL_FILES; do
                                sql_name=$(basename "$sql_file")
                                # 检查SQL文件名是否包含zip文件名的主要部分
                                if [[ "$sql_name" == *"$ZIP_BASE_NAME"* ]]; then
                                    MATCHED_FILE="$sql_file"
                                    echo "✅ 找到匹配的SQL文件: $sql_name"
                                    break
                                fi
                            done
                            
                            # 如果找到匹配的文件，只导入它
                            if [ -n "$MATCHED_FILE" ]; then
                                echo "正在导入匹配的SQL文件: $MATCHED_FILE"
                                docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB}" < "$MATCHED_FILE"
                            else
                                echo "❌ 未找到匹配的SQL文件！"
                                echo "期望找到包含 '$ZIP_BASE_NAME' 的SQL文件"
                                echo "实际找到的SQL文件: $SQL_FILES"
                                echo "请检查zip文件内容或重命名文件"
                                rm -rf "$TEMP_DIR"
                                return 1
                            fi
                        else
                            echo "❌ 解压后未找到SQL文件！"
                            rm -rf "$TEMP_DIR"
                            return 1
                        fi
                        rm -rf "$TEMP_DIR"
                    else
                        # 直接导入
                        docker exec -i ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB}" < "$SQL_FILE"
                    fi
                    ;;
                2)
                    # 从容器内导入
                    echo -n "请输入容器内的 SQL 文件路径: "
                    read SQL_FILE
                    
                    # 检查容器内文件是否存在
                    if ! docker exec ${CONTAINER_NAME} test -f "${SQL_FILE}"; then
                        echo "❌ 文件不存在: 容器内 ${SQL_FILE}"
                        return
                    fi
                    
                    echo "正在导入 SQL 文件到数据库 ${TARGET_DB}..."
                    
                    if [[ "$SQL_FILE" == *.gz ]]; then
                        # 如果是gzip压缩文件，先解压再导入
                        docker exec ${CONTAINER_NAME} bash -c "gunzip -c ${SQL_FILE} | MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB}"
                    elif [[ "$SQL_FILE" == *.zip ]]; then
                        # 如果是zip文件，先解压再导入
                        echo "检测到zip文件，正在解压..."
                        # 先检查并安装unzip
                        echo "检查容器内unzip命令..."
                        if ! docker exec ${CONTAINER_NAME} bash -c "command -v unzip 2>/dev/null || which unzip 2>/dev/null"; then
                            echo "容器内缺少unzip命令，尝试安装..."
                            # 尝试不同的包管理器
                            if docker exec ${CONTAINER_NAME} bash -c "command -v apt-get 2>/dev/null"; then
                                echo "使用apt-get安装unzip..."
                                docker exec ${CONTAINER_NAME} bash -c "apt-get update && apt-get install -y unzip"
                            elif docker exec ${CONTAINER_NAME} bash -c "command -v yum 2>/dev/null"; then
                                echo "使用yum安装unzip..."
                                docker exec ${CONTAINER_NAME} bash -c "yum install -y unzip"
                            elif docker exec ${CONTAINER_NAME} bash -c "command -v apk 2>/dev/null"; then
                                echo "使用apk安装unzip..."
                                docker exec ${CONTAINER_NAME} bash -c "apk add unzip"
                            else
                                echo "❌ 无法识别包管理器，请手动安装unzip后重试"
                                return 1
                            fi
                            
                            # 等待安装完成并验证
                            echo "等待安装完成..."
                            sleep 3
                            if ! docker exec ${CONTAINER_NAME} bash -c "command -v unzip 2>/dev/null || which unzip 2>/dev/null"; then
                                echo "❌ unzip安装失败，请手动安装后重试"
                                return 1
                            else
                                echo "✅ unzip安装成功"
                            fi
                        else
                            echo "容器内已有unzip命令"
                        fi
                        docker exec ${CONTAINER_NAME} bash -c "cd $(dirname ${SQL_FILE}) && unzip -o $(basename ${SQL_FILE})"
                        
                        # 从zip文件名提取基础名称，用于匹配SQL文件
                        ZIP_BASE_NAME=$(echo "$(basename "$SQL_FILE")" | sed 's/\.zip$//')
                        echo "查找匹配的SQL文件，zip基础名称: $ZIP_BASE_NAME"
                        
                        # 查找解压后的sql文件，优先匹配zip文件名
                        SQL_FILES=$(docker exec ${CONTAINER_NAME} bash -c "cd $(dirname ${SQL_FILE}) && find . -name '*.sql' -type f")
                        if [ -n "$SQL_FILES" ]; then
                            echo "找到SQL文件: $SQL_FILES"
                            
                            # 优先导入匹配zip文件名的SQL文件
                            MATCHED_FILE=""
                            for sql_file in $SQL_FILES; do
                                sql_name=$(basename "$sql_file")
                                # 检查SQL文件名是否包含zip文件名的主要部分
                                if [[ "$sql_name" == *"$ZIP_BASE_NAME"* ]]; then
                                    MATCHED_FILE="$sql_file"
                                    echo "✅ 找到匹配的SQL文件: $sql_name"
                                    break
                                fi
                            done
                            
                            # 如果找到匹配的文件，只导入它
                            if [ -n "$MATCHED_FILE" ]; then
                                echo "正在导入匹配的SQL文件: $MATCHED_FILE"
                                docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB} < $(dirname ${SQL_FILE})/${MATCHED_FILE#./}"
                            else
                                echo "❌ 未找到匹配的SQL文件！"
                                echo "期望找到包含 '$ZIP_BASE_NAME' 的SQL文件"
                                echo "实际找到的SQL文件: $SQL_FILES"
                                echo "请检查zip文件内容或重命名文件"
                                return 1
                            fi
                        else
                            echo "❌ 解压后未找到SQL文件！"
                            return 1
                        fi
                    else
                        # 直接导入
                        docker exec ${CONTAINER_NAME} bash -c "MYSQL_PWD=${DB_PASSWORD} mysql -u ${DB_USER} ${TARGET_DB} < ${SQL_FILE}"
                    fi
                    ;;
                *)
                    echo "❌ 无效选择！"
                    return
                    ;;
            esac
            
            if [ $? -eq 0 ]; then
                echo "✅ SQL 文件导入成功到数据库 ${TARGET_DB}！"
            else
                echo "❌ SQL 文件导入失败！"
            fi
            ;;
        0)
            return
            ;;
        *)
            echo "❌ 无效选择！"
            ;;
    esac
}

# 主循环
while true; do
    show_menu
    read choice
    
            case $choice in
        1)
            export_complete
            ;;
        2)
            export_structure
            ;;
        3)
            export_data
            ;;
        4)
            export_table
            ;;
        5)
            show_tables
            ;;

        6)
            import_sql
            ;;
        7)
            import_directory
            ;;
        0)
            echo "退出程序"
            exit 0
            ;;
        *)
            echo "❌ 无效选择，请重新输入！"
            ;;
    esac
    
    echo ""
    echo "按回车键继续..."
    read
done 