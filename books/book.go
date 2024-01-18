package books

// Объявление структуры 'Запись'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1800
	maxValue uint32 = 2023
)

// Структура для записи.
// Структура закрытая.
type book struct {
	name   string
	age    uint32
	author string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IBook interface {
	// Геттеры
	GetAuthor() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewBook(name string, value uint32) (IBook, error) {
	// Создаём пустой экземпляр.
	instance := book{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(name); error != nil {
		return nil, error
	}

	if error := instance.SetAge(value); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance book) GetAuthor() string {
	return instance.author
}

func (instance book) GetName() string {
	return instance.name
}

func (instance book) GetAge() uint32 {
	return instance.age
}

func (instance *book) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.author = "Author A"
	case "B":
		instance.author = "Author B"
	default:
		instance.author = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *book) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
