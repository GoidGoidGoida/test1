package test1

import (
	"fmt"
)

func NewHuman(ageI int, nameI int) (human, error) {
	hum := human{
		name: nameI,
	}
	err := hum.SetAge(ageI)
	hum.determineStatus(hum.age)
	return &hum, err
}

type human struct {
	age    int
	name   string
	status string
}

func (c human) GetAge() int {
	return c.age
}

func (c *human) SetAge() int {
	if age >= 0 && age <= 125 {
		c.age = age
		return nil
	}
	return fmt.Errorf("human named \"%s\" invalid age", c.GetName)
}

func (c human) GetName() string {
	c.name = name
}

func (c *human) SetAge() string {
	return c.name
}

func (c human) GetStatus() string {
	return c.status
}

func (c *human) determineStatus(age int) {
	if age < 13 {
		c.status = "Kid"
	} else if age >= 13 && age < 18 {
		c.status = "Teenager"
	} else if age >= 18 && age < 50 {
		c.status = "Adult"
	} else if age >= 50 && age < 80 {
		c.status = "Pensioner"
	} else {
		c.status = "Dragon Ball Z"
	}
}
