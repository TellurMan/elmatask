package yurlov

// циклично сдвигает срез на count элементов в сторону старших индексов
func ArrayRotator(array []int, count int) []int {
	delim := (len(array) - count%len(array)) % len(array) // индекс-разделитель делит срез на 2 части, которые меняются местами

	if delim == 0 {
		tmp := make([]int, len(array)) // копирование и возврат текущего среза, так как в случае delim = 0
		copy(tmp, array)               // в append новый срез не будет создан -> возврат ссылки на старый срез (а это плохо)
		return tmp
	}

	return append(array[delim:], array[:delim]...)
}
