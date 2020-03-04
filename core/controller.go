package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Controller handle all base methods
type Controller struct {
}

// SendJSON marshals v to a JSON structure, and send appropriate headers to w
func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error while encoding JSON: %v", err)
		io.WriteString(w, `{"error": "Internal server error"}`)
		return
	}

	w.WriteHeader(code)
	io.WriteString(w, string(b))
}
