# 打包

## linux下打包windows服务器的软件
```shell
 CGO_ENABLED=0 GOOS=windows GOARCH=amd64
```
```shell
go build main.go tcp.go udp.go  plugin.go tcp_linux.go -o main.exe
```

# 运行

## 服务器端
```shell
.\shadowsocks2-win64.exe -s 'ss://AEAD_CHACHA20_POLY1305:123456@:8488' -verbose
```
客户端加密方式是`CHACHA20-IETF-POLY1305`

## 客户端
```shell
-c 'ss://AEAD_CHACHA20_POLY1305:123456@[47.88.21.103]:8488' -verbose -socks :1080 -u
```
加上日志级别
```shell
-loglevel debug
```

# 日志设置
写入到本地磁盘

> https://www.jianshu.com/p/cfb7fb34bee5

# 绕过局域网及大陆