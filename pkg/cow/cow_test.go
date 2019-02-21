package cow

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHealthcheck(t *testing.T) {

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	cowHappy := Cow{Name: "testcow"}
	cowHappy.SetMood(100)
	cowMad := Cow{Name: "testcow"}
	cowMad.SetMood(0)

	// check mad cow
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cowMad.Healthcheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected http code %v, got %v", http.StatusBadRequest, rr.Code)
	}

	// check happy cow
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(cowHappy.Healthcheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected http code %v, got %v", http.StatusOK, rr.Code)
	}

}

func TestSay(t *testing.T) {

	cowname := "testcow"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	c := &Cow{Name: cowname}
	s := "I am a test cow"
	c.SetSay(s)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Say)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected http code %v, got %v", http.StatusOK, rr.Code)
	}

	cowregex := fmt.Sprintf("\"%v\"", s)
	r, _ := regexp.Compile(cowregex)

	if !r.MatchString(rr.Body.String()) {
		// t.Errorf("Could not find text matching regexp %v", cowregex)
		t.Errorf("Got %v, did not match regexp %v", rr.Body.String(), cowregex)
	}

}
