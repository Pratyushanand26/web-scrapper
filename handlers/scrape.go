package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Scrape(c *gin.Context, DB *gorm.DB) {
	var input ScrapeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := GoogleScrape("abcd", "en", "com", nil, 1, 30, 10)

	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}

func BuildGoogleUrls(searchTerm string, languageCode string, countryCode string, pages int, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := GoogleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
		}
	} else {
		fmt.Errorf("country code %s not supported", countryCode)
		return nil, errors.New("cannot build google query")
	}
	return toScrape, nil
}

func GoogleScrape(searchTerm string, languageCode string, countryCode string, Proxystring interface{}, pages int, count int, sleeptime int) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := BuildGoogleUrls(searchTerm, languageCode, countryCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _, page := range googlePages {
		res, err := ScrapeClientRequest(page, Proxystring)
		if err != nil {
			return nil, err
		}
		data, err := GoogleResultparsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}
	return results, nil
}
