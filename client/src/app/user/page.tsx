import axios from "axios";

interface User {
  id: string;
  username: string;
  email: string;
  first_name: string;
  last_name: string;
  join_date: string;
  language: string;
  gpt_session: string;
  dob: string;
}

const HomePage = async () => {
  return (
    <>
      <h1>General User Page</h1>
    </>
  );
};

export default HomePage;
