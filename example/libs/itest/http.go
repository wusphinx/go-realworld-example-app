package itest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// refer: https://github.com/gin-gonic/gin/blob/master/routes_test.go
func PerformRequest(r http.Handler, method, path string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
