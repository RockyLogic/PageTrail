import csv
import json
import ast

# Function to parse the CSV file and return book data.
def parse_books_csv(file_path):
    books = []
    
    with open(file_path, mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        
        for row in reader:
            books.append(row)
    
    return books

def parse_float_or_int(value):
    try:
        if '.' in value:
            return float(value)
        return int(value)
    except ValueError:
        return None


file_path = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_data_raw.csv"
books = parse_books_csv(file_path)


json_file_path = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_data.json"  

def write_books_to_json(books, json_file_path):
    with open(json_file_path, mode='w', encoding='utf-8') as json_file:
        json.dump(books, json_file, ensure_ascii=False, indent=4)
    
write_books_to_json(books, json_file_path)

# Process Genres
genreSet = set()
json_file_path = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_data_all_genres.json"  

for book in books:
    genreListAsString = book["genres"] # [Romance ,Contemporary Romance,Contemporary,Erotica,BDSM,Adult,Suspense]
    genres_list = genreListAsString.strip("[]").split(",")
    # Strip whitespace from each element
    genres_list = [genre.strip() for genre in genres_list]
    genres_list = [genre.replace('\'', '') for genre in genres_list]
    print(genres_list)
    for genre in genres_list:
        genreSet.add(genre.lower())
    

json_file_path = r"C:\Users\Rocky.BITCOINMINER\Desktop\PageTrail\data\book_all_genres.json"  
unique_genres = list(genreSet)
with open(json_file_path, mode='w', encoding='utf-8') as json_file:
    json.dump(unique_genres, json_file, ensure_ascii=False, indent=4)