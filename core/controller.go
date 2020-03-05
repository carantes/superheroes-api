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

// HandleError write error on response and return false if there is no error
func (c *Controller) HandleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	switch err := err.(type) {
	case *HTTPError:
		c.SendJSON(w, &err, err.Status)
	default:
		msg := map[string]string{
			"message": "An error occured",
		}
		c.SendJSON(w, &msg, http.StatusInternalServerError)
	}

	return true
}

// GetContent read request body and return interface
func (c *Controller) GetContent(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}
