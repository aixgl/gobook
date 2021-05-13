package main

import (
    "fmt"
    "sync"
)


// 利用通道进行通信处理并发
func concurrencyChannel() {
    fmt.Println("----------concurrencyChannel-----------")
    // 开始逻辑
    begin()
    // 中间并发处理逻辑
    c := make(chan bool, 5)
    for i := 0; i < 5; i++ {
        go func(i int) {
            concurrencyCommon(i)
            c <- true
        }(i)
    }
    // 等待并发
    for i := 0; i < 5; i++ {
        <-c
    }
    // 结尾逻辑
    end()
}
// 用waitGroup 控制并发
func concurrencyWaitGroup() {
    fmt.Println("----------concurrencyWaitGroup-----------")
    // 开始逻辑
    begin()
    // 中间并发处理逻辑 
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            concurrencyCommon(i)
        }(i)
    }
    // 等待并发
    wg.Wait()   
    // 结尾逻辑
    end()    
}

func concurrencyCommon(work int) {
    fmt.Println("concurrency common work", work)
}

func begin() {
    // todo somethings
    fmt.Println("start somethings")
}

func end() {
    // todo somethings at the end
    fmt.Println("END")
}