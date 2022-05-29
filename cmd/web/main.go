package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sir-seju/bookings/pkg/config"
	"github.com/sir-seju/bookings/pkg/handlers"
	"github.com/sir-seju/bookings/pkg/render"
)

const port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the application's main function
func main() {
	app.InProduction = false
	

	session   = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Listening on port", port)
	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
