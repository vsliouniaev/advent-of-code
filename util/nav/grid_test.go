package nav

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	g := grid{
		{1, 2},
		{3, 4},
	}
	i := g.GetIterator()
	for p := i.Next(); p != nil; p = i.Next() {
		fmt.Printf("%s = %d\n", p, g.Get(p).(int))
	}

}
