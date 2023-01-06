package service

import (
	"link-shortner/db"
	"link-shortner/model"
)

func GetUrl(shortUrl string) string {
	var link model.Link
	db.DB.First(&link, "short_url = ?", shortUrl)
	return link.Url
}
