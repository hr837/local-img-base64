# Local Image Base64 Converter

一个轻量级HTTP服务，用于将本地图片文件转换为Base64编码字符串，支持跨域访问和多平台部署。

## ✨ 功能特点
- `/` 测试页面
- `/getImageFile?file=xxx`接口接收图片路径参数，返回Base64编码结果。查询参数`file`：图片文件路径通过`encodeURIComponent`编码后（必填）

- 自动验证图片文件格式（支持.jpg/.jpeg/.png/.gif/.bmp/.webp/.svg）
- 跨域资源共享(CORS)支持
- 多平台编译支持（Windows/Linux/macOS）
- 监听端口：21333

## 🚀 快速开始

### 直接运行（预编译版本）
从`release`目录选择对应平台的可执行文件：
```bash
# Windows
c:\path\to\release\img-base64-windows-amd64.exe

# Linux
chmod +x img-base64-linux-amd64
./img-base64-linux-amd64

# macOS
chmod +x img-base64-darwin-amd64
./img-base64-darwin-amd64
```

### 源码编译
```bash
# 克隆或进入项目目录
cd /xx/img-base64

# 安装依赖
go mod tidy

# 编译当前Windows平台可执行文件
go build -o img-base64.exe

# 跨平台编译（Windows命令行）
set GOOS=linux&&set GOARCH=amd64&&go build -o release/img-base64-linux-amd64
set GOOS=darwin&&set GOARCH=amd64&&go build -o release/img-base64-darwin-amd64
```

## 🔍 使用示例

### 访问测试页面
启动服务后，在浏览器访问：http://localhost:21333


## 📄 许可证
[MIT](LICENSE) © Huang Rui