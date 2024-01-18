package pencils

// Объявление структуры 'Фрукт'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1950
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type pencil struct {
	model string
	age   uint32
	color string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IPencil interface {
	// Геттеры
	GetColor() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewPencil(model string, age uint32) (IPencil, error) {
	// Создаём пустой экземпляр.
	instance := pencil{}

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
func (instance pencil) GetColor() string {
	return instance.color
}

func (instance pencil) GetModel() string {
	return instance.model
}

func (instance pencil) GetAge() uint32 {
	return instance.age
}

func (instance *pencil) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.color = "Blue"
	case "B":
		instance.color = "Indigo"
	case "C":
		instance.color = "Light-blue"
	default:
		instance.color = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *pencil) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
