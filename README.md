# AliyunPanToken
这是根据浏览器抓包接口编写的一个从命令行获取阿里云`refresh_token`的工具，使用移动客户端APP `QRCode`扫码登录，获取的refresh_token，支持`alist`直链下载。

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
