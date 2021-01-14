package main

import (
    "flag"
    "fmt"
    "strconv"
)
// flag 解析仅有一次 因此定义接收命令行的变量可以定义在包级别
var port *int
var version *bool
var name *string

func init() {
    // 绑定命令行参数到变量
    initParseFlag()
    initParseFlagVar()
    initParseFlagDefine()
    
    // 解析命令行
    flag.Parse()
    
    // 处理命令行参数
    parseFlag()
    parseFlagVar()
    parseFlagDefine()
}
/*-----------------------------------------------------*
 * 直接使用 flag.T(命令行变量名, 默认值, 帮助说明文案) *
 * T为基础类型首字母大写 Bool, Int, String             *
 * API: flag.Bool, flag.Int, flag.String               *
 * flag.T()的返回值是指针变量                          *
 *-----------------------------------------------------*/
 // 定义命令行指针变量和默认值
func initParseFlag() {
    port = flag.Int("port", 808, "help message for port")
    version = flag.Bool("version", false, "Check version of the soft!")
    name = flag.String("name", "", "Name of project")
}
// 处理命令行参数
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
func initParseFlagVar() {
    flag.IntVar(port, "p", 808, "help message for port")
    flag.BoolVar(version, "v", false, "Check version of the soft!")
    flag.StringVar(name, "nm", "", "Name of project")
}

func parseFlagVar() {
    fmt.Println("====== parseFlagVar ======")
    // flag.Parse 只会运行一次
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
 * 直接使用 flag.Var(T, 命令行变量名, 默认值, 帮助说明文案)
 * T为自定义类型
 * T 需要实现flag包的
        // Value 接口
        type Value interface {
            String() string
            Set(string) error
        }
 * flag.TVar将命令的值赋值给存储指针变量
 *--------------------------------------------------*/
type fint int
// 实现flag.Value.Set
func (i *fint) Set(s string) error {
    v, err := strconv.ParseInt(s, 0, strconv.IntSize)
    *i = fint(v)
    return err
}
// 实现flag.Value.String
func (i *fint) String() string { 
    return strconv.Itoa(int(*i)) 
}
// ---------------------------------------------------
var vs  = fint(0)
func initParseFlagDefine() {
    flag.Var(&vs, "vs", "")
}

func parseFlagDefine() {
    if vs != 0 {
        fmt.Println("parse flag vs :", vs)
    }
}