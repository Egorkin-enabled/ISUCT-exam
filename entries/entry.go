package entries

// Объявление структуры 'Запись'

import "fmt"

// Виды записей: вид 1, вид 2, вид 3
type EntryKind uint8

const (
	Kind1 EntryKind = iota
	Kind2
	Kind3
)

// Константы для значения

// Человеческое название
const valueHumanName = "Value"

const (
	minValue uint32 = 100
	maxValue uint32 = 200
)

// Границы всеса, по которым опеределяются виды
const (
	kind1UntilValue uint32 = 133
	kind2UntilValue        = 166
)

// Структура для животных.
// Структура закрытая.
type entry struct {
	name  string
	value uint32
	kind  EntryKind
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IEntry interface {
	// Геттеры
	GetKind() EntryKind
	GetName() string
	GetValue() uint32

	// Сеттеры
	SetName(string) error
	SetValue(uint32) error
}

// Реализация конструктора
func NewEntry(name string, value uint32) (IEntry, error) {
	// Создаём пустой экземпляр.
	instance := entry{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(name); error != nil {
		return nil, error
	}

	if error := instance.SetValue(value); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance entry) GetKind() EntryKind {
	return instance.kind
}

func (instance entry) GetName() string {
	return instance.name
}

func (instance entry) GetValue() uint32 {
	return instance.value
}

func (instance *entry) SetName(name string) error {
	instance.name = name

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *entry) SetValue(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.value = value

	switch {
	case value <= kind1UntilValue:
		instance.kind = Kind1
	case value <= kind2UntilValue:
		instance.kind = Kind2
	default:
		instance.kind = Kind3
	}

	return nil
}
