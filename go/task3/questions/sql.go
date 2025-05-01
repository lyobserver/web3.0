package questions

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var sqldb *sql.DB
var sqlerr error

type Student struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

func InitSqlDB() {
	dsn := "yhqalarm:1q2w!Q@W@tcp(192.168.1.106:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	sqldb, sqlerr = sql.Open("mysql", dsn)
	if sqlerr != nil {
		panic(err)
	}
}

func CloseSqlDB() {
	if sqldb == nil {
		return
	}
	sqldb.Close()
}

// 插入学生记录
func CreateStudent(name string, age int, grade string) (int64, error) {
	result, err := sqldb.Exec("insert into Students(name,age,grade) values(?,?,?)", name, age, grade)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 查询年龄大于指定值的学生记录
func GetStudentsByAge(age int) []Student {
	var students []Student
	rows, err := sqldb.Query("select * from students where age > ?", age)
	if err != nil {
		panic("failed to get students")
	}
	defer rows.Close()
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
		if err != nil {
			panic("failed to scan student")
		}
		students = append(students, student)
	}
	return students
}

// 更新指定姓名的学生年级信息
func UpdateStudent(name string, grade string) int64 {
	result, err := sqldb.Exec("update students set grade =? where name =?", grade, name)
	if err != nil {
		return 0
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	return count
}

func DeleteStudent(age int) int64 {
	result, err := sqldb.Exec("delete from students where age <?", age)
	if err != nil {
		return 0
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	return count
}

type Account struct {
	ID      uint64
	Balance int
}

type Transaction struct {
	ID              uint64
	From_account_id uint64
	To_account_id   uint64
	Amount          int
}

func TransactionRun(aid uint64, bid uint64, amount int) {
	tx, err := sqldb.Begin()
	if err != nil {
		fmt.Println("开启事务失败")
		return
	}
	var accountA Account
	var accountB Account
	if err := tx.QueryRow("select id,balance from accounts where id=?", aid).Scan(&accountA.ID, &accountA.Balance); err != nil {
		fmt.Println("查询账户A失败")
		tx.Rollback()
		return
	}
	if accountA.Balance < amount {
		fmt.Println("余额不足")
		tx.Rollback()
		return
	}
	if _, err := tx.Exec("update accounts set balance=? where id = ?", accountA.Balance-amount, aid); err != nil {
		fmt.Println("账户A扣款失败")
		tx.Rollback()
		return
	}
	if err := tx.QueryRow("select id,balance from accounts where id=?", bid).Scan(&accountB.ID, &accountB.Balance); err != nil {
		fmt.Println("查询账户B失败")
		tx.Rollback()
		return
	}
	if _, err := tx.Exec("update accounts set balance=? where id = ?", accountB.Balance+amount, bid); err != nil {
		fmt.Println("账户B存款失败")
		tx.Rollback()
		return
	}
	if _, err := tx.Exec("insert into transactions (from_account_id,to_account_id,amount) values(?,?,?)", aid, bid, amount); err != nil {
		fmt.Println("创建交易记录失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	fmt.Println("转账成功")
}
