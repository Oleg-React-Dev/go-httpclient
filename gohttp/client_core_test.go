package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.builder = &clientBuilder{headers: commonHeaders}

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	requestHeaders.Set("Authorization", "Bearer ABC-567")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 4 {
		t.Error("Expects 3 headers but got", len(finalHeaders))
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Expects value of X-Request-Id header to equal ABC-123 but got", finalHeaders.Get("X-Request-Id"))
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Expects value of User-Agent header to equal cool-http-client but got", finalHeaders.Get("X-Request-Id"))
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Expects value of Content-Type header to equal application/json but got", finalHeaders.Get("X-Request-Id"))
	}

	if finalHeaders.Get("Authorization") != "Bearer ABC-567" {
		t.Error("Expects value of Authorization header to equal Bearer ABC-567 but got", finalHeaders.Get("X-Request-Id"))
	}

}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("should return nil, nil", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no error expected when passing nil as a body")
		}
		if body != nil && err != nil {
			t.Error("no body expected when passing nil as a body")
		}

	})

	t.Run("should return body as JSON", func(t *testing.T) {
		requestBody := []string{"one", "tow", "three"}
		body, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","tow","three"]` {
			t.Error("invalid json body")
		}

	})

	t.Run("should return body as XML", func(t *testing.T) {
		requestBody := []string{"one", "tow", "three"}
		body, err := client.getRequestBody("application/xml", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice as xml")
		}

		if string(body) != `<string>one</string><string>tow</string><string>three</string>` {
			t.Error("invalid xml body")
		}

	})

	t.Run("should return body as JSON by default", func(t *testing.T) {
		requestBody := []string{"one", "tow", "three"}
		body, err := client.getRequestBody("", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice as xml")
		}

		if string(body) != `["one","tow","three"]` {
			t.Error("invalid json body")
		}

	})

}
