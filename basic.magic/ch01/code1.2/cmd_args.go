package main                                                                                                        
/** 
 * Command Args.	
 * module:github.com/aixgl/ch01/code1.2    
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch01/code1.2
 */
import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args {
        fmt.Println("arg", i, "=", arg)
    }   
}