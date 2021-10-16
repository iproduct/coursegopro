package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func customError() error  {
	return errors.New("emit macho header corrupted")
}

func main() {
	err := customError()
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
}

