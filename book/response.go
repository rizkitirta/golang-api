package book

type BookResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"judul"`
	Description string    `json:"deskripsi"`
	Price       int       `json:"harga"`
	Rating      int       `json:"rating"`
}

func ResponseConverter(b Book) BookResponse {
	return BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}
}
