package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go_api/datasources/mysql/users_db"
	"github.com/mrmtsu/go_api/domain/users"
)

func AllGet(w http.ResponseWriter, r *http.Request) {
	users := []users.User{}
	users_db.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]

	user := users.User{}
	users_db.DB.Find(&user, userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	json.NewDecoder(r.Body).Decode(&user)

	users_db.DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	json.NewDecoder(r.Body).Decode(&user)

	users_db.DB.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	users_db.DB.Delete(users.User{}, userId)
	w.WriteHeader(http.StatusNoContent)
}
