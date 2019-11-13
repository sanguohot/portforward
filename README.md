# portforward
simple tcp/udp port forward to another address
## 环境依赖
go1.12+
## 编译
```
# git clone https://github.com/sanguohot/portforward
# cd portforward
# go build --ldflags "-linkmode external -extldflags -static" -o portforward.exe bin/app/app.go
# SET CGO_ENABLED=0
# SET GOOS=linux
# SET GOARCH=amd64
# go build --ldflags "-linkmode external -extldflags -static" -o portforward bin/app/app.go
```
## 静态文件需要系统包支持
```
# yum install glibc-static libstdc++-static -y
```