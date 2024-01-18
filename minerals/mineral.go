package minerals

// Объявление структуры 'Минерал'

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
type mineral struct {
	name string
	age  uint32
	kind string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IMineral interface {
	// Геттеры
	GetKind() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewDrink(name string, age uint32) (IMineral, error) {
	// Создаём пустой экземпляр.
	instance := mineral{}

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
func (instance mineral) GetKind() string {
	return instance.kind
}

func (instance mineral) GetName() string {
	return instance.name
}

func (instance mineral) GetAge() uint32 {
	return instance.age
}

func (instance *mineral) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.kind = "Kind-A"
	case "B":
		instance.kind = "B-kind"
	case "C":
		instance.kind = "ki-C-nd"
	default:
		instance.kind = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *mineral) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
