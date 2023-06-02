package main

import (
	"html/template"
	"os"
)

type User struct {
	Name       string
	Age        int
	Profession string
	Hobbies    []string
	Height     float32 // in meters
	Skills     map[string]bool
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name:       "John Smith",
		Age:        123,
		Profession: "Backend developer",
		Hobbies: []string{
			"Playing the guitar",
			"Playing video games",
			"Exercising",
		},
		Height: 1.80,
		Skills: map[string]bool{
			"Cooking":  false,
			"Driving":  true,
			"Outgoing": false,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
	}
}
