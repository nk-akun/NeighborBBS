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
	Author       string `json:"author"`
	Content      string `json:"content"`
	CommentCount int    `json:"commnet_count"`
	LikeCount    int    `json:"like_count"`
	CreateTime   int64  `json:"create_time"`
}
