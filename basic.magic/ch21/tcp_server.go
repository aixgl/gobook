package main

import(
    "net"
    "fmt"
    "bufio"
)

func main() {
    TCPServer()
}

func TCPServer() {
    // 用于 network I/O 开发，包括TCP/IP, UDP 协议等
    listen, err := net.Listen("tcp", ":8088")
    if err != nil {
        fmt.Println("tcp server listen failed, err:", err)
        return
    }
    for {
        // 建立连接
        // 若无连接会阻塞
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept failed, err:", err)
            continue
        }
        // 
        go processPack(conn)
    }
}

func processPack(conn net.Conn) {   
    defer conn.Close() // 延迟关闭
    for {
        reader := bufio.NewReader(conn)
        var bufStream [256]byte
        // Read
        n, err := reader.Read(bufStream[:])
        if err != nil {
            fmt.Println("read from client failed, err:", err)
            break
        }
        recvStr := string(bufStream[:n])
        fmt.Println("tcp recv from client:", recvStr)
        // 通知客户端
        conn.Write([]byte(recvStr))
    }
}

