package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mrmtsu/go_api/controllers/articles"
	"github.com/mrmtsu/go_api/controllers/users"
	"github.com/mrmtsu/go_api/datasources/mysql/users_db"
)

func sanityCheck() {
	godotenv.Load(".env")

	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	users_db.DbConnect()

	// Users
	router.HandleFunc("/users", users.AllGet).Methods("GET")
	router.HandleFunc("/users/{id}", users.Get).Methods("GET")
	router.HandleFunc("/user", users.Create).Methods("POST")
	router.HandleFunc("/users/{id}", users.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", users.Delete).Methods("DELETE")

	// Articles
	router.HandleFunc("/articles", articles.AllGet).Methods("GET")
	router.HandleFunc("/articles/{id}", articles.Get).Methods("GET")
	router.HandleFunc("/article", articles.Create).Methods("POST")
	router.HandleFunc("/articles/{id}", articles.Update).Methods("PUT")
	router.HandleFunc("/articles/{id}", articles.Delete).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
