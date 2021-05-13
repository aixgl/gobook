package main

import(
    "net"
    "fmt"
    "bufio"
    "strings"
    "os"
)

func main() {
    TCPClient()
}

func TCPClient() {
    conn, err := net.Dial("tcp", "127.0.0.1:8088")
    if err != nil {
        fmt.Println("err :", err)
        return
    }
    defer conn.Close() // 关闭连接
    inputReader := bufio.NewReader(os.Stdin)
    for {
        // 标准输入读取
        input, _ := inputReader.ReadString('\n')
        inputInfo := strings.Trim(input, "\r\n")
        // 如果输入q/Q则退出
        if strings.ToUpper(inputInfo) == "Q" {
            return
        }
        // send 发送
        _, err = conn.Write([]byte(inputInfo))
        if err != nil {
            return
        }
        // recv from client 从服务端接收数据
        buf := [512]byte{}
        n, err := conn.Read(buf[:])
        if err != nil {
            fmt.Println("recv failed, err:", err)
            return
        }
        fmt.Println("recv:", string(buf[:n]))
    }
}