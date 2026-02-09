package array

import "fmt"

func Arr() {
	anArray := [4]int{1, 2, 3, 4}
	twoD := [5][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	threeD := [3][2][1]int{{{43}, {76}}, {{78}, {89}}, {{66}, {88}}}

	fmt.Println("The length of ", anArray, "is", len(anArray))
	fmt.Println("The First Element of", twoD, "is", twoD[0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		// fmt.Println(v)
		for j := 0; j < len(v); j++ {
			m := v[j]
			// fmt.Println(m)
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}

	// The range keyword does exactly the same job as the iteration variables used in the for
	// loops of the preceding code segment, but it does so in a more elegant and clearer way.
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}
