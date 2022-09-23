package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testRenderBody = []string{"foo", "bar"}

func hasHeader(m http.Header, h, v string) (yes bool) {
	t, ok := m[h]
	if !ok {
		return
	}

	if len(t) < 1 {
		return
	}

	return t[0] == v
}

func TestRenderStringsPlain(t *testing.T) {
	t.Parallel()

	var (
		req = httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		rsp httptest.ResponseRecorder
	)

	if err := RenderStrings(&rsp, req, testRenderBody); err != nil {
		t.Fatalf("error: %v", err)
	}

	res := rsp.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %s", res.Status)
	}

	if !hasHeader(res.Header, contentTypeHeader, mimeTEXT) {
		t.Fatal("unexpected content-type")
	}
}

func TestRenderStringsJSON(t *testing.T) {
	t.Parallel()

	var (
		req = httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		rsp httptest.ResponseRecorder
	)

	req.Header.Set(acceptHeader, mimeJSON)

	if err := RenderStrings(&rsp, req, testRenderBody); err != nil {
		t.Fatalf("error: %v", err)
	}

	res := rsp.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %s", res.Status)
	}

	if !hasHeader(res.Header, contentTypeHeader, mimeJSON) {
		t.Fatal("unexpected content-type")
	}
}

func TestRenderTEXT(t *testing.T) {
	t.Parallel()

	var (
		buf bytes.Buffer
		rsp = httptest.ResponseRecorder{Body: &buf}
	)

	if err := renderTEXT(&rsp, testRenderBody); err != nil {
		t.Fatalf("error: %v", err)
	}

	res := rsp.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %s", res.Status)
	}

	if !hasHeader(res.Header, contentTypeHeader, mimeTEXT) {
		t.Fatal("unexpected content-type")
	}

	body := buf.String()

	for _, tv := range testRenderBody {
		if !strings.Contains(body, tv) {
			t.Fatalf("not contains: %s", tv)
		}
	}
}

func TestRenderJSON(t *testing.T) {
	t.Parallel()

	var (
		buf bytes.Buffer
		rsp = httptest.ResponseRecorder{Body: &buf}
	)

	if err := renderJSON(&rsp, testRenderBody); err != nil {
		t.Fatalf("error: %v", err)
	}

	res := rsp.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %s", res.Status)
	}

	if !hasHeader(res.Header, contentTypeHeader, mimeJSON) {
		t.Fatal("unexpected content-type")
	}

	var body []string

	if err := json.NewDecoder(&buf).Decode(&body); err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}

	set := make(map[string]struct{})

	for _, v := range body {
		set[v] = struct{}{}
	}

	for _, tv := range testRenderBody {
		if _, ok := set[tv]; !ok {
			t.Fatalf("not contains: %s", tv)
		}
	}
}
