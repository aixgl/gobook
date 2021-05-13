// --------------------------------------
// * 函数式
// --------------------------------------
package factory

import(
    "fmt"
    "strconv"
)

// --------------------------------------
// 因子
func sliceEach(s_arr []string, fn func(k int, v string) ) {
    if s_arr == nil {
        return
    }
    for i, v := range s_arr {
        fn(i, v)
    }
}

func demoSliceEach() {
    arr := []string{"house", "room", "city"}
    str := ""
    sliceEach(arr, func(k int, v string){
        str += strconv.Itoa(k+1) +":"+ v +";"
    })
    fmt.Println("funcEach:", str)
}

// 柯里化
func funcAdd(x int) func(y int) int {
    return func(y int) int {
        return x + y
    }
}
func demoFuncAdd() {
    add := funcAdd(2)
    fmt.Println("funcAdd:", add(3), add(1))
}


// 尾化递归
// N向下阶乘 n*n-1 * n-2 ... 1
func recursion(num int) int {
    if num == 0 {
        return 1
    } else {
        return num * recursion(num - 1)
    }
}
func tailRecursion(accumulator, num int) int {
    if num == 0 {
        return accumulator
    }
    return tailRecursion(accumulator * num, num - 1)
}

func demoRecursion() {
    fmt.Println("recursion:", recursion(10), tailRecursion(1, 10))
}

// 斐波那契
func fibonacci () func () int{
    a, b := 0, 1
    return func () int {
        a, b = b, a+b
        return a 
    }
}
func demoFib() {
    f := fibonacci ()
    for i :=0; i< 20; i++ {
        fmt.Printf("%d", f())
    }
    fmt.Println("")
}