package cf

import "fmt"

var nums [5]int = [5]int{1, 2, 3, 4, 5}

func Loop() {
	for i, num := range nums {
		fmt.Println(i, num)
	}
}
