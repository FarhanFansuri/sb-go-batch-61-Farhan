package main

import "fmt"

func main() {
	var s1 student
	s1.input("farhan", 100)

	fmt.Println(s1.name)

	murid2 := struct {
		name  string
		grade int
	}{}

	murid2.name = "John"
	fmt.Println(murid2.name)
}

type student struct {
	name  string
	grade int
}

func (s *student) input(name string, grade int) {
	s.name = name
	s.grade = grade
}
