package drinks

// Объявление структуры 'Напиток'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1000
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type drink struct {
	name string
	age  uint32
	kind string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IDrink interface {
	// Геттеры
	GetKind() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewDrink(name string, age uint32) (IDrink, error) {
	// Создаём пустой экземпляр.
	instance := drink{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(name); error != nil {
		return nil, error
	}

	if error := instance.SetAge(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance drink) GetKind() string {
	return instance.kind
}

func (instance drink) GetName() string {
	return instance.name
}

func (instance drink) GetAge() uint32 {
	return instance.age
}

func (instance *drink) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.kind = "Water"
	case "B":
		instance.kind = "Soda"
	case "C":
		instance.kind = "Alcahol"
	default:
		instance.kind = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *drink) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
