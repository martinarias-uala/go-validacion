package models

type Shape struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	A         string `json:"a"`
	B         string `json:"b"`
	CreatedBy string `json:"created_by"`
}
