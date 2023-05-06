package codegenerator

import (
	"crypto/rand"
	"io"
)

var numbers = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func RandomEncodedNumbers(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = numbers[int(b[i])%len(numbers)]
	}
	return string(b)
}
