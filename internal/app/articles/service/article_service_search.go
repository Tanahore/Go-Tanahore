package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
)

func (service *ArticleServiceImpl) ArticleSearch(req *web.GetArticlesByTitle) ([]web.AllArticlesResponse, error) {
	if err := service.Validate.Struct(req); err != nil {
		return nil, errors.New("validation error")
	}
	resDomain, err := service.Repository.GetArticlesByTitle(req.Title)
	if err != nil {
		return nil, err
	}
	if len(resDomain) == 0 {
		return nil, errors.New("data not found")
	}
	res := converter.AllArticleDomainToResponse(resDomain)
	return res, nil
}
