package main

import (
	"fmt"
	"github.com/superqbb/GoCourse/Object/pkg"
)

func Dinner(a pkg.Animal){
	a.Eat()
}

func WindCome(b pkg.Bird){
	b.Fly()
}

//非侵入式接口演示
func main(){
	myDog := &pkg.Dog{Weight:4}
	myPig := &pkg.Pig{Weight:80}

	//因为Dog,Pig都实现了Animal接口，编译器会自动识别
	Dinner(myDog)
	Dinner(myPig)

	fmt.Printf("风来了...\n")
	WindCome(myPig)
}