package controllers

import (
	"gomongo/src/database"
	"gomongo/src/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *gin.Context) {
	db, client := database.GetDatabase()
	collection := db.Collection("Users")
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
	}

	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MaxCost)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	// user.Password = ph

	_, err = collection.InsertOne(c, &user)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
	}

	client.Disconnect(c)

	c.JSON(200, gin.H{
		"message": "sucess",
	})
}

func GetAll(c *gin.Context) {
	db, client := database.GetDatabase()
	collection := db.Collection("Users")
	var users []models.User

	cursor, err := collection.Find(c, bson.M{})

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	if err = cursor.All(c, &users); err != nil {
		log.Fatal(err)
	}
	client.Disconnect(c)

	c.JSON(200, users)
}
