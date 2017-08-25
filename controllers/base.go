package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

// Controller struct
type Controller struct {
}

// Render is used for HTML
func (c *Controller) Render(w http.ResponseWriter, r *http.Request, tmpl string, p interface{}) {
	fmt.Fprintf(w, "Payload: tmpl %v data %v\n", tmpl, p)
	fmt.Fprintf(w, "running on %s with an %s CPU\n", runtime.GOOS, runtime.GOARCH)
}

// JSONResponse for JSON
func (c *Controller) JSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
