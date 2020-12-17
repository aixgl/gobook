var gla int

func f () {//作用域f函数内
    fx, fy = 1, 2
    gla = 3
    // 作用域if_1
    if fx == 1 {
        fx := 5                            // if_1的fx 与函数的fx局部变量是不同地址的变量，同名而已
        ffx := 6                           // 此变量只能在这个作用域下使用
        
        //fy是函数作用下的变量，可以正确使用，gla是全局变量也可以正确使用
        fmt.Println(fx, fy, gla)           // output: 5 2 3  
        fmt.Println(ffx)                   // output: 6
    }
    
    fmt.Println(ffx)                       // 编译不通过 未定义
    
    // 作用域 if_2
    if fy == 2 {
        ffx := 7                           // if_2下的ffx只能在这里使用
        fmt.Println(fx)                    // output: 1 //if_1的fx重定义并未影响fx
        fx := 4                            // if_2下的fx 与函数的fx是不同的变量，地址和值都不相同与 if_1下的fx也不同
        gla := 33                          // if_2下的gla与全局的gla是不同的变量，地址和值都不相同
        fmt.Println(fx, fy , gla)          // output: 4 2 33
        fmt.Println(ffx)                   // output: 7
    }
    fmt.Println(gla)                       // output:3
}

