package phones

// Объявление структуры 'Телефон'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 2000
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type phone struct {
	model        string
	age          uint32
	manufacturer string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IPhone interface {
	// Геттеры
	GetManufacturer() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewPhone(model string, age uint32) (IPhone, error) {
	// Создаём пустой экземпляр.
	instance := phone{}

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
func (instance phone) GetManufacturer() string {
	return instance.manufacturer
}

func (instance phone) GetModel() string {
	return instance.model
}

func (instance phone) GetAge() uint32 {
	return instance.age
}

func (instance *phone) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.manufacturer = "Apple"
	case "B":
		instance.manufacturer = "Samsung"
	case "C":
		instance.manufacturer = "Infinix"
	default:
		instance.manufacturer = "NoName"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *phone) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
