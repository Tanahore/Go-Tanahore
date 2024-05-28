package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"tanahore/internal/model/web"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (handler *DeviceHandlerImpl) PredictBestPlant(ctx echo.Context) error {
	id := ctx.Param("id")
	var soilType web.SoilTypeRequest
	if err := ctx.Bind(&soilType); err != nil {
		return responses.StatusBadRequest(ctx, "error", err)
	}
	req := web.DevicePredictBestPlant{
		DeviceID: id,
		SoilType: soilType.SoilType,
	}
	logrus.Info("req : ", req.SoilType, req.DeviceID)
	logrus.Info("soilType : ", soilType)
	inputData, err := handler.Service.GetDeviceInputData(&req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return responses.StatusBadRequest(ctx, "validation error", err)
		}
		return responses.StatusInternalServerError(ctx, "something happened", err)
	}
	payload, err := json.Marshal(inputData)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "something went wrong", err)
	}
	url := handler.ModelURL.PlantRecommendation
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
	var apiResponse web.PlantModelApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Error:", err)
		return responses.StatusInternalServerError(ctx, "Failed to unmarshal JSON response", err)
	}
	return responses.StatusOK(ctx, "successfully predicted plant data", apiResponse, nil)
}
