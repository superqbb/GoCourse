package pkg

import "fmt"

type Animal interface{
	Eat()
}

type Bird interface {
	Fly()
}

//我们并没有在代码的任何地方告诉Dog或者Pig这两个struct它们需要去实现Animal接口
type Dog struct {
	Weight float64
}

func(d *Dog) Eat(){
	fmt.Printf("体重%.1fkg的狗正在进食...\n",d.Weight)
}

type Pig struct {
	Weight float64
}

func(d *Pig) Fly(){
	fmt.Printf("体重%.1fkg的猪正在飞...\n",d.Weight)
}

func(d *Pig) Eat(){
	fmt.Printf("体重%.1fkg的猪正在进食...\n",d.Weight)
}