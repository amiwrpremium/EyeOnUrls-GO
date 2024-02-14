package controllers

import (
	"EyeOnUrls/models"
	"EyeOnUrls/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchUrl GET /:short_url
// Redirect to the original URL from the short URL
// @Summary Redirect to the original URL from the short URL
// @Description Redirect to the original URL from the short URL
// @ID fetch-url
// @Produce  json
// @Param short_url path string true "Short URL"
// @Success 301 {string} string "Redirect to the original URL"
// @Failure 400 {string} string "Record not found!"
// @Router /{short_url} [get]
func FetchUrl(c *gin.Context) {
	if "favicon.ico" == c.Param("short_url") {
		return
	}

	var urlData models.Url

	if err := models.DB.Where("short_url = ?", c.Param("short_url")).First(&urlData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Increment the number of visits
	urlData.ViewCount++
	models.DB.Save(&urlData)

	c.Redirect(http.StatusMovedPermanently, urlData.OriginalUrl)
	return
}

// CreateUrl POST /
// Create a new URL with a short URL
// @Summary Create a new URL with a short URL
// @Description Create a new URL with a short URL
// @ID create-url
// @Accept  json
// @Produce  json
// @Param input body models.CreateUrlInput true "URL input"
// @Success 200 {object} models.Url "data"
// @Failure 400 {string} string "error"
// @Router / [post]
func CreateUrl(c *gin.Context) {
	var input models.CreateUrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.IsValidURL(input.OriginalUrl) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	input.OriginalUrl = utils.CompleteURL(input.OriginalUrl)

	shortUrl := models.GenerateShortUrl()

	url := models.Url{OriginalUrl: input.OriginalUrl, ShortUrl: shortUrl}
	models.DB.Create(&url)

	url.ShortUrl = c.Request.Host + "/" + url.ShortUrl

	c.JSON(http.StatusOK, gin.H{"data": url})
	return
}
