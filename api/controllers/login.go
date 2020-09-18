package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yuanyu90221/golang_jwt_api_server/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"

	"github.com/yuanyu90221/golang_jwt_api_server/api/auth"
	"github.com/yuanyu90221/golang_jwt_api_server/api/models"
	"github.com/yuanyu90221/golang_jwt_api_server/api/responses"
)

//SignIn route
func (server *Server) SignIn(email, passwd string) (string, error) {
	var err error
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPasswd(user.Passwd, passwd)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}

//Login route
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate(&models.LoginAction{})
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Passwd)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
