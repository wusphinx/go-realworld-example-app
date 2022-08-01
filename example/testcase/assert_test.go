package testcase

import (
	"testing"

	la "github.com/longbridgeapp/assert"
	"github.com/stretchr/testify/assert"
)

func TestAssertInt(t *testing.T) {
	var a int64 = 1
	var b int8 = 1
	la.Equal(t, a, b)
	assert.NotEqual(t, a, b)
}
