package user_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from user ",
	})
}

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from getUsers ",
	})
}

func GetUsersById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from getUsersById ",
	})
}

func InsertUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from insertUser ",
	})
}

func UpdateUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from updateUserById ",
	})
}
