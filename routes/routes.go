package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personalidades-api/controllers/personalities"
	"personalidades-api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/personalities", personalities.List).Methods("GET")
	api.HandleFunc("/personalities/{id}", personalities.Find).Methods("GET")
	api.HandleFunc("/personalities", personalities.Create).Methods("POST")
	api.HandleFunc("/personalities/{id}", personalities.Delete).Methods("DELETE")
	api.HandleFunc("/personalities/{id}", personalities.Update).Methods("PUT")

	portEnv := os.Getenv("PORT")

	port := fmt.Sprintf(
		":%s",
		portEnv,
	)

	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
