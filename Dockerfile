FROM scratch

WORKDIR /root/

# 将编译之后的可执行文件，数据库文件、配置文件 拷贝到 image 中
COPY ./snowflakes .
COPY ./conf ./conf

EXPOSE 8000

# 启动服务
CMD ["./snowflakes"]