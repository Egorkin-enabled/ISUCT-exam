package cities

// Объявление структуры 'Город'

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
type city struct {
	name    string
	age     uint32
	country string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type ICity interface {
	// Геттеры
	GetCountry() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewCity(name string, age uint32) (ICity, error) {
	// Создаём пустой экземпляр.
	instance := city{}

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
func (instance city) GetCountry() string {
	return instance.country
}

func (instance city) GetName() string {
	return instance.name
}

func (instance city) GetAge() uint32 {
	return instance.age
}

func (instance *city) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.country = "Russia"
	case "B":
		instance.country = "China"
	case "C":
		instance.country = "UK"
	default:
		instance.country = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *city) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
