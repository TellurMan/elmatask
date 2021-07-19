package yurlov

// генератор квадратов n натуральных чисел, начиная со start
func SolutionSquareGenerator(start int, n int) []int {
	squareArray := make([]int, n)
	for i := range squareArray {
		squareArray[i] = (start + i) * (start + i)
	}
	return squareArray
}
