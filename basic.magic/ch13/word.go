package main

import (
    "unicode"
)

//s1 := "我是Gopher"
func hasChinese(s1 string) bool {
    for _, v := range s1 {
        if unicode.Is(unicode.Han, v) {
            return true
        }
    }
    return false
}

func abs(num int) int {
    // 故意写错
    return num * -1
    /* 正确求绝对值
    if num < 0 {
        return num * -1
    }
    return num
    */
}


func divide(a, b int) int {
    // 需要check b == 0
    return a/b
}