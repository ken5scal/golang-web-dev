package controllers

import (
	"html/template"
	"net/http"
	"github.com/golang-web-dev/042_mongodb/10_hands-on/sessions"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c *Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := sessions.GetUser(w, req)
	sessions.ShowSessions() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (c *Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := sessions.GetUser(w, req)
	if !sessions.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}