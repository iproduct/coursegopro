package main

import (
"errors"
"fmt"
"io/fs"
"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		fmt.Printf("%T, %v\n", err, err)
		cause := errors.Unwrap(err)
		fmt.Printf("Unwrapped: %T, %[1]v\n", cause)
		cause = errors.Unwrap(cause)
		fmt.Printf("Unwrapped: %T, %[1]v\n", cause)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}