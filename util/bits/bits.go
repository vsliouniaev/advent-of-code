package bits

var lookup = []int64{0b0, 0x1, 0x3, 0x7, 0xF, 0x1F, 0x3F, 0x7F, 0xFF, 0x1FF, 0x3FF, 0x7FF, 0xFFF, 0x1FFF, 0x3FFF, 0x7FFF, 0xFFFF}

// https://stackoverflow.com/a/3847525

func GetAsInt(bytes []byte, offset int, ln int) int64 {
	byteIndex := offset / 8
	bitIndex := offset % 8
	var val int64
	if (bitIndex + ln) > 16 {
		val = int64(bytes[byteIndex]) << 16
		if byteIndex+1 < len(bytes) {
			val |= int64(bytes[byteIndex+1]) << 8
		}
		if byteIndex+2 < len(bytes) {
			val |= int64(bytes[byteIndex+2])
		}
		val = val >> (24 - bitIndex - ln)
	} else if (offset + ln) > 8 {
		val = int64(bytes[byteIndex]) << 8
		if byteIndex+1 < len(bytes) {
			val |= int64(bytes[byteIndex+1])
		}
		val = val >> (16 - bitIndex - ln)
	} else {
		val = int64(bytes[byteIndex]) >> (8 - offset - ln)
	}

	return val & lookup[ln]
}
