package handlers

import (
	"gethub.com/atobiason/bookings/pkg/config"
	"gethub.com/atobiason/bookings/pkg/models"
	"gethub.com/atobiason/bookings/pkg/render"
	"log"
	"net/http"
)

// repo is the repository type used by the handlers
var Repo *Repository

// repository is the reporitory type
type Repository struct {
	App *config.AppConfig
}

// Nepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// new handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	log.Println("home remoteIp = ", remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	log.Println("remoteIp = ", remoteIP)
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
