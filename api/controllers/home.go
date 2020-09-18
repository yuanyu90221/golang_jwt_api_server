package controllers

import (
	"net/http"

	"github.com/yuanyu90221/golang_jwt_api_server/api/responses"
)

//Home route
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to This api")
}
