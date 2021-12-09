package util

import (
	"testing"
)

func TestRotation(t *testing.T) {
	o := Orientation{}
	o.Turn(Left)
	if o.GetDirection() != North {
		t.Errorf("Expected North but got %s", o.GetDirection())
	}
	o.Turn(Left).Turn(Left).Turn(Left)
	if o.GetDirection() != East {
		t.Errorf("Expected East but got %s", o.GetDirection())
	}
	o.Turn(Right)
	if o.GetDirection() != South {
		t.Errorf("Expected South but got %s", o.GetDirection())
	}
}

func TestDirection(t *testing.T) {
	o := Orientation{}
	v := o.Forward()
	if v.X != 1 || v.Y != 0 {
		t.Errorf("Unexpected East move is x: %d y: %d", v.X, v.Y)
	}
	v = o.Backward()
	if v.X != -1 || v.Y != 0 {
		t.Errorf("Unexpected West move is x: %d y: %d", v.X, v.Y)
	}
}
