package furnishings

// Объявление структуры 'Мебель'

import "fmt"

// Константы для значения

type FurnishingKind uint8

const (
	Reliable FurnishingKind = iota
	NonReliable
	Unknown
)

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1800
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type furnishing struct {
	material string
	age      uint32
	kind     FurnishingKind
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IFurnishing interface {
	// Геттеры
	GetKind() FurnishingKind
	GetMaterial() string
	GetAge() uint32

	// Сеттеры
	SetMaterial(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewFurnishing(model string, age uint32) (IFurnishing, error) {
	// Создаём пустой экземпляр.
	instance := furnishing{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetMaterial(model); error != nil {
		return nil, error
	}

	if error := instance.SetAge(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance furnishing) GetKind() FurnishingKind {
	return instance.kind
}

func (instance furnishing) GetMaterial() string {
	return instance.material
}

func (instance furnishing) GetAge() uint32 {
	return instance.age
}

func (instance *furnishing) SetMaterial(name string) error {
	instance.material = name

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

func (instance *furnishing) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
