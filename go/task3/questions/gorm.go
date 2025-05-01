package questions

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserID    uint64 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	PostCount int    `gorm:"type:int;default:0"`
	Posts     []Post `gorm:"foreignKey:user_id"`
}
type Post struct {
	PostID       uint64    `gorm:"primaryKey"`
	UserID       uint64    `gorm:"index"`
	Title        string    `gorm:"type:varchar(255)"`
	Content      string    `gorm:"type:text"`
	CommentState uint8     `gorm:"type:tinyint;default:0;comment:'0: 未评论, 1: 已评论'"`
	Comments     []Comment `gorm:"foreignKey:post_id"`
}

type Comment struct {
	CommentID uint64 `gorm:"primaryKey"`
	Content   string `gorm:"type:text"`
	PostID    uint64 `gorm:"index"`
}

var db *gorm.DB
var err error

func InitDb() {
	// 初始化数据库连接
	dsn := "yhqalarm:1q2w!Q@W@tcp(192.168.1.106:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
func CloseDb() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}

// 根据结构体创建表
func CreateTable() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// 读取指定用户的全部文章和评论
func GetUserPostComments(UserID uint64) ([]*Post, error) {
	var posts []*Post
	err := db.Debug().Preload("Comments").Joins("left join comments on comments.post_id=posts.post_id").Find(&posts, "posts.user_id = ?", UserID).Error
	return posts, err
}

// 读取评论数量最多的文章
func GetCommentTopOnePost() *Post {
	var post Post
	db.Debug().Select("posts.post_id,posts.title,count(*) commentCount").Joins("join comments on comments.post_id=posts.post_id").Group("posts.post_id").Order("commentCount desc").Take(&post)
	return &post
}

func CreatePost(userID uint64, title string, content string) error {
	post := Post{UserID: userID, Title: title, Content: content}
	err := db.Create(&post).Error
	return err
}

func DeletePost(postID uint64, userID uint64) error {
	post := Post{PostID: postID, UserID: userID}
	err := db.Delete(&post).Error
	return err
}

func CreateComment(postID uint64, content string) error {
	comment := Comment{PostID: postID, Content: content}
	err := db.Create(&comment).Error
	return err
}

func DeleteComment(commentID uint64, postID uint64) error {
	comment := Comment{CommentID: commentID, PostID: postID}
	err := db.Debug().Delete(&comment).Error
	return err
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	var count int64
	if err := tx.Model(&Post{}).Where("user_id=?", p.UserID).Count(&count).Error; err != nil {
		return err
	}
	if err := tx.Model(&User{}).Where("user_id = ?", p.UserID).Update("post_count", count).Error; err != nil {
		return err
	}
	return nil
}

func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	var count int64
	if err := tx.Model(&Post{}).Where("user_id=?", p.UserID).Count(&count).Error; err != nil {
		return err
	}
	if err := tx.Model(&User{}).Where("user_id =?", p.UserID).Update("post_count", count).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论状态
	fmt.Println(c.PostID)
	if err := tx.Model(&Post{}).Where("post_id =?", c.PostID).Update("comment_state", 1).Error; err != nil {
		return err
	}
	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 检查是否还有评论
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id =?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	// 更新文章的评论状态
	if count == 0 {
		if err := tx.Model(&Post{}).Where("post_id =?", c.PostID).Update("comment_state", 0).Error; err != nil {
			return err
		}
	}
	return nil
}
