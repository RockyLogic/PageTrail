package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	JoinDate    string             `bson:"join_date" json:"join_date"`
	Language    string             `bson:"language" json:"language"`
	GPTSession  string             `bson:"gpt_session" json:"gpt_session"`
	DateOfBirth string             `bson:"dob" json:"dob"`
}

type Book struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectID for primary key
	Title        string             `bson:"title" json:"title"`
	Author       string             `bson:"author" json:"author"`
	Description  string             `bson:"description" json:"description"`
	Genre        string             `bson:"genre" json:"genre"`
	Pages        string             `bson:"pages" json:"pages"`
	Rating       string             `bson:"rating" json:"rating"`
	PublishMonth string             `bson:"publish_month" json:"publish_month"`
	PublishYear  string             `bson:"publish_year" json:"publish_year"`
	Edition      string             `bson:"edition" json:"edition"`
	Language     string             `bson:"language" json:"language"`
	ISBN         string             `bson:"isbn" json:"isbn"`
}

type Booklist struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	ShareURL    string             `bson:"share_url" json:"share_url"`
}

type BooklistContent struct {
	BooklistID primitive.ObjectID `bson:"booklist_id" json:"booklist_id"`
	BookID     primitive.ObjectID `bson:"book_id" json:"book_id"`
}
