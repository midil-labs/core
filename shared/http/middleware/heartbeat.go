package middleware

import (
	"net/http"
	"strings"

	httppkg "github.com/midil-labs/core/shared/http"
	"github.com/midil-labs/core/shared/jsonapi/common"
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
				data := &response.Data{
					ResourceIdentifier: common.ResourceIdentifier{
						Type: "heartbeat",
					},
					Attributes: map[string]interface{}{
						"Ding": "Dong!",
					},
				}
				response := response.NewJsonAPIResponse[*response.Data](data)
				httppkg.OK(*response)(w)
				return
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
