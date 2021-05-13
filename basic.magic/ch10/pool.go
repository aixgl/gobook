package main

import (
    "fmt"
    "sync"
    "sort"
)

// ------------------------------------------------------
// sync.Pool 的使用
var ctxPool = sync.Pool{
    New: func() interface{} {
        return "pool str"
    },
}

func poolDemo() {
    // 获取一个初始的对象
    str := ctxPool.Get().(string)
    // 使用str
    fmt.Println(str)
    // reset str
    // 将一个还原好的对象放入pool中
    ctxPool.Put(str)
}

// ------------------------------------------------------
// 缓存map
// 缓存map
type Memo struct {
	count int
	keys  []string
	hash  map[string]interface{}
	lock  sync.RWMutex
}

// 添加kv键值对
func (this *Memo) Set(k string, v interface{}) {
	this.lock.Lock()
	if _, ok := this.hash[k]; !ok {
		this.keys = append(this.keys, k)
		sort.Strings(this.keys)
		this.count++
	}
	this.hash[k] = v
	this.lock.Unlock()
}

// 获取数据长度
func (this *Memo) Count() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.count
}

// 由key检索value
func (this *Memo) Get(k string) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	v, ok := this.hash[k]
	return v, ok
}

// 根据key排序，返回有序的vaule切片
func (this *Memo) Values() []interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	vals := make([]interface{}, this.count)
	for i := 0; i < this.count; i++ {
		vals[i] = this.hash[this.keys[i]]
	}
	return vals
}

// 更近一步，统一下调度
type CallFunc func(key string) (interface{}, error)
type MemoFun struct {
    mem     *Memo
    f       CallFunc
}

func (self *MemoFun) Get(url string) (interface{}, error) {
    data, ok := self.mem.Get(url)
    if ok {
        return data, nil
    }
    data, err := self.f(url)
    if err != nil {
        return nil, err
    }
    self.mem.Set(url, data)
    return data, nil
}
// 新生成一个方法的缓存
func NewMemoFun(f CallFunc) *MemoFun {
    return &MemoFun{
        mem : &Memo{
            keys:[]string{},
            hash:map[string]interface{}{},
        },
        f   : f,
    }
}


