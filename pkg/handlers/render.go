package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	mimeTEXT = "text/plain"
	mimeJSON = "application/json"
)

func renderTEXT(w http.ResponseWriter, body []string) (err error) {
	var buf bytes.Buffer

	for i := 0; i < len(body); i++ {
		buf.WriteString(body[i])
		buf.WriteString("\n")
	}

	w.Header().Set("Content-Type", mimeTEXT)
	_, err = buf.WriteTo(w)

	return
}

func renderJSON(w http.ResponseWriter, body []string) error {
	w.Header().Set("Content-Type", mimeJSON)

	return json.NewEncoder(w).Encode(body)
}

func RenderStrings(w http.ResponseWriter, r *http.Request, body []string) error {
	if r.Header.Get("Accept") == mimeJSON {
		return renderJSON(w, body)
	}

	return renderTEXT(w, body)
}
