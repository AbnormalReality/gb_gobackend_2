package server

import (
	"context"
	"encoding/json"
	"github.com/AbnormalReality/gb_gobackend_2/lesson4/metrics-example/logic"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/trace"
)

type S struct {
	Tr     trace.Tracer
	Server *http.Server
	Logic  *logic.Logic
}

func (s *S) Start() error {
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("my-server"))
	r.HandleFunc("/entities", s.simpleHander()).Methods(http.MethodGet)
	r.HandleFunc("/entities/", s.simpleHander()).Methods(http.MethodPost)

	s.Server = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.Server.ListenAndServe()
}

func (s *S) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

func (s *S) simpleHander() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, span := s.Tr.Start(ctx, "foo")
		defer span.End()

		data := s.Logic.Example(ctx)

		if time.Now().Second()%2 == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{"key": data})
	}
}
