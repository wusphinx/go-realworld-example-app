package testcase

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/longbridgeapp/assert"
)

type Monkey struct{}

func (m *Monkey) Foo(input int) string {
	return strconv.Itoa(input)
}

func callMonkey(x int) error {
	m := &Monkey{}
	for i := 0; i < 10; i++ {
		res := m.Foo(x + 1)
		if res == "" {
			return fmt.Errorf("get empty")
		}
	}

	return nil
}

func TestFoo(t *testing.T) {
	gomonkey.ApplyMethodFunc(reflect.TypeOf(&Monkey{}), "Foo", func(input int) string {
		switch input {
		case 1:
			return "0"
		case 2:
			return "1"
		}
		return "101"
	})

	got := callMonkey(1)
	assert.Nil(t, got)
}
