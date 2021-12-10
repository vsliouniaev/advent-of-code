package data

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stk := &Stack{}
	stk.Push('[')
	stk.Push('{')
	p := fmt.Sprintf("%c%c", stk.Pop(), stk.Pop())
	if p != "{[" {
		t.Errorf(p)
	}

}
