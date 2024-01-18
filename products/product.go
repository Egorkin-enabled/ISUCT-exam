package products

// Объявление структуры 'Запись'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Count"

const (
	minValue uint32 = 0
	maxValue uint32 = 1000
)

// Структура для записи.
// Структура закрытая.
type film struct {
	name  string
	count uint32
	price uint32
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IProduct interface {
	// Геттеры
	GetPrice() uint32
	GetName() string
	GetCount() uint32

	// Сеттеры
	SetName(string) error
	SetCount(uint32) error
}

// Реализация конструктора
func NewProduct(name string, value uint32) (IProduct, error) {
	// Создаём пустой экземпляр.
	instance := film{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(name); error != nil {
		return nil, error
	}

	if error := instance.SetCount(value); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance film) GetPrice() uint32 {
	return instance.price
}

func (instance film) GetName() string {
	return instance.name
}

func (instance film) GetCount() uint32 {
	return instance.count
}

func (instance *film) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.price = 700
	case "B":
		instance.price = 1200
	default:
		instance.price = 200
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *film) SetCount(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.count = value
	return nil
}
