package models

type ShortUrl struct {
	ID          string `json:"id" gorm:"primary_key"`
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
	VisitCount  uint   `json:"visitCount"`
}
