package main

import (
    "flag"
    "fmt"
)

func init() {
    
    initParseFlagVar()
    
    flag.Parse()
    
    parseFlag()
    parseFlagVar()
}
/*--------------------------------------------------*
 * 直接使用 flag.T(命令行变量名, 默认值, 帮助说明文案)
 * T为基础类型首字母大写 Bool, Int, String
 * API: flag.Bool, flag.Int, flag.String
 * flag.T()的返回值是指针变量
 *--------------------------------------------------*/
 // 定义命令行指针变量和默认值
var port = flag.Int("port", 808, "help message for port")
var version = flag.Bool("version", false, "Check version of the soft!")
var name = flag.String("name", "", "Name of project")
// 解析
func parseFlag() {
    fmt.Println("====== parseFlag ======")
    // 解析命令行 complie
    // flag.Parse()
    // 命令行指针port 端口非默认值打印出端口
    if *port != 808 {
        fmt.Println("The progject port is", *port)
    }
    
    // 命令行指针version 查看软件版本号
    if *version {
        fmt.Println("Project version is 8.0.1")
    }
    
    // 命令行指针name 打印软件名
    if *name != "" {
        fmt.Println("Project name is", *name)
    }
}

/*--------------------------------------------------*
 * 直接使用 flag.TVar(存储指针变量, 命令行变量名, 默认值, 帮助说明文案)
 * T为基础类型首字母大写 Bool, Int, String
 * API: flag.BoolVar, flag.IntVar, flag.StringVar
 * flag.TVar将命令的值赋值给存储指针变量
 *--------------------------------------------------*/
// 与parseFlag完成同样的工作，仅仅在变量名后面加入_var加以区分，flag包

// 定义命令行指针变量和默认值
var port_var int 
var version_var bool 
var name_var string

func initParseFlagVar() {
    flag.IntVar(&port_var, "p", 808, "help message for port")
    flag.BoolVar(&version_var, "v", false, "Check version of the soft!")
    flag.StringVar(&name_var, "nm", "", "Name of project")
}

func parseFlagVar() {
    fmt.Println("====== parseFlagVar ======")
    // flag.Parse 只会运行一次
    // flag.Parse()
    // 命令行指针port 端口非默认值打印出端口
    if port_var != 808 {
        fmt.Println("The progject port is", port_var)
    }
    
    // 命令行指针version 查看软件版本号
    if version_var {
        fmt.Println("Project version is 8.0.1")
    }
    
    // 命令行指针name 打印软件名
    if name_var != "" {
        fmt.Println("Project name is", name_var)
    }
}

/*--------------------------------------------------*
 * 直接使用 flag.Var(T, 命令行变量名, 默认值, 帮助说明文案)
 * T为自定义类型
 * flag.TVar将命令的值赋值给存储指针变量
 *--------------------------------------------------*/