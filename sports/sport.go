package sports

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
type sport struct {
	kind    string
	age     uint32
	country string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type ISport interface {
	// Геттеры
	GetCountry() string
	GetKind() string
	GetAge() uint32

	// Сеттеры
	SetKind(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewSport(name string, age uint32) (ISport, error) {
	// Создаём пустой экземпляр.
	instance := sport{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetKind(name); error != nil {
		return nil, error
	}

	if error := instance.SetAge(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance sport) GetCountry() string {
	return instance.country
}

func (instance sport) GetKind() string {
	return instance.kind
}

func (instance sport) GetAge() uint32 {
	return instance.age
}

func (instance *sport) SetKind(name string) error {
	instance.kind = name

	switch name {
	case "A":
		instance.country = "Country A"
	case "B":
		instance.country = "Country B"
	default:
		instance.country = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *sport) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
