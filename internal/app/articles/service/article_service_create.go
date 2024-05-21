package service

import (
	"errors"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/request"

	"github.com/labstack/echo/v4"
)

func (service *ArticleServiceImpl) CreateArticle(req *web.CreateArticle, ctx echo.Context) error {
	err := service.Validate.Struct(req)
	if err != nil {
		return err
	}
	if articleDomain := service.Repository.CheckExistingArticle(req.Title); articleDomain != nil {
		return errors.New("article exists")
	}
	req.Image, err = service.Cloudinary.Uploader(ctx, "image", "articles", true)
	if err != nil {
		return err
	}
	article := converter.ArticleRequestToDomain(req, req.Image)
	if err := service.Repository.CreateArticle(article); err != nil {
		return err
	}
	return nil
}
