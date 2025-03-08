package logger

import (
	"log/slog"
	"net/http"
	"os"
	// slogpretty "retsAPI/serv/logger/prettylog"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func NewLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env{
		case "local": log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case "prod": log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		default: log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	
	return log
}

// func setupPrettySlog() *slog.Logger {
// 	opts := slogpretty.PrettyHandlerOptions{
// 		SlogOpts: &slog.HandlerOptions{
// 			Level: slog.LevelDebug,
// 		},
// 	}

// 	handler := opts.NewPrettyHandler(os.Stdout)

// 	return slog.New(handler)
// }

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "middleware/logger"),
		)

		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
