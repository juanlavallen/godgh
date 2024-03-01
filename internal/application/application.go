package application

import (
	"net/http"
)

type Application interface {
	Run() error
}

func NewApplication() Application {
	return &application{}
}

type application struct{}

func (a application) Run() error {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})

	return http.ListenAndServe(":8080", r)
}
