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

type UserPageProps = {
  params: {
    id: string;
  };
};

export default async function UserPage({ params }: UserPageProps) {
  const user = await axios.get<User>(`http://localhost:8080/user/${params.id}`);

  return (
    <>
      <h1>User Page</h1>
      <div>{user.data.id}</div>
      <div>{user.data.username}</div>
      <div>{user.data.email}</div>
      <div>{user.data.first_name}</div>
      <div>{user.data.last_name}</div>
      <div>{user.data.join_date}</div>
      <div>{user.data.language}</div>
      <div>{user.data.gpt_session}</div>
      <div>{user.data.dob}</div>
    </>
  );
}
