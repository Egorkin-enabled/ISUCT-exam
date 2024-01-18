package humans

import (
	"fmt"
)

type HumanStatus int8

const (
	child = iota
	adult
	retired
)

type human struct {
	age    int32
	name   string
	status HumanStatus
}

type IHuman interface {
	GetAge() int32
	SetAge(age int32) error

	GetName() string
	SetName(name string)

	GetStatus() HumanStatus
}

func NewHuman(age int32, name string) (IHuman, error) {
	instance := human{}

	instanceError := instance.SetAge(age)

	if instanceError != nil {
		return nil, instanceError
	}

	instance.SetName(name)

	return &instance, nil

}

func (p human) GetName() string {
	return p.name
}

func (p *human) SetName(name string) {
	p.name = name
}

func (p *human) GetStatus() HumanStatus {
	return p.status
}

func (p human) GetAge() int32 {
	return p.age
}

func (p *human) SetAge(age int32) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("The age is over bounds [0..150]: %v", age)
	}

	p.age = age

	switch {
	case age < 18:
		p.status = child
	case age < 60:
		p.status = adult
	default:
		p.status = retired
	}

	return nil
}

func CalculateAgeSum(ps []IHuman) int32 {
	var sum int32 = 0

	for _, v := range ps {
		sum += v.GetAge()
	}

	return sum
}

func TryAddUnique(peoples *([]IHuman), p IHuman) bool {
	for _, v := range *peoples {
		// Check value-by-value 'coz IHuman is reference.
		if v.GetAge() == p.GetAge() && v.GetStatus() == p.GetStatus() {
			return false
		}
	}

	*peoples = append(*peoples, p)
	return true
}
