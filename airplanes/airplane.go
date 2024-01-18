package airplanes

// Объявление структуры 'Самолёт'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1900
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type airplane struct {
	model string
	age   uint32
	kind  string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IAirplane interface {
	// Геттеры
	GetKind() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewAirplane(model string, age uint32) (IAirplane, error) {
	// Создаём пустой экземпляр.
	instance := airplane{}

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
func (instance airplane) GetKind() string {
	return instance.kind
}

func (instance airplane) GetModel() string {
	return instance.model
}

func (instance airplane) GetAge() uint32 {
	return instance.age
}

func (instance *airplane) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.kind = "Kind-A"
	case "B":
		instance.kind = "B-kind"
	case "C":
		instance.kind = "Ki-C-nd"
	default:
		instance.kind = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *airplane) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
