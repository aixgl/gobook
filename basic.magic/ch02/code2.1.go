package main                                                                                                        
/** 
 * 命名.	
 * module:github.com/aixgl/ch02/code2.1
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch02/code2.1
 */
import "fmt"

type myInt int32

const PI = 3.14

func main() {
    var r = 5.0 
    var cArea = PI * r * r 
    fmt.Println("circle R:=%f, C:=%f", r, cArea)
}
