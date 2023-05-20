package handlers

import (
	"net/http"

	"github.com/smarkericmark/BNBproject/pkg/config"
	"github.com/smarkericmark/BNBproject/pkg/models"
	"github.com/smarkericmark/BNBproject/pkg/render"
)

//TemplateData holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About is the about section

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringmap := make(map[string]string)
	stringmap["test"] = "Hello this is passed data"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringmap,
	})
}
