package controllers

import (
	"gomongo/src/database"
	"gomongo/src/helpers"
	"gomongo/src/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c *gin.Context) {
	db, client := database.GetDatabase()
	collection := db.Collection("Users")
	var user models.User
	var u models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	user.ID = primitive.NewObjectID()
	user.Email = strings.ToLower(user.Email)
	collection.FindOne(c, bson.D{
		primitive.E{Key: "email", Value: user.Email},
	}).Decode(&u)

	if emailIsValid := helpers.ValidateMail(user.Email); !emailIsValid {
		c.JSON(400, gin.H{
			"Error": "email is not valid",
		})
		return
	}

	if u.Email != "" {
		c.JSON(400, gin.H{
			"Error": "User already exists",
		})
		return
	}

	user.Password, err = helpers.GenerateHash(user.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = collection.InsertOne(c, &user)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
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
		return
	}

	if err = cursor.All(c, &users); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	client.Disconnect(c)

	c.JSON(200, users)
}

func AddOrder(c *gin.Context) {
	db, client := database.GetDatabase()
	collection := db.Collection("Users")
	var order models.Order
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	order.Data = time.Now().Format("02-01-2006 3:4:5pm")

	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	err = c.ShouldBindJSON(&order)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	update := bson.M{
		"$push": bson.M{"orders": order},
	}

	err = collection.FindOneAndUpdate(c, bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}, update).Err()

	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order added to user",
	})

	client.Disconnect(c)

}
