/** 注释
 * 控制流.	
 * module:github.com/aixgl/ch04/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch04/code_main
 */
package main

import "fmt"

func forRangeSlice() {
    fmt.Println("==========forRangeSlice==========")
    for key, value := range []int{1, 2, 3, 4} {
        fmt.Printf("key:=%d value:=%d\n", key, value)
    }
}

func forRangeString() {
    fmt.Println("==========forRangeString==========")
    s := "hello世界"
    for key, value := range s {
        fmt.Printf("key:%d value:0%s\n", key, value)
    }
}

func forRangeGetPointer() {
    fmt.Println("==========forRangeGetPointer==========")
    arr := []int{1,2,3}
    for i,_ := range arr {
        fmt.Println("address:", &arr[i], "value:=", arr[i])
    }
}

func breakAndContinue() {
    fmt.Println("==========breakAndContinue==========")
    for i:=0; i < 100; i++ {
        if i > 88 {
            break       // 若是有定义标签 break 标签名
        }
        if i % 3 == 0 {
            continue
        }
    }
}

func breakAndLabel () {
    fmt.Println("==========breakAndLabel==========")
    FOR1:
    for i:=0; i < 3; i++ {
        fmt.Printf("FOR1 第%d次循环\n", i)
        // FOR2:
        for j:=0; j < 3; j++ {
            fmt.Printf("FOR2 第%d次循环\n", i)
            break FOR1
        }
    }
}

func gotoAndLabel() {
    fmt.Println("==========gotoAndLabel==========")
    for x := 0; x < 10; x++ {
        for y := 0; y < 10; y++ {
            if x == 6 {
                // 跳转到标签
                goto DONE_ONE
            }
        }
    }
    // do somethings
    // 手动返回, 避免执行进入标签
    return
    // 标签
DONE_ONE:
    fmt.Println("done one")
}

func main() {
    forRangeSlice()
    forRangeString()
    forRangeGetPointer()
    breakAndContinue()
    breakAndLabel ()
    gotoAndLabel() 
}