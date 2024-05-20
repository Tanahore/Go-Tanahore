package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (service *ArticleServiceImpl) GetArticlesByID(article_id *web.GetArticlesByID) (*web.AllArticlesResponse, error) {
	if err := service.Validate.Struct(article_id); err != nil {
		return nil, errors.New("validation error")
	}
	resDomain, err := service.Repository.GetArticlesByID(article_id)
	if err != nil {
		return nil, err
	}
	res := converter.ArticleByIDToResponse(resDomain)
	return res, nil
}
