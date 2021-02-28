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

// FavoriteArticleRequest ...
type FavoriteArticleRequest struct {
	UserID    int64 `json:"user_id"`
	ArticleID int64 `json:"article_id"`
}

// UpdateUserProfile ...
type UpdateUserProfile struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

// SetUsernameRequest ...
type SetUsernameRequest struct {
	Username string `json:"username"`
}

// SetEmailRequest ...
type SetEmailRequest struct {
	Email string `json:"email"`
}

// SetPasswordRequest ...
type SetPasswordRequest struct {
	Password string `json:"password"`
}

// UpdatePasswordRequest ...
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}
