package main

import (
	"main/task3/questions"
)

func main() {
	questions.InitDb()
	/*
		题目1：基本CRUD操作

		questions.CreateStudent("张三5", 12, "一年级")
		questions.CreateStudent("张三", 20, "三年级")
		var students []questions.Student = questions.GetStudentsByAge(18)
		if len(students) == 0 {
			fmt.Println("没有查询到学生信息")
		}
		for _, student := range students {
			fmt.Println(student.ID, student.Name, student.Age, student.Grade)
		}
		questions.UpdateStudent("张三", "四年级")
		questions.DeleteStudent(15)
	*/
	/*
		题目2：事务语句
	*/
	questions.TransactionRun(1, 2, 100)
}
