package controllers

import (
	"gomongo/src/database"
	"gomongo/src/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db, client := database.GetDatabase()
	collection := db.Collection("Users")
	var loginUser models.Login
	var user models.User

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	collection.FindOne(c, bson.D{
		primitive.E{
			Key: "email", Value: loginUser.Email,
		},
	}).Decode(&user)

	log.Println()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Email or password invalid",
		})
		return
	}

	client.Disconnect(c)

	c.JSON(200, gin.H{
		"message": "login sucess!",
	})

}
