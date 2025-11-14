FROM alpine:latest

# 设置工作目录
WORKDIR /home/admin/canal-adapter

# 安装必要的工具
RUN apk add --no-cache wget curl openjdk8-jre

# 下载Canal Adapter 1.1.5版本
RUN wget https://github.com/alibaba/canal/releases/download/canal-1.1.5/canal.adapter-1.1.5.tar.gz \
    && tar -xzf canal.adapter-1.1.5.tar.gz \
    && rm canal.adapter-1.1.5.tar.gz

# 创建必要的目录
RUN mkdir -p /home/admin/canal-adapter/conf \
    && mkdir -p /home/admin/canal-adapter/logs \
    && mkdir -p /home/admin/canal-adapter/plugin

# 设置环境变量
ENV CANAL_ADAPTER_HOME=/home/admin/canal-adapter
ENV PATH=$CANAL_ADAPTER_HOME/bin:$PATH

# 暴露端口
EXPOSE 8081

# 启动命令
CMD ["sh", "-c", "java -Xms256m -Xmx1024m -Dlogback.configurationFile=conf/logback.xml -classpath conf:lib/*:plugin/* com.alibaba.otter.canal.adapter.launcher.CanalAdapterApplication"] 