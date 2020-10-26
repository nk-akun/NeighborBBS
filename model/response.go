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
	TotalNum    int      `json:"total_num"`
	ArticleList []string `json:"article_list"`
}

// ArticleResponse ...
type ArticleResponse struct {
	Title      string `json:"json"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
