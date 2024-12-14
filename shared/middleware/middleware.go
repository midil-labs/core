package middleware

import (
	"log"
	"net/http"
	"time"
    "github.com/rs/zerolog"
    "github.com/rs/xid"
    "context"
    "github.com/midil-labs/core/pkg/logging"
)

type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}


func RequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        l := logger.Get()

        correlationID := xid.New().String()

        type contextKey string
        const correlationIDKey contextKey = "correlation_id"
        ctx := context.WithValue(r.Context(), correlationIDKey, correlationID)

        r = r.WithContext(ctx)

        l.UpdateContext(func(c zerolog.Context) zerolog.Context {
            return c.Str("correlation_id", correlationID)
        })

        w.Header().Add("X-Correlation-ID", correlationID)

        lrw := newLoggingResponseWriter(w)

        r = r.WithContext(l.WithContext(r.Context()))

        defer func() {
            panicVal := recover()
            if panicVal != nil {
                lrw.statusCode = http.StatusInternalServerError
                panic(panicVal)
            }

            l.
                Info().
                Str("method", r.Method).
                Str("url", r.URL.RequestURI()).
                Str("user_agent", r.UserAgent()).
                Int("status_code", lrw.statusCode).
                Dur("elapsed_ms", time.Since(start)).
                Msg("incoming request")
        }()

        next.ServeHTTP(lrw, r)
    })
}


func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("PANIC: %v", rec)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
