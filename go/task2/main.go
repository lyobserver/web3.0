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

	//questions.OddEvenRun()
	//questions.SchedulerRun()

	rect := questions.Rectangle{Width: 10, Height: 20}
	circle := questions.Circle{Radius: 5}
	fmt.Println("矩形的面积为：", rect.Area())
	fmt.Println("矩形的周长为：", rect.Perimeter())
	fmt.Println("圆形的面积为：", circle.Area())
	fmt.Println("圆形的周长为：", circle.Perimeter())

	persion := questions.Persion{Name: "小明", Age: 22}
	employee := questions.Employee{Persion: persion, EmployeeID: 123456}
	employee.PrintInfo()

	/*
		fmt.Println("开始执行Chnnel：")
		questions.ChannelRun()
		fmt.Println("开始执行带缓存的Chnnel：")
		questions.CacheChannelRun()
	*/
	questions.MutexRun()
	questions.AtomicRun()
}
