package main

import (
  "bytes"
  "io/ioutil"
  "log"
  "math/rand"
  "net/http"
 
  _ "net/http/pprof"
)

func httpPProfServer() {
    http.HandleFunc("/test", ProfHandler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}

func ProfHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if nil != err {
        w.Write([]byte(err.Error()))
        return
    }
    _ = logicOne(100000)
    buff := responseBytes(100000)
    b, err := ioutil.ReadAll(buff)
    if nil != err {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write(b)
}

func logicOne(times int) []byte{
    data := []byte("Xiao ming hao huai !!!")
    rtnData := []byte{}
    for i := 0; i < times; i++ {
        rtnData = append(rtnData, data...)
    }
    return rtnData
}
 
func responseBytes(times int) *bytes.Buffer {
    var buff bytes.Buffer
    for i := 1; i < times; i++ {
        buff.Write([]byte{'0' + byte(rand.Intn(10))})
    }
    return &buff
}