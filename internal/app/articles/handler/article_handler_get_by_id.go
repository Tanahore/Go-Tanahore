package handler

import (
	"strconv"
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *ArticleHandlerImpl) GetArticleByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}
	inputID := web.GetArticlesByID{
		ArticleID: uint(id),
	}
	res, err := handler.Service.GetArticlesByID(&inputID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return responses.StatusNotFound(ctx, "article not found", err)
		}
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	return responses.StatusOK(ctx, "Success", res, nil)
}
