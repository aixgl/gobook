package teacher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	// t.Fatal("not implemented")
	assert.Equal(t, 1, GoIf(0))
	assert.Equal(t, 2, GoIf(1))
	assert.Equal(t, 1, ASMIf(0))
	assert.Equal(t, 2, ASMIf(1))
}

func TestFor(t *testing.T) {
	a := GoFor(2, 10, 1)
	b := ASMFor(2, 10, 1)
	assert.Equal(t, a, b)
}
