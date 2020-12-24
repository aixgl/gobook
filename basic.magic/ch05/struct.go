/**
 * struct
 */
package main

import "fmt"

// 定义一个结构体
type User struct {
    Uid     int
    Name    string
    Age     int
    Weight  int
    Height  int
}

// 为嵌入User做准备的结构体
type Employee struct {
    User
    ManagerID   int
    Work        int
}

// 一个点的坐标结构体
type Point struct {
    X   int
    Y   int
}

// 值类型初始化
func structInitValue() {
    fmt.Println("=======structInitValue======")
    // 1 实例化
    var u1 User     // 声明User的实例
    
    // 2 可以省略字段， 全部赋值值必须与结构体属性一一匹配
    u2 := User{1, "joy", 20, 110, 170}    
    // 3 键值形式，可以部分赋值。
    u3 := User{Name : "Game", Height: 169, Weight: 120}

    fmt.Println("值类型初始化：", "\n", u1, "\n", u2, "\n",u3)
}

// 初始化成指针
func structInitPointer() {
    fmt.Println("=======structInitPointer======")
    // 1 只声明值== nil
    var up1 *User      
    // 2 new(T)一个有效的结构体指针 *upNew的每项属性被初始化
    var upNew *User = new(User)
    
    // 结构体前面加上"&"地址符号等同new
    // 3 省略字段，全部赋值
    up3 := &User{1, "p3", 20, 110, 170}
    // 4 键值形式
    up4 := &User{Uid: 2, Name:"p4", Height: 180}
    fmt.Println(1, up1 == nil)  // output:1 true
    fmt.Println(2,up2, *up2)    // output:2 &{0  0 0 0} {0  0 0 0}
    fmt.Println(3,up3, *up3)    // output:3 &{1 p3 20 110 170} {1 p3 20 110 170}
    fmt.Println(4,up4, *up4)    // output:4 &{2 p4 0 0 180} {2 p4 0 0 180}
}

// 结构体初步体验
func structGuild() {
    p1 := User{Uid: 1, Name: "p1", Height: 180, Weight: 130}
    p2 := User{Uid: 2, Name: "p2"}
    
    // 用户1
    // 改正用户p1的身高
    p1.Height = 169
    // 取值
    fmt.Printf("p1[name:%v,height:%d,weight:%d]\n", p1.Name, p1.Height, p1.Weight)
    
    // 用户2
    p2.Height = p1.Height + 5
    p2.Weight = p1.Weight - 10
    fmt.Printf("p2[name:%v,height:%d,weight:%d]\n", p2.Name, p2.Height, p2.Weight)
}

// 嵌入结构体
func structExtend() {
    fmt.Println("=======structExtend======")
    // 错误的初始化格式
    // e := Employee{Uid: 2, Name:"p4", Height: 180, Work: 5}
    // 正确的
    e := Employee{
        User:User {
            Uid: 2, 
            Name:"p4", 
            Height: 180, 
        },
        Work:5,
    }
    fmt.Println(e.Uid, e.Name, e.Height, e.Work)
}
// 结构体与内存
func structAndMemory() {
    fmt.Println("=======structAndMemory======")
    p := Point{4, 5}
    newPtr := new(Point)
    refPtr := &p
    newPtr = &p
    fmt.Printf("newPtr[%q, %v],refPtr[%q, %v]\n", newPtr, *newPtr, refPtr, *refPtr)
}