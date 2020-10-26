package model

// Models stores the models which will be create as tables in the mysql
var Models = []interface{}{&User{}, &Article{}}

// Model ...
type Model struct {
	ID int64 `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
}

// User stores user infomation
type User struct {
	Model
	Username              string `gorm:"column:username;type:varchar(20);unique;not null" json:"username"`
	Nickname              string `gorm:"column:nickname;type:varchar(30)" json:"nick_name"`
	Password              string `gorm:"column:password;type:varchar(30);not null" json:"password"`
	AvatarURL             string `gorm:"column:avatar_url;type:varchar(200)" json:"avatar_url"`
	Gender                string `gorm:"column:gender;type:tinyint;default:2" json:"gender"`
	Email                 string `gorm:"column:email;type:varchar(50)" json:"email"`
	EmailVerified         bool   `gorm:"column:email_verified;type:tinyint;default 0" json:"email_verified"`
	Description           string `gorm:"column:description;type:varchar(200)" json:"description"`
	AttentionCount        int    `gorm:"column:attention_count;type:int;default:0" json:"attention_count"`
	FavouriteArticleCount int    `gorm:"column:favourite_article_count;type:int;default:0" json:"favourite_article_count"`
	FansCount             int    `gorm:"column:fans_count;type:int;default:0" json:"fans_count"`
	PostCount             int    `gorm:"column:post_count;type:int;default:0" json:"post_count"`
	CommentCount          int    `gorm:"column:comment_count;type:int;default:0" json:"comment_count"`
	Type                  int    `gorm:"column:type;type:int;default:0" json:"type"`
	City                  string `gorm:"column:city;type:varchar(50)" json:"city"`
	Province              string `gorm:"column:province;type:varchar(50)" json:"province"`
	Country               string `gorm:"column:country;type:varchar(50)" json:"country"`
	CreateTime            int64  `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime            int64  `gorm:"column:update_time;default:null" json:"update_time"`
	DeleteTime            int64  `gorm:"column:delete_time;default:null" json:"delete_time"`
}

// Article stores article infomation
type Article struct {
	Model
	UserID       int64  `gorm:"column:user_id;type:int" json:"user_id"`
	Title        string `gorm:"column:title;type:varchar(50);not null" json:"title"`
	Status       int    `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`
	Content      string `gorm:"column:content;type:text" json:"content"`
	ViewCount    int64  `gorm:"column:view_count;type:int;default:0" json:"view_count"`
	CommentCount int64  `gorm:"column:comment_count;type:int;default:0" json:"comment_count"`
	LikeCount    int64  `gorm:"column:like_count;type:int;default:0" json:"like_count"`
	CreateTime   int64  `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime   int64  `gorm:"column:update_time;default:null" json:"update_time"`
	DeleteTime   int64  `gorm:"column:delete_time;default:null" json:"delete_time"`
}
