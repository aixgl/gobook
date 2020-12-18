/** 注释
 * 数值转换.    
 * module:github.com/aixgl/ch03
 * filename code3.2.go
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch03/code3.2
 */
package main

import (
    "fmt"
    "strcov"
)
// 简单的获取类型方法
func getType1(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 字符串转数字和bool
func stringToT() {
    fmt.Println("=======stringtoT========")
    //
    b, err := strconv.ParseBool("true")
    // 转成float64， 
    f, err := strconv.ParseFloat("3.1415", 64)
    // 转成int64，将"50"按16进制转成10进制值
    i, err := strconv.ParseInt("50", 16, 64)
    // 转成uint64，将"1024"按10进制转成10进制值。
    u, err := strconv.ParseUint("1024", 10, 64)
    
    // output: ParseBool 'true' type[bool] value[true] err[<nil>]
    fmt.Printf("ParseBool 'true' type[%v] value[%v] err[%v]\n", getType1(b), b, err)
    // output: ParseFloat  type[float64] value[3.1415] err[<nil>]
    fmt.Printf("ParseFloat  type[%v] value[%v] err[%v]\n", getType1(f), f, err)
    // output: ParseInt  type[int64] value[80] err[<nil>]
    fmt.Printf("ParseInt  type[%v] value[%v] err[%v]\n", getType1(i), i, err)
    // output: ParseUint  type[uint64] value[1024] err[<nil>]
    fmt.Printf("ParseUint  type[%v] value[%v] err[%v]\n", getType1(u), u, err)
}
// 其它基础类型转string
func TTostring() {
    fmt.Println("=======Ttostring========")
    
	s := strconv.FormatBool(true)
    fmt.Printf("ParseBool value[%v]\n", s )
    
    // float转字符串 //output:FormatFloat value[3.14159E+00]
    s = strconv.FormatFloat(3.14159, 'E', -1, 64)
    fmt.Printf("FormatFloat value[%v]\n", s )
    
    // int转字符串  //output:FormatInt value[-79d]
    s = strconv.FormatInt(-1949, 16)
    fmt.Printf("FormatInt value[%v]\n", s )
    
	// uint转字符串 //output: FormatUint value[152]
    s = strconv.FormatUint(152, 10)
    fmt.Printf("FormatUint value[%v]\n", s )
}