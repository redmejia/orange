package app

import (
	"encoding/json"
	"net/http"

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
	}

	w.Header().Set("Content-Type", "application/json")
	var resp = map[string]interface{}{
		"data": data,
	}
	a.Info.Println(resp)
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
