package recipes

type Recipe struct {
	Id           int64        	`json:"id"`
	Title        string       	`json:"title"`
	Description  string       	`json:"description"`
	CookingTime	 string 				`json:"cooking_time"`
	Ingredients  []string 			`json:"ingredients"`
	Instructions string       	`json:"instructions"`
	DateCreated  string       	`json:"-"`
}