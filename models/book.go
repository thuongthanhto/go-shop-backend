package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
