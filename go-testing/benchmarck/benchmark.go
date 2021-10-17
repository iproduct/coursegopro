package benchmarck

import (
	"bytes"
)

func ConcatenateStr(n int, s string) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func ConcatenateStr2(n int, s string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func ConcatenateStr3(n int, s string) string {
	return s
}

