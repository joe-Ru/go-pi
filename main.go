package main

import (
	"github.com/gin-gonic/gin"
	"log"

	// "context"
	// "fmt"
	// "log"
	//"encoding/json"
	"mongo_connection_test/mongo_service"
	"net/http"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type Articles struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}


func getAllArticles(c *gin.Context) {
	allArticles := mongo_service.GetArticlesFromCollection("medical_articles")
	c.IndentedJSON(http.StatusOK, allArticles)
}


func getOneArticle(c *gin.Context) {
	id := c.Param("id")[1:]
	oneArticle := mongo_service.GetOneArticleFromCollection("medical_articles", id)
	c.IndentedJSON(http.StatusOK, oneArticle)
}


func deleteOneArticle(c *gin.Context) {
	id := c.Param("id")[1:]
	mongo_service.DeleteOneArticleFromCollection("medical_articles", id)
	c.IndentedJSON(http.StatusOK, "Data Deleted")
}


func addArticle(c *gin.Context) {
	title := c.Param("title")
	link := c.Param("link")

	log.Println(title)
	log.Println(link)

	//mongo_service.AddArticleFromCollection("medical_articles", title, link)
	c.IndentedJSON(http.StatusOK, "Data Added")
}



func main() {
	router := gin.Default()
	router.GET("/getAllArticles", getAllArticles)
	router.GET("/getOneArticle:id", getOneArticle)
	router.GET("/deleteOneArticle:id", deleteOneArticle)
	router.GET("/addArticle/:title/:link", addArticle)
	router.Run("localhost:8080")

}
