package handler

import (
	"errors"
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *ArticleHandlerImpl) ArticleSearch(ctx echo.Context) error {
	query := ctx.Param("title")
	if query == "" {
		return responses.StatusBadRequest(ctx, "title needed", errors.New("title cannot be empty"))
	}
	input := web.GetArticlesByTitle{
		Title: query,
	}
	res, err := handler.Service.ArticleSearch(&input)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return responses.StatusBadRequest(ctx, "validation error", err)
		}
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	if res == nil {
		return responses.StatusNotFound(ctx, "data not found", err)
	}
	return responses.StatusOK(ctx, "success", res, nil)
}
