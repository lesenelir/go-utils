package main

import (
	"fmt"
	"github.com/lesenelir/go-utils/sliceutils"
)

func main() {
	fmt.Println("hello")

	sd := sliceutils.SliceData[int]{Data: []int{1, 2}, Size: 2}

	sd.Push(3)
	sd.Push(4, 5, 6)
	sd.Unshift(0)

	sd.ForEach(func(ele int, index int, slice []int) {
		fmt.Printf("ele: %d, index: %d, slice: %v\n", ele, index, slice)
	})

	sd.Data = sd.Reverse()
	fmt.Println(sd.Data)

	sd.Map(func(index int, ele int, slice []int) int {
		return ele * 2
	})
	fmt.Println(sd.Data)
}
