package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required"`
	Description string      `json:"description"`
	Rating      json.Number `json:"rating"`
}

type BookRequestV2 struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required"`
	SubTitle string      `json:"sub_title"`
}
