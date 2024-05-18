package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (soilPredictHandler SoilPredictHandlerImpl) GetPrediction(ctx echo.Context) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return responses.StatusBadRequest(ctx, "image file required : ", err)
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to open image file"})
	}
	defer src.Close()

	// Membuat buffer untuk menyimpan file gambar
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Membuat form field untuk file gambar
	part, err := writer.CreateFormFile("image", file.Filename)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to create form file", err)
	}

	// Menyalin konten file ke bagian multipart
	if _, err := io.Copy(part, src); err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to copy image content", err)
	}

	// Menutup writer multipart untuk menulis boundary ke buffer
	if err := writer.Close(); err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to close writer", err)
	}

	// Membuat request ke API eksternal
	externalURL := "http://localhost:8081/api/predict"
	req, err := http.NewRequest("POST", externalURL, &buf)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to create request", err)
	}

	// Menetapkan header Content-Type untuk multipart
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Mengirim request ke API eksternal
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to send request to model API", err)
	}
	defer resp.Body.Close()

	// Membaca response dari API eksternal
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.StatusInternalServerError(ctx, "Failed to read response", err)
	}

	// var soil web.SoilPredictPlansRequest
	// soilType := string(body)

	// switch soilType {
	// case string(domain.Aluvial), string(domain.Latosol), string(domain.Humus), string(domain.Andosol):
	// 	soil.SoilType = domain.SoilTypes(soilType)
	// default:
	// 	return responses.StatusBadRequest(ctx, "Invalid soil type in response: "+soilType, err)
	// }

	// res, err := soilPredictHandler.SoilPredictService.GetSoilPlantsBySoilType(&soil)
	// if err != nil {
	// 	return responses.StatusInternalServerError(ctx, "something went wrong", err)
	// }
	// return responses.StatusOK(ctx, "Successfully predicted soil image", res, nil)
	return responses.StatusOK(ctx, "Successfully predicted soil image", body, nil)
}
