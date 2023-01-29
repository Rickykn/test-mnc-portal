package dtos

type ArticleInputDTO struct {
	Content_article string `json:"content_article"`
	Source          string `json:"source"`
	Email           string `json:"email"`
}
