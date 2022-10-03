package routes

import (
	"literatur/handlers"
	"literatur/pkg/middleware"
	"literatur/pkg/mysql"
	"literatur/repositories"

	"github.com/gorilla/mux"
)

func CollectionRoutes(r *mux.Router) {
	collectionRepository := repositories.RepositoryCollection(mysql.DB)
	h := handlers.HandlerCollection(collectionRepository)

	r.HandleFunc("/collections", h.FindCollections).Methods("GET")
	r.HandleFunc("/collection/{id}", h.GetCollection).Methods("GET")
	r.HandleFunc("/collection", middleware.Auth(h.CreateCollection)).Methods("POST")
	r.HandleFunc("/collection/{id}", middleware.Auth(h.UpdateeCollection)).Methods("PATCH")
	r.HandleFunc("/collection/{id}", middleware.Auth(h.DeleteCollection)).Methods("DELETE")
}
