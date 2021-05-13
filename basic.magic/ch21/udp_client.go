package main

import(
    "net"
    "fmt"
)
func main() {
    UDPClient()
}

func UDPClient() {
    udpClient, err := net.DialUDP("udp", nil, &net.UDPAddr{
        IP:   net.IPv4(192, 168, 3, 252),
        Port: 8088,
    })
    if err != nil {
        fmt.Println("connect UDP server failed，err:", err)
        return
    }
    defer udpClient.Close()
    // 发送数据
    _, err = udpClient.Write([]byte("Hello udp"))
    if err != nil {
        fmt.Println("UDP send failed，err:", err)
        return
    }
    data := make([]byte, 4096)
    // 接收数据
    n, remoteAddr, err := udpClient.ReadFromUDP(data)
    if err != nil {
        fmt.Println("UDP recv failed，err:", err)
        return
    }
    fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}