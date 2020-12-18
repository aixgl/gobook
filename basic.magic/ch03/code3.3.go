/** 注释
 * 字符串与字节.    
 * module:github.com/aixgl/ch03
 * filename code3.3.go
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch03/code3.3
 */
package main

import (
    "fmt"
)

// 字节相当于int8的整数
func byteCal() {
    fmt.Println("=======byteCal========")
    b1 := 'H'
    b2 := b1 + 32
    // output: byte 运算 b1[H] b2[h]
    fmt.Printf("byte 运算 b1[%v] b2[%v]", string(b1), string(b2))
}

// 字符串与字节数组
func stringAndBytes() {
    fmt.Println("=======stringAndBytes========")
    s := "hello"
    bs := []byte(s)
    fmt.Printf("s => bs s[%v] bs[%v]\n", s, bs)
    
    bs[0] -= 32
    // bs[0]是切片的第一个byte因此可以计算
    fmt.Printf("string(bs)[%v]\n", string(bs))    
}

// 使用字符串为切片循环
func stringWithFor() {
    fmt.Println("=======stringWithFor========")
    s := "Hello,世界"
    
    // 按[]byte解析
    for i := 0; i < len(s); i++ { // byte
        fmt.Printf("%c,", s[i])
    }
    fmt.Println("------")
    
    // 按[]rune解析 这个类型在go中是比较特殊的存在。
    rs :=[]rune{}
    for _, r := range s { // rune
        rs = append(rs, r)
        fmt.Printf("%c,", r)
    }
	fmt.Println("\nstring([]rune)------", string(rs))
}