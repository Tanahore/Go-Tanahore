package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/request"
)

func (service *ArticleServiceImpl) CreateArticle(req *web.CreateArticle) error {
	if err := service.Validate.Struct(req); err != nil {
		return err
	}
	if exists := service.Repository.CheckExistingArticle(req.Title); exists {
		return errors.New("article already exists")
	}
	article := converter.ArticleRequestToDomain(req)
	if err := service.Repository.CreateArticle(article); err != nil {
		return err
	}
	return nil
}
