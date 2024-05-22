package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (handler SoilPredictHandlerImpl) GetExactPlant(ctx echo.Context) error {
	var request web.SoilPlantRecommendationRequest
	if err := ctx.Bind(&request); err != nil {
		return responses.StatusBadRequest(ctx, "cannot bind data", err)
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	url := "http://localhost:8081/api/recommendation"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "somethine wrong", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
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
	var apiResponse web.PlantModelApiResponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "Failed to unmarshal JSON response", err)
	}

	return responses.StatusOK(ctx, "successfully predicted plant data", apiResponse, nil)
}
