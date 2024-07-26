package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := responseWriter{ResponseWriter: w}
		next.ServeHTTP(&rw, r)

		s.Logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   rw.status,
			"duration": time.Since(start),
		}).
			Info("Request done")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
