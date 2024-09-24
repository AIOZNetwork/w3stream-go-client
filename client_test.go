package w3streamsdk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	mux.HandleFunc("/auth/api-key", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"token_type": "Bearer",
			"expires_in": 3600,
			"access_token": "fakeToken",
			"refresh_token": "fakeToken"
		  }`)
	})

	authCredentials := loadAuthCredentials()
	client = ClientBuilder(authCredentials).Build()
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func createTempFile(filename string, size int64) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}
	return f
}

func loadAuthCredentials() AuthCredentials {
	env, err := os.ReadFile("app.json")
	if err != nil {
		log.Fatal(err)
	}

	var auth AuthCredentials
	if err := json.Unmarshal(env, &auth); err != nil {
		log.Fatal(err)
	}

	return auth
}
