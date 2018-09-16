package controllers

import (
	"net/http"

	"github.com/cedricmar/go-api/models"
)

// HelloController extends Controller
type HelloController struct {
	Controller
}

// Index page
func (c *HelloController) Index(w http.ResponseWriter, r *http.Request) {
	payload := models.GetUsers()
	c.JSONResponse(w, r, payload)
}
