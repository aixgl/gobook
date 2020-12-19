/** 
 * Command Args.	
 * module:github.com/aixgl/ch01/code1.2    
 * https://github.com/aixgl/gobook/tree/master/basic.magic/ch01/code1.2
 */
package main                                                                                                        

import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args {
        fmt.Println("arg", i, "=", arg)
    }   
}
