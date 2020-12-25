/** 注释
 * 复合类型.
 * module:github.com/aixgl/ch06/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch06/code_main
 */
package main

import (
    "fmt"
    "errors"
    "os"
    "bufio"
    "io"
)

// 
func add (x int, y int) int{
    return x+y
}

// 
func add2 (x int, y int) (c int) {
    c = x + y
    return
}

// 实现一个求幂值的函数
func powMath (x, y int64) int64 {
   if y == 0 {
       return 1
   }
    r := x
   for {
       y--
       if y == 0 {
           break
       }
       r *= x
   }
   return r
}

// 递归函数
func fib(n int) int {
    if n ==0 || n == 1 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func forFib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}

func stepMul(n int, step int) int {
    if n <= 0 {
        return 1
    }
    if step == 0 {
        step = 1
    }
    return n * stepMul(n - step, step)
}

func gotoMul(n int) int {
    r := n
WALK :
    n--
    if n <= 0 {
        return r
    }
    r *= n
    goto WALK
}

// 错误处理
func errFunc () error{
    fmt.Println("=======errFunc======")
    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err == io.EOF {
            break // finished reading
        }
        if err != nil {
            return errors.New("字符不是utf格式")
        }
        // ...use r…
        fmt.Printf("%v", r)
    }
    return nil
}

// 错误处理
type ERR_CODE int
func errCodeFunc () ERR_CODE {
    fmt.Println("=======errFunc======")
    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err == io.EOF {
            break // finished reading
        }
        if err != nil {
            return -10020
        }
        // ...use r…
        fmt.Printf("%v", r)
    }
    return 0
}


func main() {
    fmt.Println("=======Chapter 6======")
}
