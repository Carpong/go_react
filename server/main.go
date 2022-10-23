package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/carpong/go_react/user"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user.Users{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler := user.NewUserHandler(db)

	r.POST("/adduser", handler.NewUser)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetWhere)
	r.PUT("/upusers/:id", handler.UpdateUser)
	r.DELETE("/deluser/:id", handler.DeleteUser)

	r.Run()
}
