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
	a := make([]int, 10, 20)
	fmt.Println(a, len(a), cap(a))
	a[9] = 42
	fmt.Printf("a[12]=%d\n", a[9])

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days", 7)
	b = b.Append([]byte(" ABCD"))
	b.AppendPointer([]byte(" DEEF"))
	fmt.Printf("%s", b)
}
