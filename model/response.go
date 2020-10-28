package model

// APIResponse ...
type APIResponse struct {
	Code    int         `json:"code"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

// Article

// ArticleListResponse ...
type ArticleListResponse struct {
	TotalNum    int                 `json:"total_num"`
	ArticleList []*ArticleBriefInfo `json:"article_list"`
}

// ArticleBriefInfo ...
type ArticleBriefInfo struct {
	ArticleID  int64  `json:"article_id"`
	Title      string `json:"title"`
	CreateTime int64  `json:"create_time"`
}

// ArticleResponse ...
type ArticleResponse struct {
	Title        string `json:"title"`
	AuthorID     int64  `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AvatarURL    string `json:"avatar_url"`
	Content      string `json:"content"`
	CommentCount int    `json:"commnet_count"`
	LikeCount    int    `json:"like_count"`
	CreateTime   int64  `json:"create_time"`
}

// Comment

// CommentListResponse ...
type CommentListResponse struct {
	ArticleID   int64          `json:"article_id"`
	TotalNum    int            `json:"total_num"`
	CommentList []*CommentInfo `json:"comment_list"`
}

// CommentInfo ...
type CommentInfo struct {
	AuthorName string `json:"author_name"`
	AuthorID   string `json:"author_id"`
	AvatarURL  string `json:"avatar_url"`
	Content    string `json:"content"`
	ParentID   int64  `json:"parent_id"`
	LikeCount  int    `json:"like_count"`
	CreateTime int64  `json:"create_time"`
}
