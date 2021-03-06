package yurlov

// поиск длины самой длинной последовательности нулей в двоичном представлении целого числа
func SolutionBinaryGap(N int) int {
	checkInt32func := func(N int) func() int { // обработка блока размером int32
		maxValue := 0 // текущий максимум
		curValue := 0 // для подсчёта нолей

		return func() int {
			if N != 0 { // выход если подан 0; или в случае если был int32 и это второй вызов
				for i := 0; i < 32; i++ { // цикл с проверкой младшего бита
					if N&1 == 0 { // если ноль, то
						curValue++ // подсчёт нолей
					}
					if N&1 == 1 { // если один, то
						if curValue > maxValue { // если подсчитано нолей больше текущего максимума, то
							maxValue = curValue // обновление максимума
						}
						curValue = 0 // очистить подсчёт нолей при встрече единицы
					}
					N >>= 1 // сдвиг числа для обработки следующего младшего бита
				}
			}
			return maxValue
		}
	}

	checkInt32 := checkInt32func(N) // состояние числа N и его свойств сохранится
	checkInt32()                    // обработка младших 32 бит
	maxValue := checkInt32()        // если int64, то обработка старших 32 бит, иначе сразу возврат ответа

	return maxValue
}
