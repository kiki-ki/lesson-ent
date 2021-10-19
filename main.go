package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kiki-ki/lesson-ent/controller"
	"github.com/kiki-ki/lesson-ent/database"
)

func main() {
	entClient := database.NewEntClient()
	defer entClient.Close()
	entClient.Migrate()

	r := chi.NewRouter()
	setRouter(r, entClient)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func setRouter(r chi.Router, dbc *database.EntClient) {
	// logging
	r.Use(middleware.Logger)

	// routing
	companyController := controller.NewCompanyController(dbc)
	r.Route("/companies", func(r chi.Router) {
		r.Get("/{companyId}", companyController.Show)
		r.Put("/{companyId}", companyController.Update)
		r.Delete("/{companyId}", companyController.Delete)
		r.Get("/{companyId}/users", companyController.IndexUsers)
		r.Post("/", companyController.CreateWithUser)
	})
}
