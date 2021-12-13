package nav

import (
	"testing"
)

func TestRotate(t *testing.T) {
	p := Point{X: 10, Y: 4}
	p.Rotate(Right)
	if p.X != 4 && p.Y != -10 {
		t.Errorf("Expected right turn to perform 90 degrees clockwise")
	}
}
