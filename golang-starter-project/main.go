package main

import (
	"fmt"
	"golang-starter-project/configs"
	"golang-starter-project/controllers"
	"golang-starter-project/infrastructures"
	"golang-starter-project/repositories"
	"golang-starter-project/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	customMiddleware "golang-starter-project/middlewares"
)

func initDependencies() controllers.Controller {
	repository := repositories.Repository{}
	service := services.Service{}
	service.SetDependencies(repository)
	controller := controllers.Controller{}
	controller.SetDependencies(service)

	return controller
}

func initMiddlewares(router *chi.Mux) {
	router.Use(middleware.Recoverer)
	router.Use(customMiddleware.CoorelationIdMiddleware)
	router.Use(middleware.Heartbeat("/status"))
	router.Use(middleware.AllowContentType("application/json"))
}

func initRoutes(router *chi.Mux, controller controllers.Controller) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/hello", controller.Hello)
	})
}

func main() {
	configs.LoadAppConfigs()
	controller := initDependencies()
	router := chi.NewRouter()
	initMiddlewares(router)
	initRoutes(router, controller)
	// dsn := configs.GetAppConfig("DB_DSN")
	// infrastructures.InitDBConnection(infrastructures.POSTGRES, dsn)
	infrastructures.InitLogger()
	if err := http.ListenAndServe(fmt.Sprintf(":%v", configs.GetAppConfig("APP_PORT")), router); err != nil {
		panic(err)
	}
}
