package main

import (
	"fmt"
	"github.com/iproduct/coursegopro/04-struct-methods-lab/copymap"
)

type Vertex struct {
	Lat  float64
	Long float64
}

func main() {
	var m copymap.GenericMap = make(copymap.GenericMap)
	m["Bell Labs"] = Vertex{Lat: 40.68433, Long: -74.39967}
	m2 := copymap.GenericMap{
		"Google":    Vertex{37.42202, -122.08408},
		"Microsoft": Vertex{52.68433, -49.39967},
	}
	copymap.CopyMap(m, m2)
	for k, v := range m {
		fmt.Printf("%s -> %#v\n", k, v)
	}

}
