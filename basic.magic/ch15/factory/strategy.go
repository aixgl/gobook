// --------------------------------------
// * 观察者
// --------------------------------------
package factory


// --------------------------------------
// 策略接口
type FormulaIF interface {
    Formula(a, b int) int
}

// --------------------------------------
// 策略模式实现
type Strategy struct {
    a   int
    b   int
    formulas    map[string]FormulaIF
    exec    func(a, b int) int
}

func (this *Strategy) Register(key string, formula FormulaIF){
    this.formulas[key] = formula
}

func (this *Strategy) Handle(args ...string) {
    this.exec =  func(a, b int) int {
        sum := 0
        for _, op:= range args {
            fn, ok := this.formulas[op]
            if !ok {
                continue
            }
            sum += fn.Formula(a, b)
        }
        return sum
    }
}

func (this *Strategy) Exec() int{
    return this.exec(this.a, this.b)
}

// --------------------------------------
// 策略公式  实现策略接口
type Add struct {}
func (this *Add) Formula(a, b int) int{
    return a + b
}

type Sub struct{}
func (this *Sub)Formula(a, b int) int{
    return a - b
}

type Multi struct{}
func (this *Multi)Formula(a, b int) int{
    return a * b
}

// --------------------------------------
// demo


