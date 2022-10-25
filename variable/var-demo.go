package main

import (
	"fmt"
	"reflect"
)

var (
	name1   = "Midhun"
	course1 = "sample"
	module1 = 4
	clip1   = 2
)

func main() {
	name := "MidhunF"
	course := "sampleF"
	fmt.Println("Name:", name, "and course:", course, ".")
	module := 5
	clip := 3
	fmt.Println("Module:", module, "and clip:", clip, ".")
	fmt.Println("Name type:", reflect.TypeOf(name))
	fmt.Println("Module type:", reflect.TypeOf(module))

	fmt.Println("Mem addr of course:", &course)

	ptr := &course

	fmt.Println("pointing course at addr:", ptr, "and its value:", *ptr)

	updateCourseCopy(course)
	fmt.Println(" course:", course)
	updateCourse(&course)
	fmt.Println(" course:", course)

}

func updateCourseCopy(course string) string {
	course = "new course"
	fmt.Println("updated course:", course)
	return course
}

func updateCourse(course *string) string {
	*course = "new course"
	fmt.Println("updated course:", *course)
	return *course
}
