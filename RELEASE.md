# 远程桌面连接客户端 - 发布说明

## 打包和发布

### 1. 编译可执行文件

在安装了 Go 环境的 Windows 系统上：

```bash
go build -ldflags="-H windowsgui" -o RemoteDesktopClient.exe .
```

### 2. 依赖文件

编译后的可执行文件是独立的，不需要额外的依赖文件。

### 3. 发布包结构

```
RemoteDesktopClient/
├── RemoteDesktopClient.exe      # 主程序
├── README.md                    # 使用说明
├── LICENSE                      # 许可证
└── docs/                        # 文档（可选）
    ├── installation_guide.md    # 安装指南
    └── user_manual.md           # 用户手册
```

### 4. 发布版本

- 版本号: 1.0.0
- 发布日期: 2026-02-07
- 平台: Windows x64

### 5. 安装说明

1. 下载发布包
2. 解压到任意目录
3. 双击 RemoteDesktopClient.exe 运行程序
4. 确保目标计算机已启用远程桌面服务

### 6. 系统要求

- Windows 10/11 (64位)
- 已启用远程桌面服务的目标计算机
- 网络连通性

### 7. 常见问题

Q: 应用程序无法启动
A: 确保使用的是 Windows 系统，此应用程序仅支持 Windows 平台

Q: 无法连接到远程计算机
A: 检查IP地址、端口号是否正确，确认目标计算机已启用远程桌面服务并配置了防火墙

Q: 连接后显示黑屏或无响应
A: 检查网络连接，确认目标计算机处于活动状态