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
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, req.Code)
	assert.Equal(t, "Hello, World!", req.Body.String())

}
