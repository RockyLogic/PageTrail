import pymongo
import json
import os
from dotenv import load_dotenv
from pymongo import MongoClient
from bson.objectid import ObjectId

load_dotenv()
client = MongoClient(os.getenv("MONGOURI_LOCAL"))
booksDB = client.PageTrail.books
genreDB = client.PageTrail.genres
bookGenresDB = client.PageTrail.book_genres

bookDataPath = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_data.json"

try:
    with open(bookDataPath, 'r', encoding='utf-8') as file:
        bookList = json.load(file)  # Ensure this loads correctly
except Exception as e:
    print(f"Error loading JSON data: {e}")
    bookList = []  # Initialize as an empty list in case of error

bookStruct = {
    "title": "",
    "author": "",
    "description": "",
    "pages": 0,
    "publish_date":"",
    "rating": 0,
    "edition":"",
    "language":"",
    "ISBN": 0,
}

# Map book data to schema
for book in bookList:
    newBook = bookStruct.copy()
    newBook["title"] = book["title"]
    newBook["author"] = book["authors"]
    newBook["description"] = book["description"]
    newBook["pages"] = book["num_pages"]
    newBook["publish_date"] = book["publication_date"]
    newBook["rating"] = book["rating_score"]
    newBook["edition"] = ""
    newBook["language"] = book["language"]
    newBook["ISBN"] = book["isbn"]

    newBookRes = booksDB.insert_one(newBook)
    newBookId = newBookRes.inserted_id
    
    # Handle genre mapping:
    genreListAsString = book["genres"]
    genres_list = genreListAsString.strip("[]").split(",")
    genres_list = list(map(str.strip, genres_list))  # Strip whitespace
    genres_list = list(map(lambda genre: genre.replace("'", ''), genres_list))  # Remove single quotes
    genres_list = list(map(str.lower, genres_list))  # Convert to lowercase
    
    for genre in genres_list:
        foundGenre = genreDB.find_one({"name": genre})
        
        if foundGenre is not None:
            foundGenreId = foundGenre['_id']
            
            bookGenreRelation = {
                "book_id": ObjectId(newBookId),
                "genre_id": ObjectId(foundGenreId),
            }
            bookGenresDB.insert_one(bookGenreRelation)