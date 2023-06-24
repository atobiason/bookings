package main

import (
	"fmt"
	"gethub.com/atobiason/bookings/pkg/config"
	"gethub.com/atobiason/bookings/pkg/handlers"
	"gethub.com/atobiason/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	app.UseCache = false

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	//	http.HandleFunc("/", handlers.Repo.Home)
	//	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on port number %s", portNumber))
	//	http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
