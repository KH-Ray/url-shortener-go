package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"url-shortener-be/db"
	"url-shortener-be/models"
	"url-shortener-be/utils"

	"github.com/gin-gonic/gin"
)

type CreateShortURLInput struct {
	Url string `json:"url" binding:"required"`
}

func GetShortUrl(c *gin.Context) {
    var shortUrl models.ShortUrl

    if err := db.DB.Where("id = ?", c.Param("id")).First(&shortUrl).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "URL not found!"}})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id": shortUrl.ID,
        "originalURL": shortUrl.OriginalUrl,
        "visitCount": shortUrl.VisitCount,
    })
}

func CreateShortUrl(c *gin.Context) {
    var shortUrlModels models.ShortUrl
    var input CreateShortURLInput

    if err := c.ShouldBindJSON(&input); err !=  nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error()}})
        return
    }

    if err := db.DB.Where("original_url = ?", input.Url).First(&shortUrlModels).Error; err == nil {
        db.DB.First(&shortUrlModels)
        c.JSON(http.StatusOK, gin.H{
            "id": shortUrlModels.ID,
            "originalURL": shortUrlModels.OriginalUrl,
            "visitCount": shortUrlModels.VisitCount,
        })
        return
    }

    visitCount := uint(0)
    urlKey := utils.GenerateNanoId()
    trimmedUrl := strings.Trim(input.Url, " ")

    if ok := utils.IsValidUrl(trimmedUrl); !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": "URL is not valid!"}})
        return
    }

    hostUrl := c.Request.Host

    if strings.Contains(trimmedUrl, hostUrl) {
        c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": "This URL is already a short URL. Please use an unshortened URL."}})
        return
    }
    
    parsedUrl, err := url.Parse(hostUrl)
    
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error()}})
        return
    }

    shortUrl := fmt.Sprintf("https://%s/%s", parsedUrl, urlKey)

    shortUrlResponse := models.ShortUrl{
        ID: urlKey,
        OriginalUrl: trimmedUrl,
        ShortUrl: shortUrl,
        VisitCount: visitCount,
    }
    db.DB.Create(&shortUrlResponse)

    c.JSON(http.StatusCreated, gin.H{
        "id": urlKey,
        "originalURL": trimmedUrl,
        "visitCount": visitCount,
    })
}

func VisitShortUrl(c *gin.Context) {
    var shortUrl models.ShortUrl

    if err := db.DB.Where("id = ?", c.Param("id")).First(&shortUrl).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "URL not found!"}})
        return
    }

    shortUrl.VisitCount += 1
    db.DB.Save(&shortUrl)

    c.JSON(http.StatusOK, gin.H{
        "id": shortUrl.ID,
        "originalURL": shortUrl.OriginalUrl,
        "visitCount": shortUrl.VisitCount,
    })
}
