package yachts

// Объявление структуры 'Яхта'

import "fmt"

// Константы для значения

type YachtKind uint8

const (
	Reliable YachtKind = iota
	NonReliable
	Unknown
)

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1900
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type yacht struct {
	model string
	age   uint32
	kind  YachtKind
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IYacht interface {
	// Геттеры
	GetKind() YachtKind
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewYacht(model string, age uint32) (IYacht, error) {
	// Создаём пустой экземпляр.
	instance := yacht{}

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
func (instance yacht) GetKind() YachtKind {
	return instance.kind
}

func (instance yacht) GetModel() string {
	return instance.model
}

func (instance yacht) GetAge() uint32 {
	return instance.age
}

func (instance *yacht) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.kind = Reliable
	case "B":
		instance.kind = NonReliable
	case "C":
		instance.kind = NonReliable
	default:
		instance.kind = Unknown
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *yacht) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
