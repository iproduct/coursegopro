package main

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}

func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func main() {
	intslice := make([]int, 10, 20)
	fmt.Println(intslice, len(intslice), cap(intslice))

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%s", b)
}
