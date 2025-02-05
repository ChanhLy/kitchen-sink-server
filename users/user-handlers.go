package users

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	database "go-server/db"
	"go-server/utils"
	"net/http"
	"strconv"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsers(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, users, http.StatusOK)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId, errParseInt := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	if errParseInt != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := GetUserById(userId, r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if user == (database.User{}) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	utils.SendResponse(w, user, http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var createUserParams database.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&createUserParams)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, createUserErr := CreateUser(createUserParams, r.Context())
	if createUserErr != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, user, http.StatusCreated)
}
