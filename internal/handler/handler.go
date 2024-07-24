package handler

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"hl3-projectmanagement/internal/service/management"
	"hl3-projectmanagement/pkg/server/router"

	"hl3-projectmanagement/docs"
	"hl3-projectmanagement/internal/handler/http"
)

type Handler struct {
	HTTP *chi.Mux
}

func New(service *management.Service) (h *Handler, err error) {
	h = &Handler{
		HTTP: router.New(),
	}
	docs.SwaggerInfo.BasePath = "/"
	h.HTTP.Get("/swagger/*", httpSwagger.WrapHandler)

	userHandler := http.NewUserHandler(service)
	taskHandler := http.NewTaskHandler(service)
	projectHandler := http.NewProjectHandler(service)

	h.HTTP.Route("/", func(r chi.Router) {
		r.Mount("/user", userHandler.Routes())
		r.Mount("/task", taskHandler.Routes())
		r.Mount("/project", projectHandler.Routes())
	})

	return
}
