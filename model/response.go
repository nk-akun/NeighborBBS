package model

// APIResponse ...
type APIResponse struct {
	Code    int         `json:"code"`
	Value   interface{} `json:"value"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}

// SysConfigResponse ...
type SysConfigResponse struct {
	SiteTitle       string       `json:"siteTitle"`       // 网站标题
	SiteDescription string       `json:"siteDescription"` // 网站描述
	SiteKeywords    []string     `json:"siteKeywords"`    // 网站关键词
	SiteNavs        []ActionLink `json:"siteNavs"`        // 网站导航
	TokenExpireDays int          `json:"tokenExpireDays"` // token有效天数
}

// ActionLink ...
type ActionLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
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
	AuthorID     int64  `json:"user_id"`
	AuthorName   string `json:"nick_name"`
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
	AuthorName string `json:"user_name"`
	AuthorID   int64  `json:"user_id"`
	AvatarURL  string `json:"avatar_url"`
	Content    string `json:"content"`
	ParentID   int64  `json:"parent_id"`
	LikeCount  int    `json:"like_count"`
	CreateTime int64  `json:"create_time"`
}
