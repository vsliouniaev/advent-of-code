package bits

import (
	"fmt"
	"testing"
)

func TestGetBitSeqAsInt(t *testing.T) {
	bytes := []byte{0b00000000, 0b00000000, 0b11111111}
	v := GetAsInt(bytes, 2, 16)
	fmt.Printf("%b\n", v)
}
