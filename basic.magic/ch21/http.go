package main

import(
    "fmt"
    "net"
    "net/http"
    "io/ioutil"
    "io"
)
func main() {
    TCP_HTTPClient()
    HTTPClient()
}
func TCP_HTTPClient() {
    conn, err := net.Dial("tcp", "www.baidu.com:80")
    if err != nil {
        fmt.Println("tcp dial failed, err:", err)
        return
    }
    defer conn.Close()
    // 发送请求 http1.0协议
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") 

    var buf [8192]byte
    // 接收响应
    for {
        n, err := conn.Read(buf[:])
        // 相应后退出
        if err == io.EOF {
            return
        }
        if err != nil {
            fmt.Println("tcp-http response failed, err:", err)
            break
        }
        fmt.Println("tcp-http response:",string(buf[:n]))
    }
}

func HTTPClient() {
    resp, err := http.Get("https://www.baidu.com/")
    if err != nil {
        fmt.Println("get failed, err:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body) )
}

// 使用tcp模拟http
func HTTPServer() {
    li, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Fatalln(err.Error())
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            fmt.Fatalln(err.Error())
            continue
        }
        // 函数前添加 go 关键字，就能使其拥有 Go 语言的并发功能
        // 这样我们可以同时处理来自不同客户端的请求
        go tcpHandle(conn)
    }
}

func tcpHandle(conn net.Conn) {
    // 短连接
    defer conn.Close()
    // 回应客户端的请求
    respond(conn)
}

func respond(conn net.Conn) {
    // 理论拆分成ServeHTTP(w http.ResponseWriter, r *http.Request)
    // web内容
    body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>Go example</title></head><body><strong>Hello World</strong></body></html>`
    
    // 组织http.ResponseWriter.Header
    // HTTP 协议及请求码
    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    // 内容长度
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) 
    // MIME
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    
    // http.ResponseWriter.Write([]byte)
    fmt.Fprint(conn, body)
}