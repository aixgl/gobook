package main

import (
    "reflect"
    "fmt"
)

// ------------------------------
type employee struct {
    name    string  `json:"pkg" typ:"string"`
    id      uint64  `serial:"incr" unqiue:""`
}

// 非指针类型方法
func (e employee) GetName() string{
    return e.name
}

func (e employee) GetId() uint64 {
    return e.id
}

// 指针类型的方法
func (e *employee) PtrGetName() string{
    return e.name
}
func (e *employee) PtrGetId() uint64 {
    return e.id
}
func (e *employee) PtrSetName(name string) bool {
    e.name = name
    return true
}
// ------------------------------
// * 结构体 反射示例
// ------------------------------
func structReflect() {
    fmt.Println("-------------------structReflect---------------------")
    e := employee{name : "xiaoming", id :8802667}
    
    // 获取类型对象
    eType := reflect.TypeOf(e)
    // reflect.Type.Name()
    fmt.Println("结构体名称：", eType.Name())
    
    // 获取字段信息
    for i :=0; i < eType.NumField(); i++ {
        // 按索引获取StructField结构体对象 
        // func (t *commonType) Field(i int) StructField
        sk := eType.Field(i)
        skByName,_ := eType.FieldByName(sk.Name)
        fmt.Printf("The %d field name[%s] and type[%s]; tag[%v]\n", i, sk.Name, sk.Type, skByName.Tag)
    }
    
    // 获取成员方法
    for i:=0; i < eType.NumMethod(); i++ {
        m := eType.Method(i)
        mByName,_ := eType.MethodByName(m.Name)
        // m 和 mByName 是值一样的2个结构体
        fmt.Printf("The %d method name[%s] and type[%s] i==mByName.Index[%v]\n", i, m.Name, m.Type, i == mByName.Index)
    }
}

func structReflect2() {
    fmt.Println("-------------------structReflect2---------------------")
    e := employee{name : "xiaoming", id :8802667}
    
    // 获取类型对象
    eType := reflect.TypeOf(e)
    fmt.Printf("employee struct Name[%s] Kind[%s]\n", eType.Name(), eType.Kind() )
    fmt.Printf("employee struct Size[%d] String[%s]\n", eType.Size(), eType.String())
    
    // 嵌入查询函数 
    // FieldByNameFunc 的参数
    sf := func (s string) bool {
        if s == "name" {
            return true
        }
        return false
    }
    
    // 依据查询条件 查询字段
    fieldStruct, ok := eType.FieldByNameFunc(sf)
    fmt.Println("employee field:", fieldStruct.Name,"type:", fieldStruct.Type, ";exist:", ok)
}

func structReflectElem() {
    fmt.Println("-------------------structReflectElem---------------------")
    e := &employee{name : "xiaoming", id :8802667}
    
    // 获取类型对象
    eType := reflect.TypeOf(e)
    // 错误的获取方式
    fmt.Printf("employee ptr Size[%d] Type[%s]\n", eType.Size(), eType.Name() )
    
    // 通过reflect.Type.Elem 获取指针的值反射类型
    // 与*e 对应
    valueType := eType.Elem()
    fmt.Printf("employee ptr[Elem] NumMethod[%d] Type[%s]\n", eType.NumMethod(), valueType.Name() )
    // 获取值类型成员方法 
    for i:=0; i < valueType.NumMethod(); i++ {
        m := valueType.Method(i)
        mByName,_ := valueType.MethodByName(m.Name)
        // m 和 mByName 是值一样的2个结构体
        fmt.Printf("Value %d method name[%s] and type[%s] i==mByName.Index[%v]\n", i, m.Name, m.Type, i == mByName.Index)
    }   
    // 获取指针类型的成员方法
    for i:=0; i < eType.NumMethod(); i++ {
        m := eType.Method(i)
        mByName,_ := eType.MethodByName(m.Name)
        // m 和 mByName 是值一样的2个结构体
        fmt.Printf("Ptr %d method name[%s] and type[%s] i==mByName.Index[%v]\n", i, m.Name, m.Type, i == mByName.Index)
    }    
}

// ------------------------------------------------
// * 通过反射修改值
type Company struct {
    Name string
}

func (c Company) Echo(s string) {
    fmt.Println("company echo:", s)
}
func structChangeValue() {
    fmt.Println("-------------------structChangeValue---------------------")
    e := &Company{Name : "xiaoming"}

    v := reflect.ValueOf(e)

    // 修改值必须是指针类型否则不可行
    if v.Kind() != reflect.Ptr {
        fmt.Println("只有引用类型")
        return
    }

    // 获取指针所指向的元素
    v = v.Elem()
    // 只能是对外开放的字段
    name := v.FieldByName("Name")
    name.SetString("小明")
    fmt.Println("change struct value", e)
    
    // 变形一下 整型
    v_int := 100
    v_int_reflect := reflect.ValueOf(&v_int)
    v_int_reflect.Elem().SetInt(200)
    fmt.Println("changed int type value", v_int)
}

// * 调用方法
func reflectCall() {
    fmt.Println("-------------------reflectCall---------------------")
    e := Company{Name : "xiaoming"}
    v := reflect.ValueOf(e)
    
    // 生成reflect.Value 获取Echo的控制
    mt := v.MethodByName("Echo")
    
    // 拼凑参数 这算是最不容易记住的地方了
    args := []reflect.Value{reflect.ValueOf("Google!")}
    // 调用
    mt.Call(args)
    fmt.Println("mt:", mt)
}