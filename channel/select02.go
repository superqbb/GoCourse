package main

import "fmt"

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main () {
	/*
	此时，select语句走的是default操作。但是这时每个case的表达式都会被执行
	1、系统会从左到右执行表达式，先执行getChan函数打印chs[0]，然后执行getNumber函数打印
	2、从上到下分别执行所有case语句中的表达式
	*/
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}