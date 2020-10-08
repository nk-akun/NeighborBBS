CREATE TABLE user_info (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  username VARCHAR(20) UNIQUE NOT NULL,
  nick_name VARCHAR(30) NOT NULL,
  password VARCHAR(30),
  avatar_url VARCHAR(200),
  gender TINYINT NOT NULL,
  email VARCHAR(50),
  email_verified TINYINT NOT NULL DEFAULT 0,
  description varchar(200),
  attention_count INT NOT NULL DEFAULT 0,
  favourite_article_count INT NOT NULL DEFAULT 0,
  fans_count INT NOT NULL DEFAULT 0,
  post_count INT NOT NULL DEFAULT 0,
  comment_count INT NOT NULL DEFAULT 0,
  type INT NOT NULL DEFAULT 0,
  city VARCHAR(40),
  province VARCHAR(40),
  country VARCHAR(40),
  create_time BIGINT,
  update_time BIGINT
);
CREATE TABLE article (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  status INT NOT NULL DEFAULT 0,
  view_count INT NOT NULL DEFAULT 0,
  like_count INT NOT NULL DEFAULT 0,
  comment_count INT NOT NULL DEFAULT 0,
  post_time BIGINT,
  update_time BIGINT
);