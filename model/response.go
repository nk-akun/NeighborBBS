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
	Cursor      int64               `json:"cursor"`
	ArticleList []*ArticleBriefInfo `json:"article_list"`
}

// ArticleBriefInfo ...
type ArticleBriefInfo struct {
	ArticleID    int64          `json:"article_id"`
	Title        string         `json:"title"`
	User         *UserBriefInfo `json:"user"`
	LikeCount    int            `json:"like_count"`
	CommentCount int            `json:"commnet_count"`
	ViewCount    int            `json:"view_count"`
	Liked        bool           `json:"liked"`
	CreateTime   int64          `json:"create_time"`
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
	AuthorNickName string       `json:"user_nickname"`
	AuthorUserName string       `json:"user_username"`
	AuthorID       int64        `json:"user_id"`
	AvatarURL      string       `json:"avatar_url"`
	Content        string       `json:"content"`
	ParentComment  *CommentInfo `json:"parent_comment"`
	LikeCount      int          `json:"like_count"`
	CreateTime     int64        `json:"create_time"`
}

// ResponseValue ...
type ResponseValue struct {
	Value map[string]interface{}
}

// NewResponseValue ...
func NewResponseValue() *ResponseValue {
	return &ResponseValue{
		Value: make(map[string]interface{}),
	}
}

// Set set data into ResponseValue
func (r *ResponseValue) Set(name string, data interface{}) *ResponseValue {
	r.Value[name] = data
	return r
}

// UserBriefInfo ...
type UserBriefInfo struct {
	ID                    int64  `json:"id"`
	Username              string `json:"username"`
	Nickname              string `json:"nickname"`
	AvatarURL             string `json:"avatar_url"`
	Gender                string `json:"gender"`
	Description           string `json:"description"`
	AttentionCount        int    `json:"attention_count"`
	FavouriteArticleCount int    `json:"favourite_article_count"`
	FansCount             int    `json:"fans_count"`
	PostCount             int    `json:"post_count"`
	CommentCount          int    `json:"comment_count"`
	Type                  int    `json:"type"`
	City                  string `json:"city"`
	Province              string `json:"province"`
	Country               string `json:"country"`
}
