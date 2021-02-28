package model

// APIResponse ...
type APIResponse struct {
	Code    int         `json:"code"`    // 返回值code
	Value   interface{} `json:"value"`   // 返回值value
	Success bool        `json:"success"` // 是否执行成功
	Message string      `json:"message"` // 附加信息
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
	Title string `json:"title"` // 导航标题
	URL   string `json:"url"`   // 导航链接
}

// Article

// ArticleListResponse ...
type ArticleListResponse struct {
	TotalNum    int                 `json:"total_num"`    // 文章总数
	Cursor      int64               `json:"cursor"`       // 游标
	ArticleList []*ArticleBriefInfo `json:"article_list"` // 文章简要信息List
}

// ArticleBriefInfo ...
type ArticleBriefInfo struct {
	ArticleID    int64          `json:"article_id"`    // 文章ID
	Title        string         `json:"title"`         // 文章标题
	Summary      string         `json:"summary"`       // 简要信息
	User         *UserBriefInfo `json:"user"`          // 作者的简要信息
	LikeCount    int            `json:"like_count"`    // 点赞数量
	CommentCount int            `json:"comment_count"` // 评论数
	ViewCount    int            `json:"view_count"`    // 浏览数
	Liked        bool           `json:"liked"`         // 当前登录用户是否已点赞(如未登录为false)
	CreateTime   int64          `json:"create_time"`
}

// ArticleResponse ...
type ArticleResponse struct {
	ArticleID    int64          `json:"article_id"`    // 文章ID
	Title        string         `json:"title"`         // 文章标题
	User         *UserBriefInfo `json:"user"`          // 作者的简要信息
	Content      string         `json:"content"`       // 文章内容
	Liked        bool           `json:"liked"`         // 当前登录用户是否已点赞(如未登录为false)
	Favortied    bool           `json:"favorited"`     // 当前登录用户是否已收藏(如未登录为false)
	CommentCount int            `json:"commnet_count"` // 评论数
	LikeCount    int            `json:"like_count"`    // 点赞数
	CreateTime   int64          `json:"create_time"`
}

// Comment

// CommentListResponse ...
type CommentListResponse struct {
	ArticleID   int64          `json:"article_id"`   // 所属的文章ID
	TotalNum    int            `json:"total_num"`    // 评论总数
	Cursor      int64          `json:"cursor"`       // 游标
	CommentList []*CommentInfo `json:"comment_list"` // 评论List
}

// CommentInfo ...
type CommentInfo struct {
	CommentID      int64        `json:"comment_id"`     // 评论ID
	AuthorNickName string       `json:"user_nickname"`  // 作者的昵称
	AuthorUserName string       `json:"user_username"`  // 作者的用户名
	AuthorID       int64        `json:"user_id"`        // 作者ID
	AvatarURL      string       `json:"avatar_url"`     // 作者的头像
	Content        string       `json:"content"`        // 评论内容
	ParentComment  *CommentInfo `json:"parent_comment"` // 父评论信息
	LikeCount      int          `json:"like_count"`     // 点赞数
	CreateTime     int64        `json:"create_time"`
}

// FavoriteResponse ...
type FavoriteResponse struct {
	TotalNum     int                 `json:"total_num"`     // 文章总数
	Cursor       int64               `json:"cursor"`        // 游标
	FavoriteList []*ArticleBriefInfo `json:"favorite_list"` // 文章简要信息List
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

// UserBriefInfo is user's brief information
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
