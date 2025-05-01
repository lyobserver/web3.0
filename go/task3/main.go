package main

import (
	"fmt"
	"main/task3/questions"
)

func main() {

	/*
		questions.InitSqlDB()
		SQL语句联系
			题目1：基本CRUD操作

			id, err := questions.CreateStudent("张三5", 12, "一年级")
			if err != nil {
				fmt.Println("插入学生记录失败")
			}
			fmt.Println("插入学生记录：ID=?", id)
			questions.CreateStudent("张三", 20, "三年级")
			if err != nil {
				fmt.Println("插入学生记录失败")
			}
			fmt.Println("插入学生记录：ID=?", id)
			var students []questions.Student = questions.GetStudentsByAge(18)
			if len(students) == 0 {
				fmt.Println("没有查询到学生信息")
			}
			for _, student := range students {
				fmt.Println(student.ID, student.Name, student.Age, student.Grade)
			}
			questions.UpdateStudent("张三", "四年级")
			questions.DeleteStudent(15)
			questions.CloseSqlDB()
	*/
	/*
			题目2：事务语句

		questions.TransactionRun(1, 2, 100)
		questions.CloseSqlDB()
	*/
	/*
		Sqlx入门
		题目1：使用SQL扩展库进行查询
	*/
	/*
		questions.InitSqlxDB()
		employees, err := questions.GetEmployeesByDepartment("技术部")
		if err != nil {
			fmt.Println("查询员工信息失败")
		}
		for _, employee := range employees {
			fmt.Println(employee.Id, employee.Name, employee.Department, employee.Salary)
		}
		employee, err := questions.GetEmployeeByTopSalary()
		if err != nil {
			fmt.Println("查询员工信息失败")
		}
		fmt.Println(employee.Id, employee.Name, employee.Department, employee.Salary)
		questions.CloseSqlxDB()
	*/
	/*
		题目2：实现类型安全映射
	*/
	/*
		books, err := questions.GetBooksByPrice(50)
		if err != nil {
			fmt.Println("查询书籍信息失败")
		}
		for _, book := range books {
			fmt.Println(book.Id, book.Title, book.Author, book.Price)
		}
		questions.CloseSqlxDB()
	*/
	/*
		进阶gorm
		题目1：模型定义
	*/
	questions.InitDb()

	questions.CreateTable()
	fmt.Println("表创建成功")
	questions.CloseDb()

	/*
		题目2：关联查询
	*/
	/*
		posts, err := questions.GetUserPostComments(2)
		if err != nil {
			fmt.Println("查询用户信息失败")
		}

		for _, post := range posts {
			if len(post.Comments) == 0 {
				fmt.Println(post.PostID, post.Title)
				continue
			}
			for _, comment := range post.Comments {
				fmt.Println(post.PostID, post.Title, comment.CommentID, comment.Content)
			}
		}

		post := questions.GetCommentTopOnePost()
		fmt.Println(post.PostID, post.Title)
		questions.CloseDb()
	*/
	/*
		题目3：钩子函数
	*/
	/*
		questions.CreatePost(1, "标题1", "内容1")
		questions.CreatePost(2, "标题2", "内容2")
		questions.CreatePost(3, "标题3", "内容3")
		questions.CreateComment(1, "评论1")
		questions.CreateComment(2, "评论2")
		questions.CreateComment(3, "评论3")
		questions.CreateComment(4, "评论4")
		questions.CreateComment(5, "评论5")

		questions.DeleteComment(17, 5)
	*/

	questions.CloseDb()
}
