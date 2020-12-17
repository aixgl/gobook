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

func int64toint32() {
	//转换 长度从大到小
    // 超过32位
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

