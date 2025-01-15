package middleware

import (
	"strings"
	"net/http"
	jhttp "github.com/midil-labs/core/shared/http"
	"github.com/midil-labs/core/shared/jsonapi/response"

)

// Heartbeat endpoint middleware useful to setting up a path like
// `/ping` that load balancers or uptime testing external services
// can make a request before hitting any routes. It's also convenient
// to place this above ACL middlewares as well. 
// Inspired by chi heatbeat middleware
func Heartbeat(endpoint string) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
				response := response.NewJsonAPIResponse(&response.Resource{})
				jhttp.OK(*response)(w)
				w.WriteHeader(http.StatusOK)
				return
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}