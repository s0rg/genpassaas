package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/s0rg/genpassaas/pkg/config"
)

const testKey = "key-ok"

func TestGET(t *testing.T) {
	t.Parallel()

	h := GET(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, httptest.NewRequest(http.MethodHead, "/", http.NoBody))

		res := rsp.Result()

		if res.StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("HEAD: unexpected status: %s", res.Status)
		}
	}

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, httptest.NewRequest(http.MethodGet, "/", http.NoBody))

		res := rsp.Result()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("GET: unexpected status: %s", res.Status)
		}
	}
}

func TestBearer(t *testing.T) {
	t.Parallel()

	const keyBAD = "key-bad"

	h := Bearer(testKey, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var req = httptest.NewRequest(http.MethodGet, "/", http.NoBody)

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #1: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, "")

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #2: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, "foo")

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #3: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, "foo bar")

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #3: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, bearerName)

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #4: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, bearerName+" ")

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #5: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, bearerName+" "+keyBAD)

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusForbidden {
			t.Fatalf("unexpected status #6: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, bearerName+" "+testKey)

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("unexpected status #7: %s", res.Status)
		}
	}

	req.Header.Set(headerAuth, bearerName+" "+testKey+" malformed")

	{
		var rsp httptest.ResponseRecorder

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("unexpected status #7: %s", res.Status)
		}
	}
}

func TestWithBearer(t *testing.T) {
	t.Parallel()

	h := withBearer(testKey, func(_ int) (rv string) {
		return testKey
	})

	var req = httptest.NewRequest(http.MethodGet, "/", http.NoBody)

	req.Header.Set(headerAuth, bearerName+" "+testKey)

	{
		var (
			buf bytes.Buffer
			rsp = httptest.ResponseRecorder{Body: &buf}
		)

		h(&rsp, req)

		res := rsp.Result()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status: %s", res.Status)
		}

		if rv := buf.String(); strings.Count(rv, testKey) != config.DefaultCount {
			t.Fatalf("unexepected body: %s", rv)
		}
	}
}

func TestGetInt(t *testing.T) {
	t.Parallel()

	if val, ok := getInt("1"); !ok || val != 1 {
		t.Fatal("fail - 1")
	}

	if _, ok := getInt("A"); ok {
		t.Fatal("fail - A")
	}
}

func TestGenHandler(t *testing.T) {
	t.Parallel()

	var N, C int

	h := genHandler(func(n int) (rv string) {
		if n > N {
			N = n
		}

		C++

		return ""
	})

	testHandler(t, h, 1, 1)

	if N != config.MinLength {
		t.Fatalf("#0 len mismatch got: %d", N)
	}

	if C != config.MinCount {
		t.Fatalf("#0 count mismatch got: %d", C)
	}

	N, C = 0, 0

	testHandler(t, h, config.MaxCount+1, config.MaxLength+1)

	if N != config.MaxLength {
		t.Fatalf("#1 len mismatch got: %d", N)
	}

	if C != config.MaxCount {
		t.Fatalf("#1 count mismatch got: %d", C)
	}
}

func testHandler(t *testing.T, h http.HandlerFunc, count, length int) {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

	q := req.URL.Query()
	q.Set("num", strconv.Itoa(count))
	q.Set("len", strconv.Itoa(length))

	req.URL.RawQuery = q.Encode()

	var r httptest.ResponseRecorder

	h(&r, req)

	res := r.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("num: %d len: %d status: %s", count, length, res.Status)
	}
}
