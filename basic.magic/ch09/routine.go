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

func channelWaitGoroutine() {
    fmt.Println("--------channelWaitGoroutine--------")
    done := make(chan struct{})
    go func() {
        sum := 0
        for i := 0; i < 100; i++ {
            sum += i
        }
        fmt.Println("done", sum)
        done <- struct{}{} // signal the main goroutine
    }()
    // 阻塞：wait for background goroutine to finish
    <-done
}

func channelSyncGoroutines() {
    fmt.Println("--------channelSyncGoroutines--------")
    counter := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 99; x >= 0; x-- {
            counter <- x
        }
    }()

    // Squarer
    go func() {
        for {
            x, ok := <-counter
            squares <- x * x
            if !ok {
                break
            }
        }
    }()
    sum := 0
    for x := range squares {
        sum += x
        // 防止 all goroutines are asleep - deadlock
        if x == 0 {
            break
        }
    }
    fmt.Println("done", sum)
}

func channelSyncGoroutines2() {
    fmt.Println("--------channelSyncGoroutines2--------")
    counter := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 99; x >= 0; x-- {
            counter <- x
        }
        close(counter)
    }()

    // Squarer
    go func() {
        for x := range counter {
            squares <- x * x
        }
        close(squares)
    }()
    sum := 0
    for x := range squares {
        sum += x
    }
    fmt.Println("done", sum)
}

func counter(input chan <- int) {
    for x := 99; x >= 0; x-- {
        input <- x
    }
    close(input)
}

func square(input chan <- int, output <- chan int) {
    for x := range output {
        input <- x * x
    }
    close(input)
}

/**
    go counter(naturals)
    go squarer(squares, naturals
 */