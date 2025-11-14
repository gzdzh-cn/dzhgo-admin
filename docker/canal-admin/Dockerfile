FROM openjdk:8-jdk-alpine

# 设置工作目录
WORKDIR /home/admin/canal-admin

# 安装必要的工具
RUN apk add --no-cache wget curl

# 下载Canal Admin
RUN wget https://github.com/alibaba/canal/releases/download/canal-1.1.8/canal.admin-1.1.8.tar.gz \
    && tar -xzf canal.admin-1.1.8.tar.gz \
    && rm canal.admin-1.1.8.tar.gz \
    && wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/8.0.28/mysql-connector-java-8.0.28.jar \
    && mv mysql-connector-java-8.0.28.jar /home/admin/canal-admin/lib/

# 创建必要的目录
RUN mkdir -p /home/admin/canal-admin/conf \
    && mkdir -p /home/admin/canal-admin/logs

# 设置环境变量
ENV CANAL_ADMIN_HOME=/home/admin/canal-admin
ENV PATH=$CANAL_ADMIN_HOME/bin:$PATH

# 暴露端口
EXPOSE 8089

# 启动命令
CMD ["sh", "-c", "java -Xms256m -Xmx512m -Dcanal.conf.dir=/home/admin/canal-admin/conf -jar /home/admin/canal-admin/lib/canal.admin-web-1.1.8.jar"] 