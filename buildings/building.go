package buildings

// Объявление структуры 'Монитор'

import "fmt"

// Константы для типов зданий
type BuildingClass uint8

const (
	BuildingI BuildingClass = iota
	BuildingII
	BuildingIII
	BuildingIV
)

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1700
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type building struct {
	address string
	age     uint32
	class   BuildingClass
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IBuilding interface {
	// Геттеры
	GetClass() BuildingClass
	GetAddress() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAddress(uint32) error
}

// Реализация конструктора
func NewBuilding(model string, age uint32) (IBuilding, error) {
	// Создаём пустой экземпляр.
	instance := building{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetModel(model); error != nil {
		return nil, error
	}

	if error := instance.SetAddress(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance building) GetClass() BuildingClass {
	return instance.class
}

func (instance building) GetAddress() string {
	return instance.address
}

func (instance building) GetAge() uint32 {
	return instance.age
}

func (instance *building) SetModel(name string) error {
	instance.address = name

	switch name {
	case "A":
		instance.class = BuildingI
	case "B":
		instance.class = BuildingII
	case "C":
		instance.class = BuildingIII
	default:
		instance.class = BuildingIV
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *building) SetAddress(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
