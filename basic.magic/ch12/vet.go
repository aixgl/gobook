package main

import (
    "fmt"
    "net/http"
)

func vetPrint() {
    s := "vet printf error format"
    fmt.Printf("%d\n", s)
}

// bool 问题
func vetBool() {
    i:= 10
    fmt.Println(i != 0 || i != 1)    // always false
    fmt.Println(i == 0 && i == 1)    // redundant check
    fmt.Println(i == 0 && i == 0)
}

// goroutine在主线程退出后还未执行
func vetGoroutineNotWork(){
    words := []string{"foo", "bar", "baz"}    
    for _, word := range words {
        go func() {
            fmt.Println(word)
        }()
    }
}

//unreachable code
func vetFibonnaqi(n int) int {   
    switch n {    
        case 0:        return 1
        case 1:        return 1
        default:       return vetFibonnaqi(n-1) + vetFibonnaqi(n-2)
    }
    fmt.Println("unreachable code")    
    return 0
}

// 小整数向右位移过多
func vetIntMove(){
    i := 100
    fmt.Println(i >> 32)
}

// 使用前未检测err
func vetErrorCheck(){
    res, err := http.Get("https://github.com/")
    defer res.Body.Close()
    // 应放在defer语句前才对
    if err != nil {
        fmt.Println("http get error!")
    }
}