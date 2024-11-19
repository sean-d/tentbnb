package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/sean-d/tentbnb/pkg/config"
	"github.com/sean-d/tentbnb/pkg/handlers"
	"github.com/sean-d/tentbnb/pkg/render"
	"log"
	"net/http"
	"time"
)

// Place this at the package level so middleware (also in package main) can access it.
var app config.AppConfig

// for accessing/mutating the appConfig's Session
var session *scs.SessionManager

func main() {

	/// Change to true for production. This controls if cookies are secure or not for CSRF and sessions
	app.InProduction = false

	// Sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	app.TemplateCache = templateCache

	// dev server: false. cache gets reloaded every page reload
	// prod: true. we assume changes are done
	app.UseCache = false

	// Provide access to the application config to the render package.
	render.NewTemplates(&app)

	// create the repo
	repo := handlers.NewRepo(&app)

	// pass the repo back to handlers so the local Repo var can hold the repository data.
	handlers.Repo.SetRepository(repo)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("started up: http://localhost:8080")

	// serve represents the http Server struct that is defined with our listening port and a handler which
	// is all of our routes and our appConfig passed in
	serve := http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	//err = http.ListenAndServe(":8080", nil)
	//
	//if err != nil {
	//	log.Fatalf("Error: %s", err)
	//}
}
