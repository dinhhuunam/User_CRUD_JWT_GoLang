package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey,column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	//CreatedAt *time.Time `json:"created_at"`
	//UpdatedAt *time.Time `json:"updated_at"`
}

func main() {
	dsn := "springstudent:springstudent@tcp(127.0.0.1:3306)/user_manage?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to database", db)

	// Migrate the schema
	db.AutoMigrate(&User{})

	//now := time.Now().UTC()
	//user := User{
	//	Id:       1,
	//	Username: "john",
	//	Email:    "john@gmail.com",
	//	Password: "123456",
	//	CreatedAt: &now,
	//	UpdatedAt: &now,
	//}
	db.Create(&User{Username: "john", Email: "john@gmail.com", Password: "123456"})

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": user,
	//	})
	//})
	//r.Run()
}
