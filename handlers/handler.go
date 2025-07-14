package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/its-asif/go-url-shortener/config"
	"github.com/its-asif/go-url-shortener/models"

	"github.com/gorilla/mux"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	tempUrl := make([]byte, length)
	for i := range tempUrl {
		tempUrl[i] = charset[rand.Intn(len(charset))]
	}
	return string(tempUrl)
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	shortCode := generateShortCode(6)
	_, err := config.DB.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, req.URL)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": "http://localhost:8000/" + shortCode,
	})
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]

	var url models.URL
	err := config.DB.Get(&url, "SELECT * FROM urls WHERE TRIM(short_code) = $1", code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
