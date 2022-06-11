package request

import "encoding/json"

type Book struct {
	Name     string      `json:"name" binding:"required"`
	Price    json.Number `json:"price" binding:"required"`
	SubTitle string      `json:"sub_title"`
}

type BookV2 struct {
	Title     string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required"`
	SubTitle string      `json:"sub_title"`
}