package handler

import (
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *ArticleHandlerImpl) CreateArticle(ctx echo.Context) error {
	var req web.CreateArticle
	if err := ctx.Bind(&req); err != nil {
		return responses.StatusBadRequest(ctx, "Failed to Bind Data", err)
	}
	if err := handler.Service.CreateArticle(&req); err != nil {
		if strings.Contains(err.Error(), "Validation") {
			return responses.StatusBadRequest(ctx, "Input data is not valid", err)
		}
		if strings.Contains(err.Error(), "exists") {
			return responses.StatusAlreadyExist(ctx, "Article Existed", err)
		}
		return responses.StatusInternalServerError(ctx, "Something went wrong", err)
	}
	return responses.StatusCreated(ctx, "Article Created", nil, nil)
}
