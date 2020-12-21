/** 注释
 * 复合类型.	
 * module:github.com/aixgl/ch05/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch05/code_main
 */
package main

import "fmt"


func arrayExample() {
    fmt.Println("=======arrayExample======")
    // 初始化 和 赋值
    //var a1 [3]int
    var a2 [3]int = [3]int{1, 2, 3}
    
    // 长度是根据初始化的值来确定的
    a3 := [...]int {1, 2, 4}

    for i, v := range a2 {
        a2[i] = v * 2
        fmt.Printf("a2 k:=%d, v:=%d, a2[i]:=%v\n", i, v, a2[i])
    }

    for i, n := 0, len(a3); i < n; i++ {
        fmt.Printf("a3[%d]:=%d \n", i, a3[i])
    }
}

func reverse(s1 []string) {
    for i, j:=0, len(s1)-1; i<j; i, j = i+1, j-1 {
        s1[i], s1[j] = s1[j], s1[i]
    }
    return s1
}
func sliceFirstBlood() {
    fmt.Println("=======arrayExample======")
    s1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}  
    //s2 := s1[3:8]
    s3 := s1[5:7]
    //s4 := s1[4:10]
    
    // 用range访问slice
    for i, v := range s3 {
        fmt.Printf("slice range k:%d, v:%s;\n", i, v)
    }
    reverse(s1)
    fmt.Println(s1)
}

func main() {
    arrayExample()
    sliceFirstBlood()
}