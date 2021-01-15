package main

import (
    "fmt"
    "encoding/json"
)
// Product 商品信息
type Product struct {
    Name      string    `json:"name"`
    ProductID int64     `json:"pid"`
    Number    int       `json:"num,omitempty"`
    Price     float64   `json:"price"`
    IsSale  bool        `json:"is_sale,string"`
    model     string    `json:"m,omitempty"`
}

// 简单的序列化
func jsonSimpleMarshal() {
    fmt.Println("=======jsonSimpleMarshal=======")
    p := &Product{
        Name:"ios",
        ProductID :7239123023421,
        Number:10,
        Price: 20.5,
        IsSale:false,
        model:"TCL",
    }
    data, _ := json.Marshal(p)
    fmt.Println(string(data))
}

// 反序列化
func jsonSimpleUnmarshal() {
    data := `{"Name":"ios","ProductID":7239123023421,"Number":10,"Price":20.5,"IsSale":false}`
    p := &Product{}
    err := json.Unmarshal([]byte(data), p)
    fmt.Println("return:", err, "; struct:", *p)
}

func jsonTagEmptyMarshal() {
    fmt.Println("=======jsonTagEmptyMarshal=======")
    p := &Product{
        Name:"ios",
        ProductID :7239123023421,
        Number:0,
        Price: 20.5,
        IsSale:false,
        model:"",
    }
    data, _ := json.Marshal(p)
    fmt.Println(string(data))    
}

/*--------------------------------------------------------*
 *  自定义json解析
 *
 *--------------------------------------------------------*/
type Order struct {
    pids map[string]int64
}

func (o *Order) MarshalJSON() ([]byte, error) {
	l := []int64{}
	for _,v := range o.pids {
		l = append(l,v)
	}

	return json.Marshal(l)
}

func (o *Order) UnmarshalJSON(b []byte) error {
	l := []int64{}
	err := json.Unmarshal(b, &l)
	if err != nil {
		return err
	}

	for i,v := range l {
		k := fmt.Sprintf("%d", i)
		o.pids[k] = v
	}	

	return nil
}

func jsonDefine() {
    fmt.Println("=======jsonDefine=======")
    m := map[string]int64{"1":234, "2":987,"3":110, "4":119}
    o := &Order{pids:m}
    
    byteOrder, err := json.Marshal(o)
    fmt.Println("JSON defined", "error:",err, "order:", string(byteOrder))
    
    mstr := `[234, 987, 110, 119]`
    o1 := &Order{pids:map[string]int64{}}
    err = json.Unmarshal([]byte(mstr), o1)
    fmt.Println("JSON defined unmarshal", "error:",err, "order:", o1)
}