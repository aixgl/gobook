package mutex

import (
    //"fmt"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestBankChannel(t *testing.T) {
    inc, times := 10, 0
    for i := 1; i <= 100; i++ {
        go DepositChannel(inc)
        times++
    }
    time.Sleep(1 * time.Second)
    assert.Equal(t, 100, times)
    assert.Equal(t, 1000, BalanceChannel())
}

func TestBankMutex(t *testing.T) {
    inc, times := 10, 0
    for i := 1; i <= 100; i++ {
        go DepositMutex(inc)
        times++
    }
    time.Sleep(1 * time.Second)
    assert.Equal(t, 100, times)
    assert.Equal(t, 1000, BalanceMutex())
}

func TestBankRWMutex(t *testing.T) {
    inc, times := 10, 0
    for i := 1; i <= 100; i++ {
        go DepositRWMutex(inc)
        times++
        go BalanceRWMutex()
    }
    time.Sleep(1 * time.Second)
    assert.Equal(t, 100, times)
    assert.Equal(t, 10*times, BalanceRWMutex())
}
