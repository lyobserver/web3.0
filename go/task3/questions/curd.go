package questions

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

var db *gorm.DB
var err error

func InitDb() {
	// 初始化数据库连接
	dsn := "yhqalarm:1q2w!Q@W@tcp(192.168.1.106:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
}

// 插入学生记录
func CreateStudent(name string, age int, grade string) {
	db.Create(&Student{Name: name, Age: age, Grade: grade})
}

// 查询年龄大于指定值的学生记录
func GetStudentsByAge(age int) []Student {
	var students []Student
	result := db.Where("age > ?", age).Find(&students)
	if result.Error != nil {
		panic("failed to get students")
	}
	return students
}

// 更新指定姓名的学生年级信息
func UpdateStudent(name string, grade string) {
	db.Model(&Student{}).Where("name = ?", name).Update("grade", grade)
}

func DeleteStudent(age int) {
	db.Where("age < ?", age).Delete(&Student{})
}
