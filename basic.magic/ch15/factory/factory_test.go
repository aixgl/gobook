package factory

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
    "container/list"
)
// 工厂测试
func TestDemoFactory(t *testing.T) {
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
    t.Log("factory ok")
}

// 观察者测试
func TestDemoPubSub(t *testing.T) {
    new1 := &NewsObserver{}
    new1.name = "新闻1号"
    new2 := &NewsObserver{}
    new2.name = "新闻2号"
    hot1 := &NewsObserver{}
    hot1.name = "热点1号"
    
    newTopic := &NewSubject{
        Topic: Topic {
            title: "写程序成明星了，怎么办？",
            ol : list.New(),
        },
    }
    
    newTopic.Add(new1)
    newTopic.Add(new2)
    newTopic.Add(hot1)
    newTopic.Send()
    
    hotTopic := &HotSubject{
        Topic: Topic {
            title: "原来都是假的！",
            ol : list.New(),
        },
    }
    hotTopic.Add(hot1)
    hotTopic.Send()
    t.Log("Observer ok")
}
// 策略模式
func TestDemoStrategy(t *testing.T) {
    st := &Strategy{
        a : 2,
        b : 3,
        formulas: make(map[string]FormulaIF),
    }
    st.Register("+", &Add{})
    st.Register("-", &Sub{})
    st.Register("*", &Multi{})
    
    st.Handle("+", "*", "-")
    fmt.Println("a = 2, b=3;(a+b) + (a * b) + (a - b) :=", st.Exec())
    assert.Equal(t, 10, st.Exec())
    t.Log("Strategy ok")
}
// 函数式编程
func TestFuncProgame(t *testing.T) {
    demoSliceEach()
    demoFuncAdd()
    assert.Equal(t, recursion(10), tailRecursion(1, 10) )
    demoFib()
    t.Log("函数式编程")
}