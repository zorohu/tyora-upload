# Typora OneIndex upload script (golang implement)

## 介绍

自动上传图片到 OneIndex 服务的 Typora 脚本
因为在macos上用typora执行python脚本失败，寻找原因无果，只好用golang实现了一个小工具QAQ
## 使用方法

### 1. clone

- 下载[本仓库](https://github.com/zorohu/tyora-upload)

### 2. 依赖下载 

`go mod tidy`

### 3. oneindex设置

![image-20210717132300365](http://pic.authorzero.com/img/image-20210717132300365.png)

### 4. 脚本配置

修改 `main.go` 中的图床地址

```
// 这里填你 OneIndex 图床服务的地址
const url = "https://www.oneindex.authorzero.com/images"
```
### 4. 编译打包
```
set GOARCH=amd64
set GOOS=linux
go build main.go
```
或者是

```
# 安装在你的gopath下面的bin目录
go install
```

- 注意需要给编译后的可执行文件权限 `chmod +x main`

### 5. Typora配置

![](http://pic.authorzero.com/img/截屏2021-07-17.png)


## 参考链接
https://www.bilibili.com/read/cv10599028
https://github.com/JeroGC/Typora-OneIndex-upload-script