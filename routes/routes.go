package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	LiteraturRoutes(r)
	AuthRoutes(r)
	CollectionRoutes(r)
}
