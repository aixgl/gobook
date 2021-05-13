package main

import (
    //"fmt"
    "github.com/slclub/link"
    "net"
    "strings"
)

var ConnMgr map[string]net.Conn = make(map[string]net.Conn)

func main() {
    service()
}

func service() {
    service_socket, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        link.ERROR("server listen start error!")
        return
    }

    defer service_socket.Close()
    link.INFO("server is running...")

    for {
        conn, err := service_socket.Accept()
        if err != nil {
            link.ERROR("server accepte conn error!")
        } else {
            link.INFO("server accepte conn success")
        }

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    for {
        data := make([]byte, 255)
        read_data, err := conn.Read(data)
        if read_data == 0 || err != nil {
            continue
        }

        //解析协议
        //第一个元素为处理流程选择条件
        msg_data := strings.Split(string(data[0:read_data]), "|")

        switch msg_data[0] {
        case "nick":
            link.INFO(conn.RemoteAddr(), "-->", msg_data[1])
            for k, v := range ConnMgr {
                if k != msg_data[1] {
                    v.Write([]byte("[" + msg_data[1] + "]: join..."))
                }
            }
            // 绑定客户端连接到昵称
            ConnMgr[msg_data[1]] = conn
        case "chat":
            for k, v := range ConnMgr {
                if k != msg_data[1] {
                    link.INFO("chat "+msg_data[2]+" to ", k)
                    v.Write([]byte("[" + msg_data[1] + "]: " + msg_data[2]))
                }
            }
        case "quit":
            for k, v := range ConnMgr {
                if k != msg_data[1] {
                    v.Write([]byte("[" + msg_data[1] + "]: quit"))
                } else {
                    link.INFO("QUIT", msg_data[1], "was left!")
                }
            }
            delete(ConnMgr, msg_data[1])
        }

    }
}
