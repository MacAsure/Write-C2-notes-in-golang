package Connection

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Listen_server(address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("[!] 服务端监听失败", err)
		return
	}
	fmt.Println("[+] 服务端监听成功!")
	for {
		// 建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败:", err)
			return
		}
		conn.Write([]byte("hostname"))
		go process(conn)

	}
}

// 处理接收发送数据
func process(conn net.Conn) {
	// 退出时关闭连接
	defer conn.Close()

	for {

		read_server(conn)
		reader_server := bufio.NewReader(os.Stdin)
		fmt.Println("请输入指令:")
		input, _ := reader_server.ReadString('\n')
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("数据发送失败:", err)
		}
	}
}

// func Trim(input string) string {
// 	inputInfo := strings.Trim(input, "\r\n")
// 	inputInfo = strings.Trim(inputInfo, " ")
// 	return inputInfo

// }

// 接收客户端数据
func read_server(conn net.Conn) {
	defer conn.Close()
	// 设置一次接收多少数据
	buf := make([]byte, 4096)
	reader := bufio.NewReader(conn)
	for {
		n, err := reader.Read(buf)

		if err != nil {
			fmt.Println("接收数据结束:", err)
		}
		result := string(buf[:n])
		fmt.Println(result)
		if n < len(buf) {
			break
		}

	}
}

