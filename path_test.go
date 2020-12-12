package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestProxy(t *testing.T) {
	r := SetupRouter()
	//w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

}
