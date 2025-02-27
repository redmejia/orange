package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/orange/app"
	"github.com/redmejia/orange/router"
)

func main() {

	var port string
	defaultPort := "8080"
	flag.StringVar(&port, "port", defaultPort, "server port")
	flag.Parse()

	app := &app.App{
		Port:     fmt.Sprintf(":%s", port),
		Info:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:    app.Port,
		Handler: router.Router(app),
	}

	app.Info.Println("Server is running on", app.Port)
	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println("Error starting server", err)
	}

}
