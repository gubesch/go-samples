package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()


	hf := http.HandlerFunc(handler)


	hf.ServeHTTP(recorder, req)


	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


	expected := `Hello Devs!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouter(t *testing.T){
	router := newRouter()
	mockServer := httptest.NewServer(router)

	response, error := http.Get(mockServer.URL + "/hello")

	if error != nil {
		t.Fatal(error)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", response.StatusCode)
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		t.Fatal(error)
	}

	responseString := string(body)
	expected := "Hello Devs!"

	if responseString != expected {
		t.Errorf("Response should be %s, got %s", expected, responseString)
	}

	response.Body.Close()

}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")
	log.Printf(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}


	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
