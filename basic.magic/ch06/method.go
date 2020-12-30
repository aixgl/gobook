package main

import "fmt"

type NewString string

type Person struct {
    name    string
    sex     byte
    age     int
    uid     int64
}

func (s NewString) Print() {
    fmt.Println(s)
}

func (s NewString) Split(sep rune) []string{
    rtnSlice := []string{}  // 返回的字符串切片定义
    beforeKey := 0      // 切分起始key
    sSrc := string(s)   // 转换成string类型为了使用字符串切片特性
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

func (p Person) GetName() string{
    return p.name
}
