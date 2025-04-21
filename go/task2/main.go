package main

import (
	"fmt"
	"main/task2/questions"
)

func main() {
	var p int = 10
	questions.ModifyValue(&p)
	fmt.Println("p的值为：", p)

}
