package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users struct {
	Name string `json:"name" binding:"required"`
	gorm.Model
}

func (Users) Tablename() string {
	return "Users"
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (t *UserHandler) NewUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := t.db.Create(&user)
	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": user.Model.ID,
	})
}

func (t *UserHandler) GetUsers(c *gin.Context) {
	var users []Users

	result := t.db.Find(&users)
	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (t *UserHandler) GetWhere(c *gin.Context) {
	var user Users

	if err := t.db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (t *UserHandler) UpdateUser(c *gin.Context) {
	var user Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.db.Model(&user).Where("id = ?", c.Param("id")).Update("name", user.Name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success !!"})
}

func (t *UserHandler) DeleteUser(c *gin.Context) {
	var user Users

	if err := t.db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.db.Delete(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete success !!"})
}
