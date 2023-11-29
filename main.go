package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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


func getAllArticles(w http.ResponseWriter, r *http.Request) {
	allArticles := mongo_service.GetArticlesFromCollection("medical_articles")

	jsonData, err := json.Marshal(allArticles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshaling JSON: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
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
	//router := gin.Default()
	//router.GET("/getAllArticles", getAllArticles)
	//router.GET("/getOneArticle:id", getOneArticle)
	//router.GET("/deleteOneArticle:id", deleteOneArticle)
	//router.GET("/addArticle/:title/:link", addArticle)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/getAllArticles", getAllArticles)


	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return 
	}

}
