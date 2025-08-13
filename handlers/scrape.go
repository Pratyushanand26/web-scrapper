package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Scrape(c *gin.Context, DB *gorm.DB) {
	var input ScrapeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := GoogleScrape("abcd", "en", "com", 1, 30)

	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}

func BuildGoogleUrls(searchTerm string, languageCode string, countryCOde string, pages int, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := GoogleDomains[countryCOde]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
		}
	} else {
		fmt.Errorf("country code %s not supported", countryCOde)
		return nil, errors.New("cannot build google query")
	}
	return toScrape, nil
}

func GoogleScrape(searchTerm string, languageCode string, countryCode string, pages int, count int) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := BuildGoogleUrls(searchTerm, languageCode, countryCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _,page:=range googlePages{
		res,err:=ScrapeClientRequest(page,Proxystring)
		if err!=nil{
			return nil,err
		}
	}
}
