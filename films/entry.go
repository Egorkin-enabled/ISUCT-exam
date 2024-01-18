package films

// Объявление структуры 'Запись'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1920
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type film struct {
	name     string
	age      uint32
	director string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IFilm interface {
	// Геттеры
	GetDirector() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewEntry(name string, value uint32) (IFilm, error) {
	// Создаём пустой экземпляр.
	instance := film{}

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
func (instance film) GetDirector() string {
	return instance.director
}

func (instance film) GetName() string {
	return instance.name
}

func (instance film) GetAge() uint32 {
	return instance.age
}

func (instance *film) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.director = "Director A"
	case "B":
		instance.director = "Director B"
	default:
		instance.director = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *film) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
