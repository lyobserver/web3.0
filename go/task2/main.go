package main

import (
	"fmt"
	"main/task2/questions"
)

func main() {
	var p int = 10
	questions.ModifyValue(&p)
	fmt.Println("p的值为：", p)

	var s []int = []int{1, 2, 3, 4, 5}
	questions.ModifySlice(&s)
	fmt.Println("s的值为：", s)

	questions.OddEvenRun()
	questions.SchedulerRun()
}
