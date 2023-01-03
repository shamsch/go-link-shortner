package model

type Link struct {
	Id int    `json:"id" gorm:"autoIncrement"`
	Url string `json:"url" gorm:"not null"`
	ShortUrl string `json:"shortUrl" gorm:"primaryKey;unique;not null"`
}
