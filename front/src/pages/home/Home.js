import { useFetch } from "../../hooks/useFetch";

import "./Home.css";

import RecipeList from "../../components/RecipeList";

export default function Home() {
  const { data, isPending, error } = useFetch("http://localhost:8080/recipes?page_id=1&page_size=5");

  return (
    <div className="home">
      {error && <p className="error">{error}</p>}
      {isPending && <p className="loading">Loading...</p>}
      {data && <RecipeList recipes={data} />}
    </div>
  );
}
