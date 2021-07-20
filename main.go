package main

import (
	//yurlov "elma/task1"
	yurlov "elma/task2"
	"fmt"
)

func main() {
	/*
		fmt.Println(yurlov.SolutionSquareGenerator(4, 3))
		fmt.Println(yurlov.SolutionBinaryGap(-9_223_363_237_502_777_052))
	*/

	arr := []int{1, 3, 4, 5, 8, 9}
	arr2 := yurlov.ArrayRotator(arr, -9)
	arr2 = yurlov.ArrayRotator(arr2, 8)
	arr2 = yurlov.ArrayRotator(arr2, 0)
	arr2 = yurlov.ArrayRotator(arr2, 2)
	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(arr2)

	arr3 := []int{3, 8, 9, 7, 3, 5, 8, 9, 6}
	res := yurlov.GetUniqCount(arr3)
	fmt.Println(res)

}
