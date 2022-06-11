package book

type BookResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"judul"`
	Description string    `json:"deskripsi"`
	Price       int       `json:"harga"`
	Rating      int       `json:"rating"`
}
