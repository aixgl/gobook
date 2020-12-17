/** 注释
 * 数值转换.	
 * module:github.com/aixgl/ch03
 * filename code3.1.go
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch03/code3.1
 */
package main

import (
    "fmt"
    "math"
)

// 长度大转长度小
func int64toint32() {
    fmt.Println("=======int64toint32========")
	//转换 长度从大到小
    // 超过32位，转换后是0值。
	num64 := int64(math.Pow(2, 38))
	num32 := int32(num64)
    // output: num64:=274877906944, num32:=0
	fmt.Printf("num64:=%d, num32:=%d\n", num64, num32)   
    
    // 少于32位
    num64 = int64(1024)
	num32 = int32(num64)
    // output: num64:=1024, num32:=1024
	fmt.Printf("num64:=%d, num32:=%d\n", num64, num32)
}

// 同长度不同类型名 也要转成相同的
func int32toint() {
    fmt.Println("=======int32toint========")
    num := 32
	num2 := int32(2)
    
    /* 编译报错：invalid operation: num * num2 (mismatched types int and int32)*/
    // fmt.Println(num * num2) // int 和 int32 的类型不匹配

    // output:64 //可以正确执行
	fmt.Println(num * int(num2))  
}

