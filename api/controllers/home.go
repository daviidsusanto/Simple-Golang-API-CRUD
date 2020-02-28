package controllers

import (
	"SimpleProject/api/responses"
	"net/http"
)

//Home method that show default home api pages
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Awesome API")
}
