package main

import (
	"context"
	"log"

	"github.com/RockyLogic/PageTrail/configs"
	"github.com/RockyLogic/PageTrail/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	// MongoDB Setup
	mongoDBClient := configs.ConnectToMongoDB()
	defer func() {
		err := mongoDBClient.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()

	router := gin.Default()
	database := mongoDBClient.Database("PageTrail")
	userController := controllers.NewUserController(database)
	bookController := controllers.NewBookController(database)
	bookListController := controllers.NewBooklistController(database)

	// Users
	userRouter := router.Group("/user")
	{
		// user.POST() // TODO: Create user: Auth issue
		userRouter.GET("/:id", userController.GetUser)
		userRouter.PATCH("/:id", userController.UpdateUser)
		userRouter.DELETE("/:id", userController.DeleteUser)
	}

	bookRouter := router.Group("/book")
	{
		bookRouter.POST("", bookController.CreateBook)
		bookRouter.GET("/:id", bookController.GetBook)
		bookRouter.PATCH("/:id", bookController.UpdateBook)
	}

	bookListRouter := router.Group("/booklist")
	{

		bookListRouter.POST("", bookListController.CreateBooklist)
		bookListRouter.GET("/:id", bookListController.GetBooklist)
		bookListRouter.PATCH("/:id", bookListController.UpdateBooklist)
		bookListRouter.DELETE("/:id", bookListController.DeleteBooklist)

		bookListRouter.POST("/content/:id", bookListController.AddToBooklist)
		bookListRouter.DELETE("/content/:id", bookListController.DeleteFromBooklist)
	}

	router.Run(":8080")
}
