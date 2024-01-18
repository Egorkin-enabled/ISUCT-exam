package restaurants

// Объявление структуры 'Запись'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1800
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type restaurant struct {
	name    string
	age     uint32
	kitchen string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IRestaurant interface {
	// Геттеры
	GetKitchen() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewRestaurant(name string, age uint32) (IRestaurant, error) {
	// Создаём пустой экземпляр.
	instance := restaurant{}

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
func (instance restaurant) GetKitchen() string {
	return instance.kitchen
}

func (instance restaurant) GetName() string {
	return instance.name
}

func (instance restaurant) GetAge() uint32 {
	return instance.age
}

func (instance *restaurant) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.kitchen = "Kitchen A"
	case "B":
		instance.kitchen = "Kitchen B"
	default:
		instance.kitchen = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *restaurant) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
