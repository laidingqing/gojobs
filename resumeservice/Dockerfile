FROM iron/base
EXPOSE 6767

ADD resumeservice-linux-amd64 /

#TODO 入参环境配置相关, configServerUrl使用环境配置
ENTRYPOINT ["./resumeservice-linux-amd64", "--configServerUrl=http://192.168.0.9:8888", "--profile=dev"]