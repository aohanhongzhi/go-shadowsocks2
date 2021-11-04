# go开发环境配置代理
```shell
GOPROXY=https://goproxy.cn,direct
```

```
go mod download
```

![img_1.png](img_1.png)


# 打包

## linux下打包windows服务器的软件
```shell
 CGO_ENABLED=0 GOOS=windows GOARCH=amd64
```
```shell
go build main.go tcp.go udp.go  plugin.go tcp_linux.go -o main.exe
```
## windwos下打包
运行
```
go run .\
```
打包
```
 go build .\
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

### deepin配置socks5代理

![](deepin-proxy-socks.png)


# 日志设置
写入到本地磁盘，目前已经切换到logrus了。

> https://www.jianshu.com/p/cfb7fb34bee5

# 绕过局域网及大陆


# windows下异常
![img.png](img.png)
golang.org/x/sys/windows.WSAECONNABORTED (10053)
![](./windows下socks代理设置错误.png)
## 解决办法
可能是系统不支持，但是firefox支持。
![](./firefox.png)
![](./firefox-proxy.png)
![](./firefox-proxy-google.png)