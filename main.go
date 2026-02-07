package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
	"syscall"
	"unsafe"
)

// RDPConfig 远程桌面配置
type RDPConfig struct {
	ID          string
	Name        string
	IPAddress   string
	Port        int
	Username    string
	Domain      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	MessageBoxW          = user32.NewProc("MessageBoxW")
	GetConsoleWindow     = user32.NewProc("GetConsoleWindow")
	ShowWindow           = user32.NewProc("ShowWindow")
	
	// Windows 窗口显示常量
	SW_HIDE int32 = 0
)

// showMessageBox 显示消息框
func showMessageBox(title, text string) {
	var (
		hwnd uintptr = 0
		caption *uint16 = syscall.StringToUTF16Ptr(title)
		message *uint16 = syscall.StringToUTF16Ptr(text)
		uType uint32 = 0x00000040 // MB_ICONINFORMATION
	)
	MessageBoxW.Call(hwnd, uintptr(unsafe.Pointer(message)), uintptr(unsafe.Pointer(caption)), uintptr(uType))
}

// TestPortConnectivity 测试端口连通性
func TestPortConnectivity(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return false
	}
	
	defer conn.Close()
	return true
}

// ConnectToRDP 连接到远程桌面
func ConnectToRDP(config RDPConfig) error {
	// 验证当前操作系统是否支持RDP
	if runtime.GOOS != "windows" {
		return fmt.Errorf("RDP only supported on Windows")
	}

	// 测试端口连通性
	if !TestPortConnectivity(config.IPAddress, config.Port) {
		return fmt.Errorf("无法连接到 %s:%d", config.IPAddress, config.Port)
	}

	// 使用mstsc命令连接到远程桌面
	var cmd *exec.Cmd

	if config.Port != 3389 {
		// 如果端口不是默认的3389端口，使用指定端口
		cmd = exec.Command("cmd", "/C", "mstsc", fmt.Sprintf("/v:%s:%d", config.IPAddress, config.Port))
	} else {
		// 使用默认端口
		cmd = exec.Command("cmd", "/C", "mstsc", fmt.Sprintf("/v:%s", config.IPAddress))
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("无法启动RDP连接: %v", err)
	}

	return nil
}

// 获取用户输入的简单函数
func getUserInput(prompt string) string {
	showMessageBox("输入", prompt)
	
	// 为了简化，我们使用命令行参数或预设值
	// 在实际应用中，可以使用更复杂的GUI输入框
	return ""
}

func main() {
	// 检查是否在Windows上运行
	if runtime.GOOS != "windows" {
		showMessageBox("错误", "此程序仅支持Windows系统")
		return
	}
	
	// 隐藏控制台窗口
	hwnd, _, _ := GetConsoleWindow.Call()
	if hwnd != 0 {
		ShowWindow.Call(hwnd, uintptr(SW_HIDE))
	}
	
	// 简单的交互流程
	ip := "127.0.0.1" // 在实际应用中，这里应该是用户输入
	port := 3389      // 在实际应用中，这里应该是用户输入
	username := ""    // 在实际应用中，这里应该是用户输入
	
	// 显示连接对话框
	showMessageBox("远程桌面连接", "请输入IP地址和端口进行连接")
	
	// 创建配置
	config := RDPConfig{
		ID:        fmt.Sprintf("rdp-%d", time.Now().Unix()),
		Name:      fmt.Sprintf("连接到 %s:%d", ip, port),
		IPAddress: ip,
		Port:      port,
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 尝试连接
	showMessageBox("连接中", fmt.Sprintf("正在连接到 %s (%s:%d)...", config.Name, config.IPAddress, config.Port))
	
	err := ConnectToRDP(config)
	if err != nil {
		showMessageBox("连接失败", err.Error())
		return
	}
	
	showMessageBox("连接成功", fmt.Sprintf("已成功连接到 %s:%d", config.IPAddress, config.Port))
}