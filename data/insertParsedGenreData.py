import pymongo
import json
import os
from dotenv import load_dotenv
from pymongo import MongoClient
from bson.objectid import ObjectId

load_dotenv()
client = MongoClient(os.getenv("MONGOURI_LOCAL"))
genresDB = client.PageTrail.genres

bookGenrePath = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_all_genres.json"

with open(bookGenrePath) as file:
    genreList = json.load(file)

for genre in genreList: 
    genresDB.insert_one({
        "name": genre
    })