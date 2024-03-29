import { useState, useRef, useEffect } from "react";

import { useFetch } from "../../hooks/useFetch";

import { useNavigate } from "react-router-dom";

import "./Create.css";

export default function Create() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [instructions, setInstructions] = useState("");
  const [cookingTime, setCookingTime] = useState("");
  const [newIngredient, setNewIngredient] = useState("");
  const [ingredients, setIngredients] = useState([]);
  const ingredientInput = useRef(null);
  const { postData, data } = useFetch("http://localhost:8080/recipes", "POST");
  const navigate = useNavigate();

  const handleAdd = (e) => {
    e.preventDefault();
    const ing = newIngredient.trim();
    if (ing && !ingredients.includes(ing)) {
      setIngredients((prevIngredients) => [...prevIngredients, ing]);
    }
    setNewIngredient("");
    ingredientInput.current.focus();
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    postData({
      title,
      description,
      ingredients,
      instructions,
      cookingTime
    });
  };

  useEffect(() => {
    if (data) {
      navigate("/");
    }
  }, [navigate, data]);

  return (
    <div className="create">
      <h2 className="page-title">Add a New Recipe</h2>
      <form onSubmit={handleSubmit}>
        <label>
          <span>Recipe Title:</span>
          <input
            type="text"
            onChange={(e) => setTitle(e.target.value)}
            value={title}
            required
          />
        </label>
        <label>
          <span>Recipe Description:</span>
          <textarea
            type="text"
            onChange={(e) => setDescription(e.target.value)}
            value={description}
          />
        </label>
        <label>
          <span>Recipe Ingredients:</span>
          <div className="ingredients">
            <input
              type="text"
              onChange={(e) => setNewIngredient(e.target.value)}
              value={newIngredient}
              ref={ingredientInput}
            />
            <button className="btn" onClick={handleAdd}>
              Add
            </button>
          </div>
        </label>
        <p>
          Current Ingredients: {ingredients.join(", ")}
        </p>
        <label>
          <span>Recipe Instructions:</span>
          <textarea
            onChange={(e) => setInstructions(e.target.value)}
            value={instructions}
            required
          />
        </label>
        <label>
          <span>Cooking Time:</span>
          <input
            type="text"
            onChange={(e) => setCookingTime(e.target.value)}
            value={cookingTime}
            required
          />
        </label>
        <button className="btn">Submit</button>
      </form>
    </div>
  );
}
