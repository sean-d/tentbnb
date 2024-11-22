package handlers

import (
	"github.com/sean-d/tentbnb/pkg/config"
	"github.com/sean-d/tentbnb/pkg/models"
	"github.com/sean-d/tentbnb/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository using the app config supplied
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetRepository sets the repository for the handlers using the repository passed in from main
func (rep *Repository) SetRepository(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rep.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.TemplateRender(w, "home.page.gohtml", &models.TemplateData{})

}

// About is the /about handler
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := rep.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := map[string]string{
		"test":         "hello",
		"another test": "goodbye",
		"remote ip":    remoteIP,
	}

	render.TemplateRender(w, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}

// Contact is the contact page handler
func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "contact.page.gohtml", &models.TemplateData{})
}

// Availability is the search availability handler
func (rep *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "search-availability.page.gohtml", &models.TemplateData{})
}

// Reservation is the reservation handler
func (rep *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

// Parry is the parry handler
func (rep *Repository) Parry(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "parry.page.gohtml", &models.TemplateData{})
}

// Simon is the simon handler
func (rep *Repository) Simon(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, "simon.page.gohtml", &models.TemplateData{})
}
