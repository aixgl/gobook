package main

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	inc, times := 10, 0
	for i := 1; i <= 100; i++ {
		go Deposit(inc)
		times++
	}
	time.Sleep(1 * time.Second)
	assert.Equal(t, 100, times)
	assert.Equal(t, 1000, Balance())
}
