package service

import (
	"fmt"
	"link-shortner/db"
	"link-shortner/model"
	"link-shortner/util"
) 

func CreateShortenUrl(url string) string {
	shortenUrl := util.RandString(8)
	link := model.Link{
		Url: url,
		ShortUrl: shortenUrl,
	}
	//TODO: Check if shortUrl already exists
	db.DB.Create(&link)
	fmt.Println("Link Created")
	return shortenUrl
}

