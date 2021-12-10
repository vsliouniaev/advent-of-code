package util

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stk := &Stack{}
	stk.Push('[')
	stk.Push('{')
	fmt.Printf("%c%c\n", stk.Pop(), stk.Pop())

}
