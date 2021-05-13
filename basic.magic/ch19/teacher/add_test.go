package teacher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// t.Fatal("not implemented")
	assert.Equal(t, int64(20), Add(5, 15))
}

func TestMyNum(t *testing.T) {
	assert.Equal(t, 8888, MyNum)
	assert.Equal(t, 4, intMov2)
	assert.True(t, trueBool)
	assert.False(t, falseBool)
}

func TestAssempleString(t *testing.T) {
	assert.Equal(t, "gopher and phper", Job)
}

func TestAssempleSlice(t *testing.T) {
	assert.Equal(t, 4, len(helloSlice))
	assert.Equal(t, 6, cap(helloSlice))
	assert.Equal(t, int32(3), helloSlice[0])
	assert.Equal(t, int32(4), helloSlice[1])
	assert.Equal(t, int32(64), helloSlice[2])
	assert.Equal(t, int32(21), helloSlice[3])
}

func TestSwap(t *testing.T) {
	a, b := Swap(10, 20)
	assert.Equal(t, 10, b)
	assert.Equal(t, 20, a)
}
