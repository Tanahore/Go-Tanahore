package handler

import (
	"strings"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *ArticleHandlerImpl) GetAllArticles(ctx echo.Context) error {
	res, err := handler.Service.GetAllArticles()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return responses.StatusNotFound(ctx, "articles not found", err)
		}
		if strings.Contains(err.Error(), "went wrong") {
			return responses.StatusInternalServerError(ctx, "Something went wrong", err)
		}
	}
	return responses.StatusOK(ctx, "Success", res, nil)
}
