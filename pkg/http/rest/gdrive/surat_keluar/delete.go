package rest

import (
	"net/http"

	"github.com/alifudin-a/go-gdrive/pkg/domain/helper"
	"github.com/alifudin-a/go-gdrive/pkg/domain/helper/models"

	"github.com/labstack/echo/v4"
)

type delete struct{}

func NewDeleteFileHandler() *delete {
	return &delete{}
}

func (*delete) DeleteFile(c echo.Context) (err error) {
	var resp models.Response
	var ds = helper.DriveService

	fID := c.Param("fID")

	err = helper.DeleteFile(ds, fID)
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "File is not exist!"
		return c.JSON(http.StatusOK, resp)
	}

	resp.Code = http.StatusOK
	resp.Message = "File successfully deleted!"

	return c.JSON(http.StatusOK, resp)
}
