package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("Before running tests")
	_ = os.Setenv("workshop_golang", "Hello World")
	m.Run()
	println("After running tests")
	_ = os.Setenv("workshop_golang", "")
}

func TestNewPerson(t *testing.T) {
	println("Running TestNewPerson")
	env := os.Getenv("workshop_golang")
	println(env)
	println("End TestNewPerson")
}
