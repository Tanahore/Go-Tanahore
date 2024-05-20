package request

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func ArticleRequestToDomain(req *web.CreateArticle) *domain.Articles {
	return &domain.Articles{
		Title:    req.Title,
		Content:  req.Content,
		SoilType: req.SoilType,
		// ImageURL: imageURL,
	}
}
