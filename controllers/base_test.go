package controllers_test

import (
	"apps/api/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestController(t *testing.T) {
	fmt.Println("Given a server and a controller")
	c := common.Controller{}
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	fmt.Println("When index handler is called")
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Message string `json:"message"`
		}{
			Message: "Hello",
		}

		c.Render(w, nil, "index", &v)
	})
	resp, err := http.Get(server.URL + "/index")

	fmt.Println("Then response should be correct")
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// 200
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 got", resp.StatusCode)
	}
	// Body
	s := []string{"Payload", "running on"}
	for _, p := range s {
		if !strings.Contains(string(body), p) {
			t.Errorf("Expected to contain \"%v\"", p)
		}
	}
}
