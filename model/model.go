package model

// User stores user infomation
type User struct {
	ID                    int64  `gorm:"column:id" json:"id"`
	UserName              string `gorm:"column:username" json:"username"`
	NickName              string `gorm:"column:nick_name" json:"nick_name"`
	Password              string `gorm:"column:password" json:"password"`
	AvatarURL             string `gorm:"column:avatar_url" json:"avatar_url"`
	Gender                string `gorm:"column:gender" json:"gender"`
	Email                 string `gorm:"column:email" json:"email"`
	EmailVerified         bool   `gorm:"column:email_verified" json:"email_verified"`
	Description           string `gorm:"column:description" json:"description"`
	AttentionCount        int    `gorm:"column:attention_count" json:"attention_count"`
	FavouriteArticleCount int    `gorm:"column:favourite_article_count" json:"favourite_article_count"`
	FansCount             int    `gorm:"column:fans_count" json:"fans_count"`
	PostCount             int    `gorm:"column:post_count" json:"post_count"`
	CommentCount          int    `gorm:"column:comment_count" json:"comment_count"`
	Type                  int    `gorm:"column:type" json:"type"`
	City                  string `gorm:"column:city" json:"city"`
	Province              string `gorm:"column:province" json:"province"`
	Country               string `gorm:"column:country" json:"country"`
	CreateTime            int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime            int64  `gorm:"column:update_time" json:"update_time"`
}

// Article stores article infomation
type Article struct {
	ID     int64  `gorm:"column:id" json:"id"`
	UserID int64  `gorm:"column:user_id" json:"user_id"`
	Title  string `gorm:"column:title" json:"title"`
}
