package main

import (
	"fmt"
	"github.com/nicoabatedaga/golang_workshop_using_library/models"
	"time"
)

func main() {

	people := []*models.Person{
		models.NewPerson("John", 1, time.January, 2000),
		models.NewPerson("Nico", 1, time.January, 2010),
		models.NewPerson("Jane", 1, time.January, 1990),
		models.NewPerson("Alice", 1, time.January, 1980),
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
