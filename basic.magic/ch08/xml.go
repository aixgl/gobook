package main

import (
    "fmt"
    "encoding/xml"
)
type Computer struct {
    //XMLName xml.Name `xml:"computer"` // XMLName标签名
    Id  int     `xml:"id"`
    Memory  int `xml:"memory"`
    Cpu     int `xml:"cpu"`
    Keyboard bool   `xml:"keyboard"`
    Display string  `xml:"display"`
}

func XmlMarshal() {
    fmt.Println("--------XmlMarshal--------")
    c := &Computer {
        Id:82032,
        Memory:16,
        Cpu:4,
        Keyboard:true,
        Display:"LCD",
    }
    
    output, err := xml.MarshalIndent(c, "  ", "    ")
    fmt.Println("xml-marshal")
    fmt.Printf("%v\n", string(output))
    fmt.Println("error:", err)
}

func XmlUnmarshal() {
    fmt.Println("--------XmlUnmarshal--------")
    data := `
    <?xml version="1.0" encoding="UTF-8" ?>
    <computer>
        <id>82032</id>
        <memory>16</memory>
        <cpu>4</cpu>
        <keyboard>true</keyboard>
        <display>LCD</display>
    </computer>
    `
    c := Computer{}
    xml.Unmarshal([]byte(data), &c)
    fmt.Println(c)
}