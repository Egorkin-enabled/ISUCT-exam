package humans

// Объявление структуры 'Человек'

import "fmt"

// Виды записей: дитя, взрослый, на пенсии
type HumanStatus uint8

const (
	Child HumanStatus = iota
	Adult
	Retired
)

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 0
	maxValue uint32 = 150
)

// Структура для записи.
// Структура закрытая.
type human struct {
	name   string
	age    uint32
	status HumanStatus
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IHuman interface {
	// Геттеры
	GetStatus() HumanStatus
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewHuman(name string, value uint32) (IHuman, error) {
	// Создаём пустой экземпляр.
	instance := human{}

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
func (instance human) GetStatus() HumanStatus {
	return instance.status
}

func (instance human) GetName() string {
	return instance.name
}

func (instance human) GetAge() uint32 {
	return instance.age
}

func (instance *human) SetName(name string) error {
	instance.name = name

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *human) SetAge(age uint32) error {
	if age < minValue || age > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, age)
	}

	instance.age = age

	switch {
	case age < 18:
		instance.status = Child
	case age < 60:
		instance.status = Adult
	default:
		instance.status = Retired
	}

	return nil
}
