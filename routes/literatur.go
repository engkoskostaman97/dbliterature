package routes

import (
	"literatur/handlers"
	"literatur/pkg/middleware"
	"literatur/pkg/mysql"
	"literatur/repositories"

	"github.com/gorilla/mux"
)

func LiteraturRoutes(r *mux.Router) {
	literaturRepository := repositories.RepositoryLiteratur(mysql.DB)
	h := handlers.HandlerLiteratur(literaturRepository)

	r.HandleFunc("/literaturs", middleware.Auth(h.FindLiteraturs)).Methods("GET")
	r.HandleFunc("/literatur/{id}", h.GetLiteratur).Methods("GET")
	r.HandleFunc("/literatur", middleware.Auth(middleware.UploadFile(h.CreateLiteratur))).Methods("POST")

}
