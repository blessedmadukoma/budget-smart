package number

import (
	"crypto/rand"
	"io"
	"strconv"
)

func UintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

func StringToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return uint(i), err
	}
	return uint(i), nil
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateRandom(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
