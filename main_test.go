package main


import (
	"log"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go_pi/mongo_service"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)


func GetAllArticlesTest(t *testing.T) {
	req := httptest.NewRequest("GET", "/getAllArticles", nil)
	rec := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())

}