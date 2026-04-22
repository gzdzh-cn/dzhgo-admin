#!/bin/bash

# 检查参数数量
if [ $# -ne 2 ]; then
    echo "用法: $0 <源路径> <目标路径>"
    echo "示例: $0 /Volumes/disk/site/go/dzhgo/dzhgo-admin/addons/customer_pro/ /Volumes/disk/site/go/dzhgo/extend/customer_pro/ "
    exit 1
fi

# 获取命令行参数
SOURCE_PATH="$1"
TARGET_PATH="$2"

# 检查源路径是否存在
if [ ! -d "$SOURCE_PATH" ]; then
    echo "错误: 源路径 '$SOURCE_PATH' 不存在"
    exit 1
fi

echo "开始同步..."
echo "源路径: $SOURCE_PATH"
echo "目标路径: $TARGET_PATH"

rsync -avh --delete \
    --exclude-from='rsync-exclude.txt' \
    --filter=':- .gitignore' \
    "$SOURCE_PATH" "$TARGET_PATH"

echo "同步完成!"
