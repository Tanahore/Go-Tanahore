package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (service *ArticleServiceImpl) InformationRetrieval() ([]web.AllArticlesResponse, error) {
	resDomain, err := service.Repository.GetAllArticles()
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	if len(resDomain) == 0 {
		return nil, errors.New("not found")
	}
	res := converter.ToInformationRetrieval(resDomain)
	return res, nil
}
