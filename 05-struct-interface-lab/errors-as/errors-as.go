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
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
			//fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}