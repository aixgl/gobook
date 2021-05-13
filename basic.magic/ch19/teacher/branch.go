package teacher

func GoIf(ok int) int {
	if ok == 0 {
		return 1
	}
	return 2
}

func ASMIf(ok int) int

// go 模拟for 逻辑
func GoFor(startLoop, endLoop, step int) (ret int) {
	for i := startLoop; i < endLoop; i += step {
		ret += i
	}
	return ret
}

/*
func GoFor(startLoop, endLoop, step int) (ret int){
LOOP_BEFORE:
	ret = 0
	i = 0

LOOP_IF:
	if i < endLoop
		goto LOOP_BODY
	goto LOOP_END

LOOP_BODY:
	i += step
	ret += step
	goto LOOP_IF

LOOP_END:
	return ret
}
*/
// go汇编实现 for
func ASMFor(startLoop, endLoop, step int) (ret int)
