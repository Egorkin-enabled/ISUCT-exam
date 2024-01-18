package books

// Опреации над срезами

// Неопределённоть:
//     'метод для вычисления среднего веса'
//     но
//     'возвращает сумму весов'
// Разделим задачу на две.

// Заготовим ошибки.
type stringError string

const AvgZeroCollectionError stringError = "Collection is empty! Unable to calculate average value."

func (error stringError) Error() string {
	return string(error)
}

// Вычисление суммы.
func CalculateAgeSum(collection []IBook) uint32 {
	var result uint32 = 0
	for _, v := range collection {
		result += v.GetAge()
	}

	return result
}

// Вычисление среднего.
func CalculateAgeAvg(collection []IBook) (float64, error) {
	if len(collection) == 0 {
		return 0.0, AvgZeroCollectionError
	}

	avg := float64(CalculateAgeSum(collection)) / float64(len(collection))
	return avg, nil
}

// Добавление уникальной структуры в срез
func TryAddUniqueInstance(collection *[]IBook, instance IBook) bool {
	for _, v := range *collection {
		// Сравниваем значения попарно,
		// т.к.  интерфейс - указатель и оператор == сравнивает адреса
		// Проверку kind упкскаем, т.к. kind зависит от проверяемых значений
		if v.GetName() == instance.GetName() && v.GetAge() == instance.GetAge() {
			// Нашли такую-же структуру - выходим.
			return false
		}
	}

	// Не нашли, значит структуру можно добавить -- повторений не будет.
	*collection = append(*collection, instance)
	return true
}
