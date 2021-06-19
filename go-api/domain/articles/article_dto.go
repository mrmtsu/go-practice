package articles

import "github.com/mrmtsu/go_api/domain/users"

type Article struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	User        users.User `json:"-" gorm:"foreignKey:UserId"`
	UserId      int        `json:"user_id"`
}

type JoinArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"user_name"`
}
