package main

import (
    "fmt"
    "net"
)

var nick string = ""

func main() {
    // 登录
    conn, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("connect server fail [127.0.0.1:8000]")
    }
    defer conn.Close()
    fmt.Println("connect server successed \n")

    // 取名
    fmt.Printf("Make a nickname:")
    fmt.Scanf("%s", &nick)
    fmt.Println("Hello : ", nick)
    // 将昵称同步到服务端 等同登录
    conn.Write([]byte("nick|" + nick))

    go handleClient(conn)

    // 通信
    var msg string
    for {
        msg = ""
        fmt.Scan(&msg)
        conn.Write([]byte("chat|" + nick + "|" + msg))
        // 登出
        if msg == "quit" || msg == "q" {
            conn.Write([]byte("quit|" + nick))
            break
        }
    }

}

// 实时读取服务端response
func handleClient(conn net.Conn) {
    for {
        data := make([]byte, 255)
        msg_read, err := conn.Read(data)
        if msg_read == 0 || err != nil {
            break
        }

        fmt.Println("@", string(data[0:msg_read]))
    }
}
