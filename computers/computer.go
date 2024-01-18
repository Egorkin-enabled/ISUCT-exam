package computers

// Объявление структуры 'Компьютер'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1990
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type computer struct {
	model string
	age   uint32
	kind  string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IComputer interface {
	// Геттеры
	GetKind() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewComputer(model string, age uint32) (IComputer, error) {
	// Создаём пустой экземпляр.
	instance := computer{}

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
func (instance computer) GetKind() string {
	return instance.kind
}

func (instance computer) GetModel() string {
	return instance.model
}

func (instance computer) GetAge() uint32 {
	return instance.age
}

func (instance *computer) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.kind = "x86"
	case "B":
		instance.kind = "amd64"
	case "C":
		instance.kind = "ARM"
	default:
		instance.kind = "Chinese"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *computer) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
