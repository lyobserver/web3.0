package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sqlxdb *sqlx.DB
var sqlxerr error

func InitSqlxDB() {
	dsn := "yhqalarm:1q2w!Q@W@tcp(192.168.1.106:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	sqlxdb, sqlxerr = sqlx.Open("mysql", dsn)
	if sqlxerr != nil {
		panic("failed to connect database")
	}
}

func CloseSqlxDB() {
	sqlxdb.Close()
}

type employees struct {
	Id         uint64  `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type books struct {
	Id     uint64  `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 根据部门读取员工信息
func GetEmployeesByDepartment(department string) ([]employees, error) {
	var employees []employees
	err := sqlxdb.Select(&employees, "SELECT id,name,department,salary FROM employees WHERE department = ?", department)
	return employees, err
}

// 根据工资最高的员工信息
func GetEmployeeByTopSalary() (employees, error) {
	var employee employees
	err := sqlxdb.Get(&employee, "SELECT id,name,department,salary FROM employees ORDER BY salary DESC LIMIT 1")
	return employee, err
}

// 读取价格大于指定值的书籍信息
func GetBooksByPrice(price float64) ([]books, error) {
	var books []books
	err := sqlxdb.Select(&books, "SELECT id,title,author,price FROM books WHERE price >?", price)
	return books, err
}
