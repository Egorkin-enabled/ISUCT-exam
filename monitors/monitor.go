package monitors

// Объявление структуры 'Монитор'

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
type monitor struct {
	model      string
	age        uint32
	screenSize string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IMonitor interface {
	// Геттеры
	GetScreenSize() string
	GetModel() string
	GetAge() uint32

	// Сеттеры
	SetModel(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewComputer(model string, age uint32) (IMonitor, error) {
	// Создаём пустой экземпляр.
	instance := monitor{}

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
func (instance monitor) GetScreenSize() string {
	return instance.screenSize
}

func (instance monitor) GetModel() string {
	return instance.model
}

func (instance monitor) GetAge() uint32 {
	return instance.age
}

func (instance *monitor) SetModel(name string) error {
	instance.model = name

	switch name {
	case "A":
		instance.screenSize = "1920x1080"
	case "B":
		instance.screenSize = "640x480"
	case "C":
		instance.screenSize = "256x176"
	default:
		instance.screenSize = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *monitor) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
