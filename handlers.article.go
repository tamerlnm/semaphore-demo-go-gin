package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			//c.HTML(
			//	// Set the HTTP status to 200 (OK)
			//	http.StatusOK,
			//	// Use the article.html template
			//	"article.html",
			//	// Pass the data that the page uses
			//	gin.H{
			//		"title":   article.Title,
			//		"payload": article,
			//	},
			//)
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")
		} else {
			// If the article is not found, abort with an error
			log.Println(c.AbortWithError(http.StatusNotFound, err))
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
