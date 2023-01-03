package model

type Link struct {
	Id int    `json:"id" gorm:"autoIncrement"`
	Url string `json:"url" gorm:"primaryKey;unique;not null"`
	ShortUrl string `json:"shortUrl"`
}
