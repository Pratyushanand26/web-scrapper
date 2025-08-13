package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

func ScrapeClientRequest(searchURL string,Proxystring interface{})(*http.Response,error){
  baseClient:=getScrapeClient(Proxystring)
  req,_:=http.NewRequest("GET",searchURL,nil)
  req.Header.Set("User-Agent",RandomUserAgent())

  res,err:=baseClient.Do(req)
  if res.StatusCode!=200{
	err:=fmt.Errorf("got a non 200 response suggesting a ban")
	return nil,err
  }

  if err!=nil{
	return nil,err
  }
  return res,nil
}

func getScrapeClient(proxystring interface{}) *http.Client{

	switch v:=proxystring.(type){
	case string:
		proxyUrl,_:=url.Parse(v);
		return &http.Client{Transport:&http.Transport{Proxy:http.ProxyURL(proxyUrl)}}

    default:return &http.Client{}

	}
}
