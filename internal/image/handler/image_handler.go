package handler

import (
	"ps-halo-suster/internal/image/dto"
	"ps-halo-suster/internal/image/service"
	"ps-halo-suster/pkg/helper"
	"ps-halo-suster/pkg/httphelper/response"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageService service.ImageService
}

func NewImageHandler(imageService service.ImageService) *ImageHandler {
	return &ImageHandler{imageService: imageService}
}

func (h *ImageHandler) UploadImage(ctx echo.Context) *response.WebResponse {
	fileHeader, err := ctx.FormFile("file")
	helper.Panic400IfError(err)

	file, err := fileHeader.Open()
	helper.Panic400IfError(err)

	fileUrl, err := h.imageService.UploadImage(file, fileHeader)
	helper.PanicIfError(err, "failed to upload image")

	return &response.WebResponse{
		Status:  200,
		Message: "File uploaded successfully",
		Data:    dto.ImageUploadResponse{ImageUrl: fileUrl},
	}
}
