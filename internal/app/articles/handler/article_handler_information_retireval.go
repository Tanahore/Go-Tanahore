package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"tanahore/internal/model/web"
	converter "tanahore/internal/pkg/converter/response"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler *ArticleHandlerImpl) InformationRetrieval(ctx echo.Context) error {
	query := ctx.Param("query")
	if query == "" {
		return responses.StatusBadRequest(ctx, "query cannot be empty", errors.New("query not found"))
	}

	res, err := handler.Service.InformationRetrieval() // Assuming this is a method in your service
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return responses.StatusNotFound(ctx, "articles not found", err)
		}
		if strings.Contains(err.Error(), "went wrong") {
			return responses.StatusInternalServerError(ctx, "Something went wrong", err)
		}
		return err
	}
	convertedRes := converter.ToInformationRetrievalPayload(res, query)
	for _, article := range convertedRes.Articles {
		fmt.Println("article", article)
	}
	payload, err := json.Marshal(convertedRes)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	url := handler.ModelURL.InformationRetrieval
	fmt.Println("url : ", url)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "Failed to read response", err)
	}

	var apiResponse web.FromInformationRetrieval
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "Failed to unmarshal JSON response", err)
	}

	if len(apiResponse.Articles) == 0 {
		return responses.StatusNotFound(ctx, "cannot find "+query+" in articles", errors.New("articles not found"))
	}
	return responses.StatusOK(ctx, "Success", apiResponse, nil)
}
