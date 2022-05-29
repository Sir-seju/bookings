package handlers

import (
	"net/http"

	"github.com/sir-seju/bookings/pkg/config"
	"github.com/sir-seju/bookings/pkg/models"
	"github.com/sir-seju/bookings/pkg/render"
)

// A variable to def ine the storage of new handlers
var Repo *Repository

// Defines the repository type
type Repository struct {
	App *config.AppConfig
}

// Instantiates a new app in the Repository struct
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html", &models.TemplateData{})
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Write some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, World!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	
	// Pass data into template 
	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
