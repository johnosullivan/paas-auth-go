package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/johnosullivan/paas-auth-go/controllers"
)

func GetRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/healthz", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(controllers.PingLink)))

	return router
}
