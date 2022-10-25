package main

import (
	"fmt"
	"golang.org/x/text/cases"
	_ "golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func main() {
	author := "test author"
	course := "some course"
	fmt.Println(converter(author, course))
}

func converter(author string, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = cases.Title(language.English).String(course)
	return author, course
}
