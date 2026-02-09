package slices

import "fmt"

func Sli() {
	aSliceLiteral := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(aSliceLiteral); i++ {
		m := aSliceLiteral[i]
		fmt.Println(m)
	}
}
