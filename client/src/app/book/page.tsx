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

export default async function GeneralBookPage() {
  return (
    <>
      <h1>General Book Page</h1>
    </>
  );
}
