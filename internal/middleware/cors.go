package middleware

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS allows cross-origin requests from the site root and any of its
// subdomains. Preflights answer 204 and skip the rest of the chain, so they
// are not counted by the rate limiter.
func CORS(siteHost string) gin.HandlerFunc {
	siteHost = strings.ToLower(siteHost)

	// sameSite reports whether rawOrigin is on siteHost or a subdomain of it.
	sameSite := func(rawOrigin string) bool {
		if siteHost == "" {
			return false
		}
		u, err := url.Parse(rawOrigin)
		if err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.Port() != "" {
			return false
		}
		host := strings.ToLower(u.Hostname())
		return host == siteHost || strings.HasSuffix(host, "."+siteHost)
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" || !sameSite(origin) {
			// Not a same-site origin: no CORS grant, response stays blocked.
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			c.Next()
			return
		}

		h := c.Writer.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		h.Set("Access-Control-Max-Age", "600")
		// Origin is echoed per-request, so caches must key on it.
		h.Add("Vary", "Origin")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
