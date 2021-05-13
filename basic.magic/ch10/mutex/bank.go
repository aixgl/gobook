package mutex

import "sync"

// -------------------------------------------
// sync.Mutex
var (
    balanceMu int
    mu        sync.Mutex
)

// 存款
func DepositMutex(amount int) {
    mu.Lock()
    // 通常会用延迟函数调度Unlock：
    // defer mu.Unlock()
    balanceMu = balanceMu + amount
    mu.Unlock()
}

func BalanceMutex() int {
    mu.Lock()
    // defer mu.Unlock()
    b := balanceMu
    mu.Unlock()
    return b
}

// -------------------------------------------
// channel lock
var balance int

// 用channel 做同步控制
// channel的缓存一定是1才可以。
var lock_ch = make(chan int, 1)

// 存款
func DepositChannel(amount int) {
    // lock_ch 是空第一次发送 不会阻塞，因为缓存是1；
    // 若缓存满了则会阻塞
    lock_ch <- 1
    // 赋值完成取出token，释放锁定
    balance = balance + amount
    <-lock_ch
}

// 余额
func BalanceChannel() int {
    lock_ch <- 1
    b := balance
    return b
}

// -------------------------------------------
// 读写锁sync.RWMutex
var (
    balanceRW int
    muR       *sync.RWMutex = new(sync.RWMutex)
)

// 存款
func DepositRWMutex(amount int) {
    muR.Lock()
    defer muR.Unlock()
    balanceRW = balanceRW + amount
}

func BalanceRWMutex() int {
    muR.RLock()
    defer muR.RUnlock()
    return balanceRW
}
