import axios from "axios";

interface Book {
  id: string;
  title: string;
  description: string;
  pages: number;
  publish_date: string;
  rating: number;
  edition: string;
  language: string;
  ISBN: string;
}

type BookPageProps = {
  params: {
    id: string;
  };
};

export default async function BookPage({ params }: BookPageProps) {
  const book = await axios.get<Book>(`http://localhost:8080/book/${params.id}`);

  console.log(book);

  return (
    <>
      <h1>Book Page</h1>
      <div>{book.data.id}</div>
      <div>{book.data.title}</div>
      <div>{book.data.pages}</div>
      <div>{book.data.publish_date}</div>
      <div>{book.data.rating}</div>
      <div>{book.data.edition}</div>
      <div>{book.data.language}</div>
      <div>{book.data.ISBN}</div>
      <div>{book.data.description}</div>
    </>
  );
}
