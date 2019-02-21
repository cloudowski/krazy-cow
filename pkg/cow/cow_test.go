package cow

import (
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

// TODO - fix after resolving issues with reading conf from cow
func TestSay(t *testing.T) {

	cowname := "testcow"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	c := Cow{Name: cowname}
	handler := http.HandlerFunc(c.Say)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected http code %v, got %v", http.StatusOK, rr.Code)
	}

	// cowregex := fmt.Sprintf("\"%v\"", cowname)
	cowregex := "FIXME"
	r, _ := regexp.Compile(cowregex)

	if !r.MatchString(rr.Body.String()) {
		// t.Errorf("Could not find text matching regexp %v", cowregex)
		t.Errorf("Got %v, did not match regexp %v", rr.Body.String(), cowregex)
	}

}
