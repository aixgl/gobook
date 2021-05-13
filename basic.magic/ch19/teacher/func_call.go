package teacher

import (
	"fmt"
	//"runtime"
)

// go 实现
func GoAddSub(a, b int) (c, d int) {
	return a + b, a - b
}

func GoCaller() {
	// go demo
	c, d := GoAddSub(90, 70)
	fmt.Println("GO:", c, d)

	// 验证汇编函数调度
	cc, dd := ASMCaller()
	fmt.Println("ASM:", cc, dd)
}

// 汇编实现
func ASMAddSub(a, b int) (c, d int)

func ASMCaller() (c, d int)
