package main

import (
    "fmt"
    "time"
    "context"
)

// -------------------------------------------
// select 选择case 会比较随机
func randSelectDemo() { 
    fmt.Println("----------randSelectDemo-----------")
    c1 := make(chan int)
    c2 := make(chan int)
    go pushRandtSelect(c1, c2)
    go randSelect(c1, c2)
    
    time.Sleep(2 * time.Second)
    //select {}
}
// 3次监听select 来查看运行情况
func randSelect(c1, c2 chan int) {
    
    time.Sleep(1 * time.Second)
    for i :=0; i< 3 ; i++ {
        select {
        case out := <-c1 :
            fmt.Println("select rec c1:=", out)
        case c2 <- 22 :
            fmt.Println("select send c2")
        default :
            fmt.Println("no select!")
        }
    }
}
// 第二个goroutine 运行发送和接收
func pushRandtSelect(c1, c2 chan int) {
    c1 <- 3
    fmt.Println("select rec", <-c2)
}

// -------------------------------------------
// 使用 select 实现 timeout 机制
func timeoutSelect() {
    fmt.Println("----------timeoutSelect-----------")
    timeout := make (chan bool, 1)
    go func() {
        time.Sleep(1 * time.Second) // sleep one second
        timeout <- true
    }()
    select {
    case <- timeout:
        // 延时执行逻辑
        fmt.Println("timeout!")
    }
}

// -------------------------------------------
// close chan 通知退出goroutine
func exitCloseChan() {
    fmt.Println("----------exitCloseChan-----------")
    done :=make(chan bool)
    go func() {// goroutine 1
        for{
            select {
            case <-done:
                fmt.Println("exit goroutine 1")
                return
            default:
                fmt.Println("task goroutine 1")
                time.Sleep(1*time.Second)
            }
        }
    }()
 
    go func() {// goroutine 2
        // 没有消息阻塞状态，chan关闭 for 循环结束
        for msg :=range done{ 
            fmt.Println("goroutine 2 do somethings.", msg)
        }
        fmt.Println("exit goroutine 2")
    }()
 
    go func() {// goroutine 3
        for{
            select {
            case <-done:
                fmt.Println("exit goroutine 3")
                return
            default:
                fmt.Println("task goroutine 3")
                time.Sleep(1*time.Second)
            }
        }
    }()
    time.Sleep(2*time.Second)
    close(done)
    time.Sleep(3*time.Second)
    fmt.Println("退出程序")
}

// 退出协程
func exitContext() {
    fmt.Println("----------exitContext-----------")
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    go func() {
        for{
            select {
            case <-ctx.Done():
                fmt.Println("exit goroutine 1")
                return
            default:
                fmt.Println("running...")
                time.Sleep(1*time.Second)
            }
        }
    }()
    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("exit goroutine 2")
                return
            }
        }
    }()
    time.Sleep(6 * time.Second)
    cancel() //也可以手动调用 cancel()方法退出
    time.Sleep(1*time.Second)
    fmt.Println("程序退出")
}

// 使用WithCancel 退出协程
func exitContextCancel() {
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        for{
            select {
            case <-ctx.Done():
                fmt.Println("exit goroutine")
                return
            default:
                fmt.Println("running...")
                time.Sleep(1*time.Second)
            }
        }
    }()
 
    time.Sleep(3*time.Second)
    cancel()
    time.Sleep(1*time.Second)
    fmt.Println("退出程序")
}  

// WithDeadline 退出协程
func exitDeadLine() {
    taskTime := "2021-03-25 16:23:59"
    loc, _ := time.LoadLocation("Local")
    c_time, _ := time.ParseInLocation("2006-01-02 15:04:05", taskTime, loc)
    ctx, _ := context.WithDeadline(context.Background(), c_time)
    go func() {
        for{
            select {
            case <-ctx.Done():
                fmt.Println("exit goroutine")
                return
            default:
                fmt.Println("event ...")
                time.Sleep(1*time.Second)
            }
        }
    }()
 
    time.Sleep(60*time.Second)
    fmt.Println("程序退出")
}
