import { useParams } from "react-router-dom";

import { useFetch } from "../../hooks/useFetch";

import "./Recipe.css";

export default function Recipe() {
  const { id } = useParams();
  const url = "http://localhost:8080/recipes/" + id;
  const { data: recipe, error, isPending } = useFetch(url);
  return (
    <div className="recipe">
      {error && <p className="error">{error}</p>}
      {isPending && <p className="loading">Loading...</p>}
      {recipe && (
        <>
          <h2 className="page-title">{recipe.title}</h2>
          <p>Takes {recipe.cooking_time} to make.</p>
          <ul>
            {recipe.ingredients.map((ing) => (
              <li key={ing}>{ing}</li>
            ))}
          </ul>
          <p className="description">{recipe.description}</p>
          <p className="instructions">{recipe.instructions}</p>
        </>
      )}
    </div>
  );
}
