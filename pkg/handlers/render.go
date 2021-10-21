package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	mimeTEXT = "text/plain"
	mimeJSON = "application/json"
)

func renderTEXT(w http.ResponseWriter, body []string) error {
	var buf bytes.Buffer

	for i := 0; i < len(body); i++ {
		buf.WriteString(body[i])
		buf.WriteString("\n")
	}

	w.Header().Set("Content-Type", mimeTEXT)

	if _, err := buf.WriteTo(w); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func renderJSON(w http.ResponseWriter, body []string) error {
	w.Header().Set("Content-Type", mimeJSON)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

func RenderStrings(w http.ResponseWriter, r *http.Request, body []string) error {
	if r.Header.Get("Accept") == mimeJSON {
		return renderJSON(w, body)
	}

	return renderTEXT(w, body)
}
