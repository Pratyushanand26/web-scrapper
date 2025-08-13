package handlers

import (
	"math/rand"
	"time"
)

type RegisterInput struct{
  Username   string  `json:"username" binding:"required,min=3"`
  Email      string  `json:"email" binding:"required,email"`
  Password   string  `json:"password" binding:"required,min=6"`
}

type LoginInput struct{
  Email      string  `json:"email" binding:"required,email"`
  Password   string  `json:"password" binding:"required,min=6"`
}

type ScrapeInput struct{
  Text      string   `json:"text" binding:"max:50"`
}

var GoogleDomains = map[string]string{

}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDes   string
}

var UserAgents = []string{}

func RandomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randnum := rand.Int() % len(UserAgents)
	return UserAgents[randnum]
}
