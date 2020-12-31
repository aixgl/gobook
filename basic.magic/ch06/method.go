package main

import (
	"fmt"
	"github.com/aixgl/ch06/order"
)

type NewString string

type Person struct {
	name string
	sex  byte
	age  int
	uid  int64
}

type Teacher struct {
	Person
	School string
	office int
}

/** =============================================
 * 结构体NewString的方法
 ** =============================================*/
func (s NewString) Print() {
	fmt.Println(s)
}

func (s NewString) Split(sep rune) []string {
	rtnSlice := []string{} // 返回的字符串切片定义
	beforeKey := 0         // 切分起始key
	sSrc := string(s)      // 转换成string类型为了使用字符串切片特性
	for i, r := range s {
		if r == sep {
			rtnSlice = append(rtnSlice, sSrc[beforeKey:i])
			beforeKey = i + 1
		}
	}
	if beforeKey >= len(s) {
		rtnSlice = append(rtnSlice, "")
	} else {
		rtnSlice = append(rtnSlice, sSrc[beforeKey:])
	}

	return rtnSlice
}

/* (ptr *Person) PtrSetName(name string) error
 * 方法运行时可以看成定义的这个函数
 */
func PtrSetName(ptr *Person, name string) error {
	ptr.name = name
	return nil
}

/** =============================================
 * 结构体Person的方法
 ** =============================================*/
// 值方法
func (pv Person) ValGetName() string {
	return pv.name
}
func (pv Person) ValSetName(name string) error {
	pv.name = name
	return nil
}

// 指针方法
func (ptr *Person) PtrGetName() string {
	return ptr.name // 等价(*ptr).name // 错误*ptr.name
}
func (ptr *Person) PtrSetName(name string) error {
	ptr.name = name
	return nil
}
func (ptr *Person) PrintInfo() {
	fmt.Println("I am", ptr.name, "and", ptr.age, "year old.")
}

/** =============================================
 * 结构体Teacher的方法
 ** =============================================*/
func (t *Teacher) GetSchool() string {
	return t.School
}
func (t *Teacher) SetSchool(name string) {
	t.School = name
}

// 覆盖重写
func (ptr *Teacher) PrintInfo() {
	fmt.Println("My school is", ptr.School)
}

//==============================================
type OrderInfo struct {
	order.User
}
