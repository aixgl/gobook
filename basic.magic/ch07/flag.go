package main

import (
    "flag"
    "fmt"
)

var port = flag.Int("port", 808, "help message for port")

func parseFlag() {
    fmt.Println("The process port is", *port)
}