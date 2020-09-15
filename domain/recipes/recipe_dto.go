package recipes

type Recipe struct {
	Id           int64    `json:"id"`
	Ingredients  []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
	DateCreated  string   `json:"date_created"`
	Status       string   `json:"status"`
}
