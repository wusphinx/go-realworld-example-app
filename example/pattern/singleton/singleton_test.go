package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	for i := 0; i < 10; i++ {
		GetInstance()
	}

	// count is: 1
	t.Logf("count is: %d", count)

	if count == 10 {
		t.Errorf("GetInstance() should only once")
	}
}
