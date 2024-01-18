package cars

// Объявление структуры 'Запись'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1900
	maxValue uint32 = 2023
)

// Структура для записи.
// Структура закрытая.
type car struct {
	name  string
	age   uint32
	speed uint32
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type ICar interface {
	// Геттеры
	GetSpeed() uint32
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewEntry(name string, age uint32) (ICar, error) {
	// Создаём пустой экземпляр.
	instance := car{}

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
func (instance car) GetSpeed() uint32 {
	return instance.speed
}

func (instance car) GetName() string {
	return instance.name
}

func (instance car) GetAge() uint32 {
	return instance.age
}

func (instance *car) updateSpeed() {
	switch {
	case instance.age < 1885:
		instance.speed = 0

	case instance.age < 2000:
		switch instance.name {
		case "A":
			instance.speed = 120
		case "B":
			instance.speed = 100
		default:
			instance.speed = 80
		}

	default:
		switch instance.name {
		case "A":
			instance.speed = 200
		case "B":
			instance.speed = 150
		default:
			instance.speed = 120
		}
	}
}

func (instance *car) SetName(name string) error {
	instance.name = name

	instance.updateSpeed()

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *car) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	instance.updateSpeed()

	return nil
}
