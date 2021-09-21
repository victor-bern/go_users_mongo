package controllers

import (
	"gomongo/src/database"
	"gomongo/src/helpers"
	"gomongo/src/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	user.Password, err = helpers.GenerateHash(user.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

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
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	client.Disconnect(c)

	c.JSON(200, users)
}
