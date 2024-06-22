package response

import (
	"tanahore/internal/model/domain"
	"tanahore/internal/model/web"
)

func AllArticleDomainToResponse(resDomain []domain.Articles) []web.GetAllArticlesResponse {
	var res []web.GetAllArticlesResponse
	for i := range resDomain {
		articleResponse := web.GetAllArticlesResponse{
			ArticleID: resDomain[i].ArticleID,
			Title:     resDomain[i].Title,
			ImageURL:  resDomain[i].ImageURL,
		}
		res = append(res, articleResponse)
	}
	return res
}

func ToInformationRetrieval(resDomain []domain.Articles) []web.AllArticlesResponse {
	var res []web.AllArticlesResponse
	for i := range resDomain {
		articleResponse := web.AllArticlesResponse{
			ArticleID: resDomain[i].ArticleID,
			SoilType:  resDomain[i].SoilType,
			Title:     resDomain[i].Title,
			Content:   resDomain[i].Content,
			ImageURL:  resDomain[i].ImageURL,
		}
		res = append(res, articleResponse)
	}
	return res
}

func ToInformationRetrievalPayload(articles []web.AllArticlesResponse, query string) *web.ToInformationRetrievalFormat {
	return &web.ToInformationRetrievalFormat{
		Query:    query,
		Articles: articles,
	}
}

func ArticleByIDToResponse(resDomain *domain.Articles) *web.AllArticlesResponse {
	return &web.AllArticlesResponse{
		ArticleID: resDomain.ArticleID,
		Title:     resDomain.Title,
		Content:   resDomain.Content,
		SoilType:  resDomain.SoilType,
		ImageURL:  resDomain.ImageURL,
	}
}
