package model

type ArticleTag struct {
	*Model
	Name string `json:"name"`
}

func (at ArticleTag) TableName() string {
	return "article_tag"
}
