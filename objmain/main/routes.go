package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	reviews "github.com/Alejandraarrieta/pruebaGo/reviews/web"
	"net/http"
)

func Routes(
	//sph *web.CreateSmartphoneHandler,
	alumnoHandler *reviews.AlumnoHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	//mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)
	mux.Post("/reviews", alumnoHandler.AddAlumnoHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "miqueas")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}

