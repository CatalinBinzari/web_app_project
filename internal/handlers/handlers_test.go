package handlers

import (
	"net/http"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name                string
	url                 string
	method              string
	params              []postData
	excepctedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	// {"about", "/about", "GET", []postData{}, http.StatusOK},
	// {"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	// {"ms", "/major-suite", "GET", []postData{}, http.StatusOK},
}

// func TestHandlers(t *testing.T) {
// 	routes := getRoutes()
// 	ts := httptest.NewTLSServer(routes) // starts test server
// 	defer ts.Close()                    // defer is executed after current function(TestHandlers) is finished

// 	for _, e := range theTests {
// 		if e.method == "GET" {
// 			resp, err := ts.Client().Get(ts.URL + e.url)
// 			if err != nil {
// 				t.Log(err)
// 				t.Fatal(err)
// 			}

// 			if resp.StatusCode != e.excepctedStatusCode {
// 				t.Errorf("for %s, expected %d, but got %d", e.name, e.excepctedStatusCode, resp.StatusCode)
// 			}
// 		} else {
// 			// TBD
// 		}
// 	}

// }
