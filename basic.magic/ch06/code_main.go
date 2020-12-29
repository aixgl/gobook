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

// 匿名函数
func anonymousExs() {
    sq := func (x int) int {
        return x * x
    }
    fmt.Println("anonymous  sq1:=", sq(2))  // sq1:=4
}
// 一个求和加法器
func adder() func(int) int {
    sum := 0// 试着修改sum的初始值为1 
    return func(n int) int {
        sum += n
        return sum
    }
}
func anonymousEnv()  {
    fmt.Println("=======anonymousEnv======")
    a := adder()
    fmt.Println(a(0), a(2), a(4))
}
// 信号量处理
func signalHandle() {
    for {
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2,syscall.SIGHUP)
        sig := <-ch
        link.DEBUG("Signal received: %v", sig)
        switch sig {
        
        case syscall.SIGHUP:
            link.DEBUG("get sighup\n")
        case syscall.SIGINT:
            os.Exit(1)
        case syscall.SIGUSR1:
            link.INFO("usr1\n")
        case syscall.SIGUSR2:
            link.INFO("usr2\n")
        default:
            panic("we don't have the process")
            link.INFO("get sig=%v\n",sig)
        }
    }
}

func deferStack() {
    fmt.Println("=======deferStack======")
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    defer fmt.Println("defer 3")
}

func deferReturn() int {
    i := 0
    addi := func() {
        i++
        fmt.Println("defer", i)
    }
    defer addi()
    defer addi()
    // 此处return i 是值拷贝操作 将i理解为return函数的参数
    // 等价于s:=i; return s;
    return i
}

func deferReturnRef()(i int) {
    //i := 0  // 与deferReturn的区别
    addi := func() {
        i++
        fmt.Println("defer", i)
    }
    defer addi()
    defer addi()
    // 此处return i 是值拷贝操作 将i理解为return函数的参数
    // 等价于s:=i; return s;
    return i
}

func recoverPanic() {
    defer func () {
        if err := recover(); err == nil {
            return
        }
        // do somethings
    }()
    n := 0
    n--
    if n == -1 {
        panic("n must bigger than -1")
    }
    // return // 无返回值定义，可将花括号的前一行有return
}

func main() {
    fmt.Println("=======Chapter 6======")
    deferStack()
    fmt.Println("return", deferReturn() )
    fmt.Println("return", deferReturnRef() )
}
