package main

import (
	"fmt"
)

func main() {
	for {
		msg := ""
		fmt.Print("问:")
		fmt.Scan(&msg)
		fmt.Println("答:", handle(msg))
		if msg == "q" || msg == "Q" {
			return
		}
	}
}

func handle(msg string) string {
	ni := []rune("您你我")
	ma := []rune("吗")
	msgArr := []rune(msg)
	for i, v := range msgArr {
		if v == ni[0] || v == ni[1] {
			msgArr[i] = ni[2]
		}

		if v == ma[0] {
			msgArr = msgArr[:i]
			break
		}
	}
	msg = string(msgArr)
	return msg
}
