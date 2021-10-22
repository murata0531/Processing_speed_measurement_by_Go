package handler

import (
	"net/http"

	"github.com/murata0531/Processing_speed_measurement_by_Go/backend/article"

	"github.com/gin-gonic/gin"
)

func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := articles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type ArticlePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ArticlePost(post *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := ArticlePostRequest{}
		c.Bind(&requestBody)

		item := article.Article{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}