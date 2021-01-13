/** 注释
 * 接口.
 * module:github.com/aixgl/ch07/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch07/code_main
 */
package main

import (
    "fmt"
    "io"
    "bufio"
    "strings"
)

func InterfaceCom() {
    fmt.Println("======InterfaceCom======")
    sp := &space {
        str : "hello, pite\n hi sam",
    }
    
    // reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    // reader := bufio.NewReader(fb) // 从文件中读取
    reader := bufio.NewReader(sp)
    text, _ := reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
    fmt.Printf("bufio.NewReader %#v\n", text)
}

// 测试空接口
func EmptyInterface() {
    fmt.Println("======EmptyInterface======")
    var r io.Reader
    fmt.Printf("io.Reader type[%v] value[%T]\n", r, r)
}

func main () {
    EmptyInterface()

}