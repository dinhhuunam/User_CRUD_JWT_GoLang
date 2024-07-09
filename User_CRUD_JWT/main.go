package main

import (
	"User_CRUD_JWT/modules/item/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

//	type User struct {
//		gorm.Model
//		Id        int       `json:"id" gorm:"primaryKey,column:id"`
//		Username  string    `json:"username" gorm:"not null,unique,column:username"`
//		Email     string    `json:"email" gorm:"not null,unique,column:email"`
//		Password  string    `json:"password" gorm:"not null,column:password"`
//		CreatedAt time.Time `json:"created_at"`
//		UpdatedAt time.Time `json:"updated_at"`
//	}
//
//	type UserCreation struct {
//		Username string `json:"username" gorm:"not null,unique,column:username"`
//		Email    string `json:"email" gorm:"not null,unique,column:email"`
//		Password string `json:"password" gorm:"not null,column:password"`
//	}
//
//	type UserRead struct {
//		Id       int    `json:"id" gorm:"primaryKey,column:id"`
//		Username string `json:"username" gorm:"not null,unique,column:username"`
//	}
//
//	func (UserRead) TableName() string {
//		return "users"
//	}
//
//	func (UserCreation) TableName() string {
//		return "users"
//	}
func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to database", db)

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	//now := time.Now().UTC()
	//user := User{
	//	Id:       1,
	//	Username: "john",
	//	Email:    "john@gmail.com",
	//	Password: "123456",
	//	CreatedAt: &now,
	//	UpdatedAt: &now,
	//}
	//db.Create(&User{Username: "john", Email: "john@gmail.com", Password: "123456"})

	//CRUD: Create, Read,Update,Delete
	//POST: /v1/users (Create new user)
	//GET: /v1/users (list items)
	//GET: /v1/users/:id (get items detail by id)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", CreateUser(db))
			users.GET("", ListUser(db))
			users.GET("/:id", readUserById(db))
			users.PATCH("/:id", editUserById(db))
			users.DELETE("/:id", deleteUserById(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Create user

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UserCreation

		//Get data from URL
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

//Read item by id

func readUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.UserRead

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

//Edit item by Id

func editUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var data model.UserCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

//List user

func ListUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.UserRead

		if err := db.Find(&result).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func deleteUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(model.UserCreation{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
