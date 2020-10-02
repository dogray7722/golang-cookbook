package recipes

type Recipe struct {
	Id           int64        `json:"id"`
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	DateCreated  string       `json:"date_created"`
	Status       string       `json:"status"`
}

type Ingredient struct {
	Id          int64  `json:"id"`
	ServingSize string `json:"serving_size"`
	Item        string `json:"item"`
}
