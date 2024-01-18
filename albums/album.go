package albums

// Объявление структуры 'Альбом'

import "fmt"

// Константы для значения

// Человеческое название
const valueHumanName = "Age"

const (
	minValue uint32 = 1900
	maxValue uint32 = 2024
)

// Структура для записи.
// Структура закрытая.
type album struct {
	name   string
	age    uint32
	artist string
}

// Чтобы обеспечить доступ к структуре из другого пакета (main),
// вводим публичный интерфейс.
type IAlbum interface {
	// Геттеры
	GetArtist() string
	GetName() string
	GetAge() uint32

	// Сеттеры
	SetName(string) error
	SetAge(uint32) error
}

// Реализация конструктора
func NewAlbum(name string, age uint32) (IAlbum, error) {
	// Создаём пустой экземпляр.
	instance := album{}

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
func (instance album) GetArtist() string {
	return instance.artist
}

func (instance album) GetName() string {
	return instance.name
}

func (instance album) GetAge() uint32 {
	return instance.age
}

func (instance *album) SetName(name string) error {
	instance.name = name

	switch name {
	case "A":
		instance.artist = "Artist A"
	case "B":
		instance.artist = "Artist B"
	default:
		instance.artist = "Unknown"
	}

	// В данной реализации ошибок быть не может,
	// Но закладываемся на будущее.
	return nil
}

func (instance *album) SetAge(value uint32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("Value '%v' out of bounds [%v, %v]: %v", valueHumanName, minValue, maxValue, value)
	}

	instance.age = value
	return nil
}
