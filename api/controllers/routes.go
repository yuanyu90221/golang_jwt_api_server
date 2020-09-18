package controllers

import (
	"github.com/yuanyu90221/golang_jwt_api_server/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
