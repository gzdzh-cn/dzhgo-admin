FROM openjdk:8-jdk-alpine

# 设置工作目录
WORKDIR /home/admin/canal-server

# 安装必要的工具
RUN apk add --no-cache wget curl

# 下载Canal Server
RUN wget https://github.com/alibaba/canal/releases/download/canal-1.1.8/canal.deployer-1.1.8.tar.gz \
    && tar -xzf canal.deployer-1.1.8.tar.gz \
    && rm canal.deployer-1.1.8.tar.gz

# 创建必要的目录
RUN mkdir -p /home/admin/canal-server/conf \
    && mkdir -p /home/admin/canal-server/logs \
    && mkdir -p /home/admin/canal-server/plugin

# 设置环境变量
ENV CANAL_HOME=/home/admin/canal-server
ENV PATH=$CANAL_HOME/bin:$PATH

# 暴露端口
EXPOSE 11110 11111 11112

# 启动命令
CMD ["sh", "-c", "java -Xms256m -Xmx1024m -Dlogback.configurationFile=conf/logback.xml -classpath conf:lib/*:plugin/* com.alibaba.otter.canal.deployer.CanalLauncher"] 