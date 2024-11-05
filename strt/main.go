package main

import (
	"fmt"
	"strt/helper"
)

func main() {

	fmt.Println(helper.Helo)
}

type student struct {
	name  string
	grade int
}

func (s *student) input(name string, grade int) {
	s.name = name
	s.grade = grade
}
