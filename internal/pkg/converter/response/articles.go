package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func AllArticleDomainToResponse(resDomain []domain.Articles) []web.AllArticlesResponse {
	var res []web.AllArticlesResponse
	for i := range resDomain {
		articleResponse := web.AllArticlesResponse{
			ArticleID: resDomain[i].ArticleID,
			Title:     resDomain[i].Title,
			Content:   resDomain[i].Content,
			SoilType:  resDomain[i].SoilType,
		}
		res = append(res, articleResponse)
	}
	return res
}

func ArticleByIDToResponse(resDomain *domain.Articles) *web.AllArticlesResponse {
	return &web.AllArticlesResponse{
		ArticleID: resDomain.ArticleID,
		Title:     resDomain.Title,
		Content:   resDomain.Content,
		SoilType:  resDomain.SoilType,
	}
}
