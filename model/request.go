package model

// APIRequest ...
type APIRequest interface{}

// RegisterRequest ...
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest ...
type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// ArticleRequest ...
type ArticleRequest struct {
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CommentRequest ...
type CommentRequest struct {
	UserID    int64  `json:"user_id"`
	ArticleID int64  `json:"article_id"`
	Content   string `json:"content"`
	ParentID  int64  `json:"parent_id"`
}

// LikeArticleRequest ...
type LikeArticleRequest struct {
	UserID    int64 `json:"user_id"`
	ArticleID int64 `json:"article_id"`
}
