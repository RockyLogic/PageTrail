import os
from dotenv import load_dotenv
from pymongo import MongoClient
from bson.objectid import ObjectId

load_dotenv()
client = MongoClient(os.getenv("MONGOURI_LOCAL"))

db = client.PageTrail

usersDB = db.users
userActivityDB = db.user_activity
favouritesDB = db.favourites
ratingsDB = db.ratings
bookListsDB = db.booklists
bookListContentDB = db.booklist_contents

booksDB = db.books
genreDB = db.genres
bookGenreDB = db.book_genre

booksReadDB = db.books_read
booksToReadDB = db.books_to_read
booksInProgressDB = db.books_in_progress

followingDB = db.following

usersDB.insert_one(
    {
        "username":"Rocky",
        "email":"Rocky@gmail.com",
        "first_name": "Rocky",
        "last_name" : "Kuang",
        "join_date" : "10/28/2024",
        "language" : "english",
        "gpt-session" : "",
        "dob" : "09-19-2000",
    }
)

books = [{
    "title": "Summer Story",
    "author": "Jill Barklem",
    "description": "It was such a hot summer. The sky was deep blue and the sun never faltered.All along Brambly Hedge, the mice did their best to keep cool. Poppy Eyebright sought refuge in the mossy shadows of the mill wheel; Dusty Dogwood took to walking by the banks of the cooling stream. Dusty and Poppy spent more and more time together, so no one was at all surprised when they announced their engagement. They decided on a very unusual setting for the wedding ceremony, but even they didn't realize just how unusual it would prove to be!",
    "pages": 32,
    "publish_date":"January 1, 1980",
    "rating": 4.45,
    "edition":"",
    "language":"English",
    "ISBN": 9780689830594,
},
{
    "title": "The Lake of Darkness",
    "author": "Ruth Rendell",
    "description": "Martin Urban is a quiet bachelor with a comfortable life, free of worry and distractions. When he unexpectedly comes into a small fortune, he decides to use his newfound wealth to help out those in need. Finn also leads a quiet life, and comes into a little money of his own. Normally, their paths would never have crossed. But Martin’s ideas about who should benefit from his charitable impulses yield some unexpected results, and soon the good intentions of the one become fatally entangled with the mercenary nature of the other. In the Lake of Darkness , Ruth Rendell takes the old adage that no good deed goes unpunished to a startling, haunting conclusion.",
    "pages": 210,
    "publish_date":"January 1, 1980",
    "rating": 3.76,
    "edition":"",
    "language":"English",
    "ISBN": 9780375704970,
}, 
{
    "title": "Beyond the Blue Event Horizon",
    "author": "Frederik Pohl",
    "description": "In Book Two of the Heechee Saga, Robinette Broadhead is on his way to making a fortune by bankrolling an expedition to the Food Factory--a Heechee spaceship that can graze the cometary cloud and transfor the basic elements of the universe into untold quantities of food. But even as he gambles on the breakthrough technology, he is wracked with the guilt of losing his wife, poised forever at the ""event horizon"" of a black hole where Robin had abaondoned her. As more and more information comes back from the expedition, Robin grows ever hopeful that he can rescue his beloved Gelle-Klara Moynlin. After three and a years, the factory is discovered to work, and a human is found aboard. Robin's suffering may be just about over....",
    "pages": 336,
    "publish_date":"January 1, 1980",
    "rating": 3.95,
    "edition":"",
    "language":"English",
    "ISBN": 9780345446671,
},
{
    "title": "St. Peter's Fair",
    "author": "Ellis Peters",
    "description": "A pause in the civil war offers Shrewsbury's townsfolk hope that the upcoming fair will be successful, but the discovery of the body of a wealthy merchant could destroy that hope.",
    "pages": 217,
    "publish_date":"May 1, 1981",
    "rating": 4.12,
    "edition":"",
    "language":"English",
    "ISBN": 9780446403016,
},
{
    "title": "Twice Shy",
    "author": "Dick Francis",
    "description": "A computerized horse-betting system falls into Jonathan Derry's hands--and unless he returns it to the rightful owners, the odds of his survival are slim to none.",
    "pages": 304,
    "publish_date":"January 1, 1981",
    "rating": 3.92,
    "edition":"",
    "language":"English",
    "ISBN": 9780425198773,
}]

booksDB.insert_many(books)

ratingsDB.insert_one(
    {
        "user_id": ObjectId("671fe512767ee83551ebaec2"),
        "book_id": ObjectId("671fe9022dfa7143767a4a6e"),
        "rating": 8.5,
        "comment": "Good Read."
    }
)

favouritesDB.insert_one({
    "user_id": ObjectId("671fe512767ee83551ebaec2"),
    "book_id": ObjectId("671fec2a92eeecec8f1c047f")
})

genres = [
  { "genreName": "Picture Books" },
  { "genreName": "Childrens" },
  { "genreName": "Fiction" },
  { "genreName": "Animals" },
  { "genreName": "Fantasy" },
  { "genreName": "Classics" },
  { "genreName": "Short Stories" },
  { "genreName": "Mystery" },
  { "genreName": "Crime" },
  { "genreName": "Thriller" },
  { "genreName": "British Literature" },
  { "genreName": "Suspense" },
  { "genreName": "Mystery Thriller" },
  { "genreName": "Science Fiction" },
  { "genreName": "Space Opera" },
  { "genreName": "Science Fiction Fantasy" },
  { "genreName": "Aliens" },
  { "genreName": "Adventure" },
  { "genreName": "Historical Fiction" },
  { "genreName": "Historical" },
  { "genreName": "Medieval" },
  { "genreName": "Historical Mystery" },
  { "genreName": "Horse Racing" },
  { "genreName": "Young Adult" },
  { "genreName": "Fairy Tales" },
  { "genreName": "Retellings" },
  { "genreName": "Romance" }
]
genreDB.insert_many(genres)

