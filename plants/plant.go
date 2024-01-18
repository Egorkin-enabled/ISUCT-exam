package albums

// Объявление структуры 'Растение'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1700
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type plant struct {
	name string
	age  uint32
	kind string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IPlant interface {
	// Геттеры
	GetKind() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewPlant(name string, age uint32) (IPlant, error) {
	// Создаём пустой экземпляр.
	instance := plant{}

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
func (instance plant) GetKind() string {
	return instance.kind
}

func (instance plant) GetName() string {
	return instance.name
}

func (instance plant) GetAge() uint32 {
	return instance.age
}

func (instance *plant) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.kind = "CumInSinA"
	case "B":
		instance.kind = "SibIsIna"
	case "C":
		instance.kind = "NajEbina"
	default:
		instance.kind = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *plant) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
