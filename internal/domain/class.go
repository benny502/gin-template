package domain

type Class struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Items []Item `json:"items"`
}
