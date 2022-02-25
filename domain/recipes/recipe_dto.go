package recipes

type Recipe struct {
	Id           int64        `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	CookingTime	 string 			`json:"cooking_time"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	DateCreated  string       `json:"date_created"`
	Status       string       `json:"status"`
}

type Ingredient struct {
	Id          int64  `json:"id"`
	Item        string `json:"item"`
}