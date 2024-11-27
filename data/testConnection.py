import os
from dotenv import load_dotenv
from pymongo import MongoClient
from bson.objectid import ObjectId


load_dotenv()
client = MongoClient(os.getenv("MONGOURI_LOCAL"))
db = client.PageTrail
usersDB = db.users

user = usersDB.find_one({"_id": ObjectId("671fe512767ee83551ebaec2")})
print(user)
