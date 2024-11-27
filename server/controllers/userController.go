package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/RockyLogic/PageTrail/models"
)

type UserController struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserController(database *mongo.Database) *UserController {
	return &UserController{
		database:   database,
		collection: database.Collection("users"),
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {}

func (controller *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var user models.User
	result := controller.collection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user)
	if result == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var updateData bson.M
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	update := bson.M{"$set": updateData}
	var updatedUser models.User
	err = controller.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, update).Decode(&updatedUser)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var user models.User
	result := controller.collection.FindOneAndDelete(context.Background(), bson.M{"_id": objId}).Decode(&user)
	if result == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
