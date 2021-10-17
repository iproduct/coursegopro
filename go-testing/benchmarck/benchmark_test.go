package benchmarck

import (
	"fmt"
	"testing"
)

func BenchmarkConcatenateStr1(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
		j++
		ConcatenateStr(10000, "a")
	}

	fmt.Println(j)
}

func BenchmarkConcatenateStr2(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
		j++
		ConcatenateStr2(10000, "a")
	}
	fmt.Println(j)
}

func BenchmarkConcatenateStr3(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
		j++
		ConcatenateStr3(10000, "a")
	}
	fmt.Println(j)
}