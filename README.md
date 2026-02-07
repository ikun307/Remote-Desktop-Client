# 远程桌面连接客户端

这是一个用 Go 语言开发的远程桌面连接客户端，支持通过指定 IP 地址和端口连接到远程计算机。

## 功能特性

- 通过 IP 地址和端口连接远程桌面
- 支持自定义端口（默认为 3389）
- 端口连通性测试
- Windows 原生界面
- 两种运行模式：命令行和GUI

## 技术架构

- 远程桌面连接: 使用 Windows 自带的 mstsc 命令
- 网络连通性测试: 使用 Go 内置的 net 包
- GUI 界面: 使用 Windows API (syscall)

## 注意事项

- 此软件仅支持 Windows 系统
- 目标计算机必须启用远程桌面服务
- 需要正确的凭据才能连接到远程计算机
- 防火墙必须允许远程桌面连接（默认端口 3389）
- 在连接前会自动测试端口连通性

## 编译源码

1. 确保已安装 Go 语言环境 (1.19 或更高版本)
2. 在项目目录中运行：
   ```
   go build -o RemoteDesktopClient.exe .
   ```
   或构建GUI版本：
   ```
   go build -ldflags="-H windowsgui" -o RemoteDesktopClientGUI.exe .
