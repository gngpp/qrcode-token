# AliyunPanToken
根据浏览器抓包接口编写的一个从命令行获取阿里云盘`refresh_token`的工具，使用移动客户端APP `QRCode`扫码登录，获取的refresh_token，支持`alist`直链下载。

### 使用
- 编译安装（依赖cgo，确保存在gcc环境）
```shell
# 拉取源码
git clone https://github.com/zf1976/AliyunPanToken.git && cd AliyunPanToken

# 拉取依赖
go mod tidy

# 编译
go build -ldflags="-s -w" -o app main.go

# 执行
./app
```
- openwrt环境下需要安装gcc（固件一般不会自带）
```shell
opkg update && opkg install gcc
ar -rc /usr/lib/libpthread.a
```

### 示例

<img src="img/img1.png"/>
<img src="img/img2.png"/>
