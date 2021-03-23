package main

import (
    "fmt"
    "time"
)

func printer(num int) {
    fmt.Println("I am printer", num)
}

func forRountine() {
    fmt.Println("--------forRountine--------")
    for i := 0; i < 10; i++ {
        go printer(i)
    }
    time.Sleep(5 * time.Second)
}

func channelWaitGoroutine(){
    fmt.Println("--------channelWaitGoroutine--------")
    done := make(chan struct{})
    go func() {
        sum := 0
        for i:=0; i < 100; i++ {
            sum += i
        }
        log.Println("done", sum)
        done <- struct{}{} // signal the main goroutine
    }()
    // 阻塞：wait for background goroutine to finish    
    <-done 
}