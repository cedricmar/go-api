package controllers

import (
	"net/http"

	"apps/api/models"
)

// HelloController extends Controller
type HelloController struct {
	Controller
}

// Index page
func (c *HelloController) Index(w http.ResponseWriter, r *http.Request) {
	payload := models.NewPage("Index Page", "en", "Test Page")
	c.Render(w, r, "index", payload)
}
