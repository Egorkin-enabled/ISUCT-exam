package animals

// Объявление структуры 'Животное'

import "fmt"

// Виды животных: малые, средние, большие
type AnimalKind uint8

const (
	Small AnimalKind = iota
	Middle
	Large
)

// Константы для веса
const weightName = "Weight"

const (
	minWeightValue uint32 = 0
	maxWieghtValue uint32 = 1000
)

// Границы всеса, по которым опеределяются виды
const (
	smallUntil  uint32 = 1
	middleUntil        = 8
)

// Структура для животных.
// Структура закрытая.
type animal struct {
	kind   AnimalKind
	name   string
	weight uint32
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IAnimal interface {
	// Геттеры
	GetKind() AnimalKind
	GetName() string
	GetWeight() uint32

	// Сеттеры
	SetName(string) error
	SetWeight(uint32) error
}

// Реализация конструктора
func NewAnimal(name string, weight uint32) (IAnimal, error) {
	// Создаём пустой экземпляр.
	instance := animal{}

	// Устанавливаем поля с проверками на ошибки
	if error := instance.SetName(name); error != nil {
		return nil, error
	}

	if error := instance.SetWeight(weight); error != nil {
		return nil, error
	}

	// Вид установится автоматически

	return &instance, nil
}

// Реализация геттеров/сеттеров
func (instance animal) GetKind() AnimalKind {
	return instance.kind
}

func (instance animal) GetName() string {
	return instance.name
}

func (instance animal) GetWeight() uint32 {
	return instance.weight
}

func (instance *animal) SetName(name string) error {
	instance.name = name

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *animal) SetWeight(weight uint32) error {
	if weight < minWeightValue || weight > maxWieghtValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", weightName, minWeightValue, maxWieghtValue, weight)
	}

	instance.weight = weight

	switch {
	case weight <= smallUntil:
		instance.kind = Small
	case weight <= middleUntil:
		instance.kind = Middle
	default:
		instance.kind = Large
	}

	return nil
}
