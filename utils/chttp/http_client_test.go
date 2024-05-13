package chttp_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qnify/api-server/utils/chttp"
	"github.com/qnify/api-server/utils/consts"
	. "github.com/qnify/api-server/utils/helper"
)

func TestGet(t *testing.T) {
	expectedResp := `{"message": "success"}`
	tokenSecret := "super secret"
	urlPath := "/test/one/two"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Check(r.Method == "GET", t)
		Check(r.URL.String() == urlPath, t)
		Check(r.Header.Get(consts.AuthHeader) == tokenSecret, t)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedResp))
	}))
	defer server.Close()

	client := chttp.NewClient(server.URL, true)
	resp, body, err := client.Get(urlPath, map[string]string{consts.AuthHeader: tokenSecret})

	NoErr(err, t)
	Check(resp.StatusCode == http.StatusOK, t)
	Check(string(body) == expectedResp, t)
}

func TestPost(t *testing.T) {
	expectedBody := "Hello World"
	expectedResp := `{"message": "success"}`
	urlPath := "/test"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Check(r.Method == "POST", t)
		Check(r.URL.String() == urlPath, t)

		body, err := io.ReadAll(r.Body)
		NoErr(err, t)
		Check(string(body) == expectedBody, t)
		Check(r.Header.Get(consts.ContentType) == consts.TextType, t)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(expectedResp))
	}))
	defer server.Close()

	client := chttp.NewClient(server.URL, false)
	payload := bytes.NewBuffer([]byte(expectedBody))
	resp, body, err := client.Post(urlPath, map[string]string{consts.ContentType: consts.TextType}, payload)

	NoErr(err, t)
	Check(resp.StatusCode == http.StatusCreated, t)
	Check(string(body) == expectedResp, t)
}

func TestPostJson(t *testing.T) {
	payloadKey := "key"
	expectedPayload := map[string]string{payloadKey: "value"}
	expectedResp := `{"message": "success"}`
	urlPath := "/test/path"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Check(r.Method == "POST", t)
		Check(r.URL.String() == urlPath, t)

		body, err := io.ReadAll(r.Body)
		NoErr(err, t)
		Check(r.Header.Get(consts.ContentType) == consts.JsonType, t)
		payload := make(map[string]string)
		err = json.Unmarshal(body, &payload)
		NoErr(err, t)
		Check(payload[payloadKey] == expectedPayload[payloadKey], t)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(expectedResp))
	}))
	defer server.Close()

	client := chttp.NewClient(server.URL, false)
	resp, body, err := client.PostJson(urlPath, expectedPayload)

	NoErr(err, t)
	Check(resp.StatusCode == http.StatusCreated, t)
	Check(string(body) == expectedResp, t)
}

func TestRetry(t *testing.T) {
	reqCount := 0
	expectedResp := "OK"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if reqCount == 2 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedResp))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		reqCount++
	}))
	defer server.Close()

	client := chttp.NewClient(server.URL, true)
	resp, body, err := client.Get("/test", nil)
	NoErr(err, t)
	Check(resp.StatusCode == http.StatusOK, t)
	Check(string(body) == expectedResp, t)
}
