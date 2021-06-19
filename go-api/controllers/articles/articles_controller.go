package articles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go_api/datasources/mysql/users_db"
	"github.com/mrmtsu/go_api/domain/articles"
)

func AllGet(w http.ResponseWriter, r *http.Request) {
	articles := []articles.Article{}
	users_db.DB.Preload("User").Find(&articles)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId := params["id"]

	article := articles.Article{}
	users_db.DB.Find(&article, articleId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func Create(w http.ResponseWriter, r *http.Request) {
	article := articles.Article{}
	json.NewDecoder(r.Body).Decode(&article)

	users_db.DB.Create(&article)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func Update(w http.ResponseWriter, r *http.Request) {
	article := articles.Article{}
	json.NewDecoder(r.Body).Decode(&article)

	users_db.DB.Save(&article)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId := params["id"]
	users_db.DB.Delete(articles.Article{}, articleId)
	w.WriteHeader(http.StatusNoContent)
}

func GetShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId := params["id"]

	article := articles.Article{}
	users_db.DB.Find(&article, articleId)
	userId := article.UserId
	joinArticle := []articles.JoinArticle{}
	users_db.DB.Table("articles").
		Select("articles.id, articles.title, articles.description, users.name").
		Joins("left join users on users.id = articles.user_id").
		Where("articles.user_id = ? AND users.id = ?", userId, userId).
		Find(&joinArticle)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joinArticle)
}
