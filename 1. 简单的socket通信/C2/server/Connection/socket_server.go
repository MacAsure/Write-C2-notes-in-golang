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
		if certified(conn) {
			go process(conn)
		} else {
			log.Println("[-] 认证失败!")
		}
		// conn.Write([]byte("hostname"))

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

// 认证流程
func certified(conn net.Conn) bool {
	defer conn.Close()
	buf := make([]byte, 32)
	n, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Println(err)
	}
	if string(buf[:n]) == "877869CBFED11FC453C218174121CC7C" {
		return true
	} else {
		return false
	}

}

// func read_server(conn net.Conn) {

// 	// 设置一次接收多少数据
// 	buf := make([]byte, 2)
// 	reader := bufio.NewReader(conn)

// 	// 设置期望接收长度
// 	// expectlength := 1024
// 	n, err := io.ReadFull(reader, buf)
// 	fmt.Println("我在这1")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("客户端执行结果:")
// 	fmt.Println(string(buf[:n]))

// }
