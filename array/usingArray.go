package array

import "fmt"

func Arr() {
	anArray := [4]int{1, 2, 3, 4}
	twoD := [5][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	threeD := [3][2][1]int{{{43}, {76}}, {{78}, {89}}, {{66}, {88}}}
	fmt.Println("The length of ", anArray, "is", len(anArray))
	fmt.Println("The First Element of", twoD, "is", twoD[0])
	fmt.Println("The length of", threeD, "is", len(threeD))
}
