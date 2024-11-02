package domain

type Item struct {
	ID          int    `json:"id"`
	ClassId     int    `json:"class_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
