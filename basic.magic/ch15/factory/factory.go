// --------------------------------------
// * 工厂
// --------------------------------------
package factory

import (
    "errors"
    "fmt"
)
// 定义一个简易流程
type Executer interface{
    Start()
    Exec() bool
    End()
}

// 工厂类
type FactoryClass struct {
    flows   map[string]Executer
}

// 工厂取对象实例 方式一
func(f *FactoryClass) Register(key string, exec Executer) {
    f.flows[key] = exec
}

// 工厂取对象实例 方式一
func(f *FactoryClass) FactoryStatic(key string) (Executer, error){
    exec, ok := f.flows[key]
    if !ok {
        return nil, errors.New("Not found[" + key + "]")
    }
    return exec, nil
}

// 工厂取对象实例 方式二
// 每次new一次操作
func(f *FactoryClass) Factory(key string)(Executer) {
    switch key {
    case "demo1":
        // 实际使用中 new 完对象还需要做一系列初始化等，可单独放一个函数里
        return &Demo1{}
    case "middle":
        return &Middle{}
    }
    return nil
}

// --------------------------------------------
// 产品结构体
type Demo1 struct {
}
// 实现流程接口
func (d *Demo1) Start() {
}
func (d *Demo1) Exec() bool {
    fmt.Println("Demo1")
    return true
}
func (d *Demo1) End() {
}


type Middle struct {
}
// 实现流程接口
func (d *Middle) Start() {
}
func (d *Middle) Exec() bool {
    fmt.Println("Middle")
    return true
}
func (d *Middle) End() {
}

func init() {
    
}
/*
    //ft := &factory.FactoryClass{}
    ft := &FactoryClass{}
    exe := ft.Factory("demo1")
    exe.Start()
    exe.Exec()
    exe.End()
    
    exe = ft.Factory("middle")
    exe.Start()
    exe.Exec()
    exe.End()
 */