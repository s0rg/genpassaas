package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	acceptHeader      = "Accept"
	mimeTEXT          = "text/plain"
	mimeJSON          = "application/json"
)

func RenderStrings(
	w http.ResponseWriter,
	r *http.Request,
	s []string,
) error {
	if r.Header.Get(acceptHeader) == mimeJSON {
		return renderJSON(w, s)
	}

	return renderTEXT(w, s)
}

func renderTEXT(w http.ResponseWriter, s []string) error {
	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		b.WriteString(s[i])
		b.WriteString("\n")
	}

	w.Header().Set(contentTypeHeader, mimeTEXT)

	if _, err := b.WriteTo(w); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func renderJSON(w http.ResponseWriter, s []string) error {
	w.Header().Set(contentTypeHeader, mimeJSON)

	if err := json.NewEncoder(w).Encode(s); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}
