package main

import (
	"C2/server/Common"
	"C2/server/Connection"
	"bufio"
	"fmt"
	"os"
)

func main() {
	Common.Flag()
	// 配置监听
	reader_server := bufio.NewReader(os.Stdin)
	fmt.Println("输入监听ip及端口:")
	input, _, _ := reader_server.ReadLine()
	Connection.Listen_server(string(input))

}

// func read_server(conn net.Conn) {
// 	defer conn.Close() // 关闭连接
// 	reader := bufio.NewReader(conn)
// 	fmt.Println(reader)
// 	for {
// 		// 读取数据长度
// 		lenByte, err := reader.ReadByte()
// 		if err != nil {
// 			fmt.Println("接收数据结束:", err)
// 			break
// 		}
// 		len := int(lenByte) // 转换为整数
// 		// 读取数据内容
// 		data, err := reader.ReadBytes('\r') // 使用\0作为分隔符
// 		if err != nil {
// 			fmt.Println("接收数据结束:", err)
// 			break
// 		}
// 		res := string(data[:len]) // 截取有效长度
// 		fmt.Println("客户端执行结果:", res)
// 		break
// 	}

// }
