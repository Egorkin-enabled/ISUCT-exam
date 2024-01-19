package keyboards

// Объявление структуры 'Монитор'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1990
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type keyboard struct {
	model string
	age   uint32
	size  string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IKeyboard interface {
	// Геттеры
	GetSize() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewKeyboard(model string, age uint32) (IKeyboard, error) {
	// Создаём пустой экземпляр.
	instance := keyboard{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetModel(model); error != nil {
		return nil, error
	}

	if error := instance.SetAge(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance keyboard) GetSize() string {
	return instance.size
}

func (instance keyboard) GetModel() string {
	return instance.model
}

func (instance keyboard) GetAge() uint32 {
	return instance.age
}

func (instance *keyboard) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.size = "192x108 cm"
	case "B":
		instance.size = "640x480 cm"
	case "C":
		instance.size = "256x176 cm"
	default:
		instance.size = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *keyboard) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
