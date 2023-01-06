package service

import (
	"fmt"
	"link-shortner/db"
	"link-shortner/model"
	"link-shortner/util"
) 

func CreateShortenUrl(url string) string {
	var existingUrl model.Link

	db.DB.First(existingUrl, "url = ?", url)
	
	if existingUrl.Url == url {
		fmt.Println("Link already exists")
		return existingUrl.ShortUrl
	}
	
	shortenUrl := util.RandString(8)
	link := model.Link{
		Url: url,
		ShortUrl: shortenUrl,
	}

	db.DB.Create(&link)
	fmt.Println("Link Created")
	return shortenUrl
}

