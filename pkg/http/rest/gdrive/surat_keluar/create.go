package rest

import (
	"net/http"
	"os"

	"github.com/alifudin-a/go-gdrive/pkg/domain/helper"
	"github.com/alifudin-a/go-gdrive/pkg/domain/helper/models"

	"github.com/labstack/echo/v4"
)

type create struct{}

func NewCreateFileHandler() *create {
	return &create{}
}

func (*create) CreateFileHandler(c echo.Context) (err error) {
	var resp models.Response
	var ds = helper.DriveService

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fID := os.Getenv("FID_CREATE_SK")

	create, err := helper.CreateFile(ds, file.Filename, "application/octet-stream", src, fID)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success!"
	resp.Data = map[string]interface{}{
		"file": create,
	}

	return c.JSON(http.StatusOK, resp)
}
