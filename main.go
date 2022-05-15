package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Dummy Data User
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User //null
//var Users2 = []User{} //[]

func main() {
	r := gin.Default()

	// Grouping
	userRoutes := r.Group("/v1/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", getHandler)

	r.Run() //default port 8080
}

// todo Handler
func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  true,
		"message2": "hello world",
	})
}

// READ
func GetUsers(c *gin.Context) {
	c.JSON(200, Users)
}

// CREATE
func CreateUser(c *gin.Context) {
	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	// Menggunakan ID uuid
	reqBody.ID = uuid.New().String()
	// memasukkan value ke reqbody(User)
	Users = append(Users, reqBody)

	//res.body
	c.JSON(200, gin.H{
		"error": false,
	})
}

// UPDATE
func EditUser(c *gin.Context) {
	id := c.Param("id")

	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

// DELETE
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})

			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
