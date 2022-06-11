package book

import "encoding/json"

type BookInput struct {
	Name     string      `json:"name" binding:"required"`
	Price    json.Number `json:"price" binding:"required"`
	SubTitle string      `json:"sub_title"`
}

type BookInputV2 struct {
	Title     string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required"`
	SubTitle string      `json:"sub_title"`
}