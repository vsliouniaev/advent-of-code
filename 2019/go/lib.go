package _go

import (
	"bufio"
	"os"
)

func ParseFile(file string) (data []string) {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return
}
