package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/orange/models"
)

func (a *App) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := []models.Profile{
		{
			Name:    "John Doe",
			Hobbies: []string{"Reading"},
			Image:   "http://localhost:8080/art/img/0x1.jpg",
		},
		{
			Name:    "Sofia Doe",
			Hobbies: []string{"Reading", "Travelling"},
			Image:   "http://localhost:8080/art/img/0x0.jpg",
		},
		{
			Name:    "Rebecca Doe",
			Hobbies: []string{"Swimming", "Travelling"},
			Image:   "http://localhost:8080/art/img/0x3.jpeg",
		},
		{
			Name:    "Jane Doe",
			Hobbies: []string{"Swimming", "Travelling"},
			Image:   "http://localhost:8080/art/img/0x4.jpg",
		},
		{
			Name:    "Mary Doe",
			Hobbies: []string{"Programming", "C++", "Go", "Kotlin", "Java", "Python", "Swimming", "Travelling", "Reading", "Cooking", "Dancing", "Singing", "Playing"},
			Image:   "http://localhost:8080/art/img/0x5.jpg",
		},
		{
			Name:    "Amy Doe",
			Hobbies: []string{"Travelling", "Reading", "Cooking", "Dancing", "Singing", "Playing"},
			Image:   "http://localhost:8080/art/img/0x6.jpg",
		},
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 2
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(data) {
		start = len(data)
	}

	if end > len(data) {
		end = len(data)
	}

	pagedData := data[start:end]

	w.Header().Set("Content-Type", "application/json")
	var resp = map[string]interface{}{
		"data":       pagedData,
		"page":       page,
		"limit":      limit,
		"total":      len(data),
		"totalPages": len(data) / limit,
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	a.Info.Println(string(b))
	json.NewEncoder(w).Encode(resp)
}

func (a *App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.HealthCheck{
		Status:  "OK",
		Version: "1.0.0",
		AppName: "Orange API",
		Endpoints: []string{
			"/",
			"/profiles",
		},
	})
}
