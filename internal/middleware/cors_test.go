package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func corsTester(siteHost string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(CORS(siteHost))
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	return r
}

func TestCORS(t *testing.T) {
	r := corsTester("syfq.my.id")

	cases := []struct {
		name       string
		origin     string
		method     string
		wantStatus int
		wantAllow  string
	}{
		{"same host", "https://syfq.my.id", "GET", 200, "https://syfq.my.id"},
		{"subdomain", "https://waas.syfq.my.id", "GET", 200, "https://waas.syfq.my.id"},
		{"deep subdomain", "http://a.b.syfq.my.id", "GET", 200, "http://a.b.syfq.my.id"},
		{"preflight", "https://syfq.my.id", "OPTIONS", 204, "https://syfq.my.id"},
		{"no origin", "", "GET", 200, ""},
		{"suffix lookalike", "https://evilsyfq.my.id", "GET", 200, ""},
		{"other domain", "https://evil.com", "GET", 200, ""},
		{"other domain preflight", "https://evil.com", "OPTIONS", 403, ""},
		{"port mismatch", "https://syfq.my.id:8443", "GET", 200, ""},
		{"non-http scheme", "ftp://syfq.my.id", "GET", 200, ""},
		{"garbage", "not a url", "GET", 200, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/", nil)
			if tc.origin != "" {
				req.Header.Set("Origin", tc.origin)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", w.Code, tc.wantStatus)
			}
			if got := w.Header().Get("Access-Control-Allow-Origin"); got != tc.wantAllow {
				t.Fatalf("Allow-Origin = %q, want %q", got, tc.wantAllow)
			}
		})
	}
}

func TestCORSDisabledWhenNoHost(t *testing.T) {
	r := corsTester("")
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "https://syfq.my.id")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("Allow-Origin = %q, want empty", got)
	}
}
