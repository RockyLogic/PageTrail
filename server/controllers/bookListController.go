package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/RockyLogic/PageTrail/models"
)

type BooklistController struct {
	database            *mongo.Database
	BLCollection        *mongo.Collection
	BLContentCollection *mongo.Collection
}

func NewBooklistController(database *mongo.Database) *BooklistController {
	return &BooklistController{
		database:            database,
		BLCollection:        database.Collection("booklists"),
		BLContentCollection: database.Collection("booklist_contents"),
	}
}

// Create Booklist
func (controller *BooklistController) CreateBooklist(c *gin.Context) {
	var newBooklist models.Booklist
	if err := c.BindJSON(&newBooklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result, err := controller.BLCollection.InsertOne(context.Background(), newBooklist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create booklist"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get booklist and contents
func (controller *BooklistController) GetBooklist(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booklist ID format"})
		return
	}

	println(id)
	var booklist models.Booklist
	result := controller.BLCollection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&booklist)
	if result == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booklist not found"})
		return
	}

	// Books part of booklist
	var booklistContents []models.BooklistContent
	cursor, err := controller.BLContentCollection.Find(context.Background(), bson.M{"booklist_id": objId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error finding contents of booklist"})
		return
	}

	for cursor.Next(context.Background()) {
		var content models.BooklistContent
		if err := cursor.Decode(&content); err != nil {
			// Handle error (e.g., log it or return a response to the client)
			log.Println("Error decoding booklist content:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode booklist content"})
			return
		}
		booklistContents = append(booklistContents, content)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading from cursor"})
		return
	}

	responseData := map[string]interface{}{
		"booklist":         booklist,
		"booklistContents": booklistContents,
	}

	c.JSON(http.StatusOK, responseData)
}

// Update general info on booklist booklists table
func (controller *BooklistController) UpdateBooklist(c *gin.Context) {
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

	var updatedBookList models.Booklist
	err = controller.BLCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, update, opts).Decode(&updatedBookList)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update book"})
		return
	}
	c.JSON(http.StatusOK, updatedBookList)
}

// Remove booklist in booklists table and its contents in M-N table
func (controller *BooklistController) DeleteBooklist(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	results, err := controller.BLContentCollection.DeleteMany(context.Background(), bson.M{"booklist_id": objId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete booklist contents"})
		return
	}

	log.Printf("Deleted %d booklist contents.", results.DeletedCount)

	var booklist models.Booklist
	res := controller.BLCollection.FindOneAndDelete(context.Background(), bson.M{"_id": objId}).Decode(&booklist)
	if res == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booklist not found"})
		return
	}
	c.JSON(http.StatusOK, booklist)
}

// Create new document in M-N booklist_contents table
func (controller *BooklistController) AddToBooklist(c *gin.Context) {
	booklistId := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(booklistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var newBooklistContent models.BooklistContent
	if err := c.BindJSON(&newBooklistContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	newBooklistContent.BooklistID = objId

	// BookID is nonempty
	if newBooklistContent.BookID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BookID cannot be empty"})
		return
	}

	result, err := controller.BLContentCollection.InsertOne(context.Background(), newBooklistContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create booklist"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Delete document in M-N booklist_contents table
func (controller *BooklistController) DeleteFromBooklist(c *gin.Context) {
	booklistId := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(booklistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var newBooklistContent models.BooklistContent
	if err := c.BindJSON(&newBooklistContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	newBooklistContent.BooklistID = objId

	// BookID is nonempty
	if newBooklistContent.BookID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BookID cannot be empty"})
		return
	}

	var booklistContent models.BooklistContent
	result := controller.BLContentCollection.FindOneAndDelete(context.Background(), bson.M{"booklist_id": objId, "book_id": newBooklistContent.BookID}).Decode(&booklistContent)
	if result == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found in booklist"})
		return
	}
	c.JSON(http.StatusOK, booklistContent)
}
