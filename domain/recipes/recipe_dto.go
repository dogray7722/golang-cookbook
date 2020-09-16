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
	servingSize string `json:"serving_size"`
	item        string `json:"item"`
}
