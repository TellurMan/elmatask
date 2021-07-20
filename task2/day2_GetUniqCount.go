package yurlov

// подсчёт количества уникальных чисел в срезе
func GetUniqCount(arr []int) int {
	set := make(map[int]bool) // эмуляция множества
	for _, elem := range arr {
		set[elem] = true // если такой элемент был, то значение по ключу перезапишется; иначе создастся новое
	}
	return len(set) // размер множества = количество ключей = количество уникальных чисел
}
