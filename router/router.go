package router

import (
	"net/http"

	"github.com/redmejia/orange/app"
)

func Router(app *app.App) http.Handler {
	mux := http.NewServeMux()

	var fs = http.FileServer(http.Dir("static"))
	mux.Handle("/art/img/", http.StripPrefix("/art/img/", fs))

	mux.HandleFunc("/", app.HealthCheckHandler)
	mux.HandleFunc("/profiles", app.ProfileHandler)

	return mux
}
