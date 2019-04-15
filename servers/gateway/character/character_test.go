package character

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSummaryHandler(t *testing.T) {
	//verify that response has
	// - correct response status code
	// - correct Content-Type header
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/summary?url=https://ddb.ac/characters/10750901/CtT69P", nil)

	CharacterHandler(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("incorrect response status code: expected %d but got %d", http.StatusOK, resp.Code)
	}
	expectedctype := "application/json"
	ctype := resp.Header().Get("Content-Type")
	if len(ctype) == 0 {
		t.Errorf("No `Content-Type` header found in the response: must be there start with `%s`", expectedctype)
	} else if !strings.HasPrefix(ctype, expectedctype) {
		t.Errorf("incorrect `Content-Type` header value: expected it to start with `%s` but got `%s`", expectedctype, ctype)
	}
}
