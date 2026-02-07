package main

import (
	"fmt"
	"net"
	"time"
)

// TestPortConnectivity 测试端口连通性
func TestPortConnectivity(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	
	fmt.Printf("正在测试 %s 的连通性...\n", address)
	
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf("端口测试失败: 无法连接到 %s - %v\n", address, err)
		return false
	}
	
	defer conn.Close()
	fmt.Printf("端口测试成功: %s 可达\n", address)
	return true
}

func main() {
	fmt.Println("远程桌面连接客户端 - 端口测试工具")
	fmt.Println("================================")
	
	// 示例测试
	TestPortConnectivity("127.0.0.1", 3389)  // 本地RDP端口
	TestPortConnectivity("127.0.0.1", 80)    // 本地HTTP端口
	TestPortConnectivity("8.8.8.8", 53)      // Google DNS
	
	fmt.Println("\n测试完成。")
	fmt.Println("按任意键退出...")
	fmt.Scanln()
}