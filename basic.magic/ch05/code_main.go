/** 注释
 * 复合类型.
 * module:github.com/aixgl/ch05/main
 * source: https://github.com/aixgl/gobook/tree/master/basic.magic/ch05/code_main
 */
package main

import (
	"fmt"
)

func arrayExample() {
	fmt.Println("=======arrayExample======")
	// 初始化 和 赋值
	//var a1 [3]int
	var a2 [3]int = [3]int{1, 2, 3}

	// 长度是根据初始化的值来确定的
	a3 := [...]int{1, 2, 4}

	for i, v := range a2 {
		a2[i] = v * 2
		fmt.Printf("a2 k:=%d, v:=%d, a2[i]:=%v\n", i, v, a2[i])
	}

	for i, n := 0, len(a3); i < n; i++ {
		fmt.Printf("a3[%d]:=%d \n", i, a3[i])
	}
}

func reverse(s1 []string) {
	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
}
func sliceFirstBlood() {
	fmt.Println("=======arrayExample======")
	s1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	//s2 := s1[3:8]
	s3 := s1[5:7]
	//s4 := s1[4:10]

	// 用range访问slice
	for i, v := range s3 {
		fmt.Printf("slice range k:%d, v:%s;\n", i, v)
	}
	reverse(s1)
	fmt.Println(s1)
}

func sliceAppend() {
	fmt.Println("=======sliceAppend======")
	var sRunes []rune
	for _, r := range "hello, world" {
		sRunes = append(sRunes, r)
	}
	fmt.Printf("append []rune%q\n", sRunes)

	sRunes = append(sRunes, sRunes[0:5]...)
	fmt.Printf("append []rune%q\n", sRunes)
}

func sliceStack() {
	fmt.Println("=======sliceStack======")
	var sb = []int{1, 2, 4, 8, 16, 32}
	// 入栈
	sb = append(sb, 64)
	fmt.Println(sb) //[1 2 4 8 16 32 64]
	// 出栈
	k := sb[len(sb)-1]
	sb = sb[:len(sb)-1]
	fmt.Println(k, sb) //64 [1 2 4 8 16 32]
}

func mapFirstBlood() {
	fmt.Println("=======mapFirstBlood======")
	scoreMap := map[string]int{ // 声明且初始化键值
		"a": 75, // 键：值
		"b": 86,
	}
	// 增
	scoreMap["c"] = 90
	fmt.Printf("%v\n", scoreMap) // output: map[a:75 b:86 c:90]
	// 删
	delete(scoreMap, "b")
	fmt.Printf("%v\n", scoreMap) // output: map[a:75 c:90]
	// 改
	scoreMap["c"] = 92
	scoreMap["a"] = 100
	fmt.Printf("%v\n", scoreMap) // output: map[a:100 c:92]
	// 查
	fmt.Println("search:", scoreMap["a"], scoreMap["c"]) // output: search: 100 92
	for k, v := range scoreMap {
		fmt.Println(k, ":=", v)
	}

	// 安全查 ok:true/false
	c, ok := scoreMap["c"]
	if ok {
		c += 25
	}
}

func main() {
	arrayExample()
	sliceFirstBlood()
	sliceAppend()
	sliceStack()
	mapFirstBlood()
	dialogWithStdin()
}
