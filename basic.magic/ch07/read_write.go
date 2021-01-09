package main
import (
    "fmt"
    "io"
)

type space struct {
    str   string
}

// 实现io.Reader
func (sp *space) Read(p []byte) (int, error){
    n := len(sp.str)
    // 循环读取str 此处禁止使用append
    for i :=0; i < n; i++ {
        p[i] = sp.str[i]
    }
    // 等价上面循环的部分
    // n = copy(p, sp.str[:])
    return n, nil
}

// 实现io.Writer
func (sp *space) Write(p []byte)(int, error) {
    sp.str = string(p)
    return len(sp.str), nil
}