package main

import (
	"fmt"
	"github.com/lesenelir/go-utils/maputils"
)

func main() {
	md := maputils.MapData[string, int]{
		Map:  map[string]int{},
		Size: 0,
	}

	md.Set("a", 1)
	md.Set("b", 2)
	md.Set("c", 3)
	fmt.Println(md.Map, md.Size)

	md.Delete("b")

	for key, value := range md.Map {
		fmt.Println(key, value)
	}

	md.Clear()
	fmt.Println(md.Map)
}
