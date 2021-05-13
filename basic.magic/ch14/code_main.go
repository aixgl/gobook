/** 注释
 * 接口.
 * module:github.com/aixgl/ch14/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch14/code_main
 */
package main

import (
    "reflect"
    "fmt"
)

// --------------------------------
// * kind 
// --------------------------------
type MyInt int

var a int = 1
var b MyInt = 2
// a 和 b有相同的Kind “int”
func typeAndKind() {
    at := reflect.TypeOf(a)
    bt := reflect.TypeOf(b)
    fmt.Printf("int a type[%s] kind[%s]\n", at.Name(), at.Kind())
    fmt.Printf("MyInt b type[%s] kind[%s]\n", bt.Name(), bt.Kind())
}

func valueSample() {
    var a int = 100
    // 获取变量a的反射值对象
    valueOfA := reflect.ValueOf(a)
    // 获取interface{}类型的值, 通过类型断言转换
    var valueA int = valueOfA.Interface().(int)
    // 获取64位的值, 强制类型转换为int类型
    var valueB int = int(valueOfA.Int())
    fmt.Println("Value 断言:",valueA, "Value 强制转换:", valueB)
}
func main() {
    // http pprof
    // httpPProfServer()
}
