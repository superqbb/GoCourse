package main

import "github.com/superqbb/GoCourse/Object/demo01"

func main() {
	mark := demo01.Student{demo01.Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := demo01.Employee{demo01.Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}