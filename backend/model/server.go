package model

type Link struct {
	Id int    `json:"id" gorm:"primaryKey"`
	Url string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}
