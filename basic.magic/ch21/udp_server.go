package main

import(
    "net"
    "fmt"
)
func main() {
    UDPServer()
}

func UDPServer() {
    listen, err := net.ListenUDP("udp", &net.UDPAddr{
        IP:   net.IPv4(0, 0, 0, 0),//可以允许所有IP访问
        Port: 8088,
    })
    if err != nil {
        fmt.Println("UDP listen failed, err:", err)
        return
    }
    defer listen.Close()
    for {
        var data [1024]byte
        // 接收数据，对比TCP只有 地址对象，没有连接对象
        n, addrClient, err := listen.ReadFromUDP(data[:])
        if err != nil {
            fmt.Println("UDP read failed, err:", err)
            continue
        }
        fmt.Printf("data:%v address:%v lenght:%v\n", string(data[:n]), addrClient, n)
        // 发送报文数据 依据地址
        _, err = listen.WriteToUDP(data[:n], addrClient)
        if err != nil {
            fmt.Println("UDP write failed, err:", err)
            continue
        }
    }
}