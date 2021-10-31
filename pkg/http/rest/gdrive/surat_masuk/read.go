package rest

import (
	"net/http"

	"github.com/alifudin-a/go-gdrive/pkg/domain/helper"
	"github.com/alifudin-a/go-gdrive/pkg/domain/helper/models"

	"github.com/labstack/echo/v4"
)

type read struct{}

func NewReadFileHandler() *read {
	return &read{}
}

func (*read) ReadFileHandler(c echo.Context) (err error) {
	var resp models.Response
	var ds = helper.DriveService

	fID := c.Param("fID")

	getFile, err := helper.GetFile(ds, fID)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success!"
	resp.Data = map[string]interface{}{
		"file": getFile,
	}

	return c.JSON(http.StatusOK, resp)
}
