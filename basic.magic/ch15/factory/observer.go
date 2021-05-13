// --------------------------------------
// * 观察者
// --------------------------------------
package factory

import(
    "fmt"
    "container/list"
)

// 观察者模型接口
type SubjectIF interface{
    Add(ObserverIF)
    Send()
}

type ObserverIF interface {
    Receive(info string)
}

// ----------------------------------------
// topic
type Topic struct {
    title string
    ol  *list.List
}

func (t *Topic) Add(o ObserverIF) {
    t.ol.PushBack(o)
}

func (t *Topic) Send() {
    for ov := t.ol.Front(); ov != nil; ov = ov.Next() {
        (ov.Value).(ObserverIF).Receive(t.title)
    }
}
// 热点主题
type HotSubject struct {
    Topic
}
// 新主题
type NewSubject struct {
    Topic
}

// ----------------------------------------
// 观察者
type Observer struct {
    name    string
}
func (o *Observer) Receive(info string) {
    fmt.Println("主题[", info, "] 观察者[", o.name, "]")
}
// 新闻观察者
type NewsObserver struct{
    Observer
}
// 热点观察者
type HotObserver struct {
    Observer
}
