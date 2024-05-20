package models

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	BirthDate time.Time
}

func NewPerson(name string, day int, month time.Month, year int) *Person {
	return &Person{
		Name:      name,
		BirthDate: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

func (p *Person) Age() int {
	now := time.Now()
	years := now.Year() - p.BirthDate.Year()

	if now.Month() < p.BirthDate.Month() || (now.Month() == p.BirthDate.Month() && now.Day() < p.BirthDate.Day()) {
		years--
	}

	return years
}

func (p *Person) IsAdult() bool {
	return p.Age() >= 18
}

func (p *Person) Presentation() string {
	return fmt.Sprintf("hello my name is %s, i'm %d years old", p.Name, p.Age())
}
