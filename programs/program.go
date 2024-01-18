package programs

// Объявление структуры 'Самолёт'

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
type program struct {
	name string
	age  uint32
	kind string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IProgram interface {
	// Геттеры
	GetKind() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewProgram(model string, age uint32) (IProgram, error) {
	// Создаём пустой экземпляр.
	instance := program{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(model); error != nil {
		return nil, error
	}

	if error := instance.SetAge(age); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance program) GetKind() string {
	return instance.kind
}

func (instance program) GetName() string {
	return instance.name
}

func (instance program) GetAge() uint32 {
	return instance.age
}

func (instance *program) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.kind = "Illustrate"
	case "B":
		instance.kind = "CAD"
	case "C":
		instance.kind = "Game"
	default:
		instance.kind = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *program) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
