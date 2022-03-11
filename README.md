# AliyunPanToken
这是一个从命令行获取阿里云refresh_token的工具，使用手机客户端QRCode扫码登录，获取的refresh_token可以在alist支持直链下载。

### 使用
- 编译安装
```shell
# 拉取源码
git clone https://github.com/zf1976/AliyunPanToken.git && cd AliyunPanToken

# 编译
go build -ldflags="-s -w" -o app main.go

# 执行
./app
```

### 示例

<img src="img/img1.png"/>
<img src="img/img2.png"/>
