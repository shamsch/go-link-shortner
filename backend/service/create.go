package service

import (
	"fmt"
	"link-shortner/db"
	"link-shortner/model"
	"link-shortner/util"
) 

func CreateShortenUrl(url string) string {
	var existingUrl model.Link

	fmt.Println("Checking if this link already exists in the database")

	db.DB.First(&existingUrl, "url = ?", url)
	
	if existingUrl.Url == url {
		fmt.Println("Link already exists")
		return existingUrl.ShortUrl
	}
	
	shortenUrl := util.RandString(8)

	// run a loop to check if the shortenUrl already exists

	for {
		var existingShortUrl model.Link
		db.DB.First(&existingShortUrl, "short_url = ?", shortenUrl)
		if existingShortUrl.ShortUrl == shortenUrl {
			shortenUrl = util.RandString(8)
		} else {
			break
		}
	}

	link := model.Link{
		Url: url,
		ShortUrl: shortenUrl,
	}

	db.DB.Create(&link)
	fmt.Println("Link Created")
	return shortenUrl
}

