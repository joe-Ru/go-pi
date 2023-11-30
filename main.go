package main

import (
	"log"

	"encoding/json"
	"github.com/go-chi/chi/v5"
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


func getOneArticle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	oneArticle := mongo_service.GetOneArticleFromCollection("medical_articles", idParam)

	log.Println(oneArticle)

	jsonData, err := json.Marshal(oneArticle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshaling JSON: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func deleteOneArticle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	mongo_service.DeleteOneArticleFromCollection("medical_articles", idParam)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item deleted successfully"))

}


func addArticle(w http.ResponseWriter, r *http.Request) {

	var newArticle Articles

	if err := json.NewDecoder(r.Body).Decode(&newArticle); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON: " + err.Error()))
		return
	}

	log.Println(newArticle)

	mongo_service.AddArticleFromCollection("medical_articles", newArticle.Title, newArticle.Link)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item added successfully"))
}



func main() {

	r := chi.NewRouter()
	r.Get("/getAllArticles", getAllArticles)
	r.Get("/getOneArticle/{id}", getOneArticle)
	r.Post("/addOneArticle", addArticle)
	r.Get("/deleteOneArticle/{id}", deleteOneArticle)


	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return 
	}

}
