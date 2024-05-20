package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (service *ArticleServiceImpl) GetAllArticles() ([]web.AllArticlesResponse, error) {
	resDomain, err := service.Repository.GetAllArticles()
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	if resDomain == nil {
		return nil, errors.New("not found")
	}
	res := converter.AllArticleDomainToResponse(resDomain)
	return res, nil
}
