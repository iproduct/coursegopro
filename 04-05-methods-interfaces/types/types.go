package main

import (
	"fmt"
)

// ByteSlice type
type ByteSlice []byte

// Append appends data and returns new bvyte slice
func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}

// AppendPointer appends data to existing ByteSlice
func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This week has %d days - ", 7)
	b = b.Append([]byte("A\n"))
	b.AppendPointer([]byte("B\n"))
	b = (ByteSlice).Append(b, []byte("C\n"))
	b = (*ByteSlice).Append(&b, []byte("D\n"))
	//(ByteSlice).AppendPointer(b, []byte("E\n")) // Error: invalid method expression ByteSlice.AppendPointer (needs pointer receiver: (*ByteSlice).AppendPointer)
	(*ByteSlice).AppendPointer(&b, []byte("F\n"))
	fmt.Printf("%#v", string(b))
}
