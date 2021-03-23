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
    time.Sleep(2 * time.Second)
}