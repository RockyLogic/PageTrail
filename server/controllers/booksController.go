package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/RockyLogic/PageTrail/models"
)

type BookController struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewBookController(database *mongo.Database) *BookController {
	return &BookController{
		database:   database,
		collection: database.Collection("books"),
	}
}

func (controller *BookController) CreateBook(c *gin.Context) {

	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result, err := controller.collection.InsertOne(context.Background(), newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create book"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (controller *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var book models.Book
	result := controller.collection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&book)
	if result == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (controller *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var updateData bson.M
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": updateData}

	var updatedBook models.Book
	err = controller.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, update, opts).Decode(&updatedBook)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update book"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}
