package rest

import (
	"net/http"
	"os"

	"github.com/alifudin-a/go-gdrive/pkg/domain/helper"
	"github.com/alifudin-a/go-gdrive/pkg/domain/models"

	"github.com/labstack/echo/v4"
)

type list struct{}

func NewListFileSKHandler() *list {
	return &list{}
}

func (*list) ListFileSKHandler(c echo.Context) (err error) {
	var resp models.Response
	var ds = helper.DriveService

	// file, err := c.FormFile("file")
	// if err != nil {
	// 	return err
	// }

	// src, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()

	fID := os.Getenv("FID_SURAT_KELUAR")

	listFile, err := helper.ListFile(ds, fID)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success!"
	resp.Data = map[string]interface{}{
		"list": listFile,
	}

	return c.JSON(http.StatusOK, resp)
}
