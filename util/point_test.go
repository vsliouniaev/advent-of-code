package util

import (
	"testing"
)

func TestRotate(t *testing.T) {
	p := Point{X: 10, Y: 4}
	p.Rotate(Right)
	if p.X != -10 && p.Y != 4 {
		t.Errorf("Expected right turn to perform 90 degrees clockwise")
	}
}
