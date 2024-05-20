package main

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

func main() {

	people := []*Person{
		NewPerson("John", 1, time.January, 2000),
		NewPerson("Nico", 1, time.January, 2010),
		NewPerson("Jane", 1, time.January, 1990),
		NewPerson("Alice", 1, time.January, 1980),
	}

	for _, p := range people {
		fmt.Println(p.Presentation())
		if p.IsAdult() {
			fmt.Println("I'm an adult")
		} else {
			fmt.Println("I'm not an adult")
		}
	}

}
