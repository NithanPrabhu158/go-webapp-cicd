
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestStaticFileServing(t *testing.T) {
    // Set up the handler as in main.go
    fs := http.FileServer(http.Dir("landing_page"))
    handler := http.StripPrefix("/", fs)

    // Test serving index.html
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", rr.Code)
    }

    // You can add more tests for other static files if needed
}