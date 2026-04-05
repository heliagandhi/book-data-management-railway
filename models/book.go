package models

type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	ReleaseYear int      `json:"release_year"`
	Price       float64  `json:"price"`
	TotalPage   int      `json:"total_page"`
	Thickness   string   `json:"thickness"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"category"`
}