package games

// Объявление структуры 'Игра'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1980
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type game struct {
	name   string
	age    uint32
	genere string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IGame interface {
	// Геттеры
	GetGenere() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewGame(name string, age uint32) (IGame, error) {
	// Создаём пустой экземпляр.
	instance := game{}

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
func (instance game) GetGenere() string {
	return instance.genere
}

func (instance game) GetName() string {
	return instance.name
}

func (instance game) GetAge() uint32 {
	return instance.age
}

func (instance *game) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.genere = "Genere A"
	case "B":
		instance.genere = "Ganere B"
	default:
		instance.genere = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *game) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
