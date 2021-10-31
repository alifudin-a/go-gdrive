package router

import (
	"os"

	sk "github.com/alifudin-a/go-gdrive/pkg/http/rest/gdrive/surat_keluar"
	sm "github.com/alifudin-a/go-gdrive/pkg/http/rest/gdrive/surat_masuk"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, URI=${uri}, STATUS=${status}, " +
			"HOST=${host}, ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))

	api := e.Group("/api")
	gd := api.Group("/gdrive")

	// surat masuk
	gd.GET("/surat_masuk", sm.NewListFileSMHandler().ListFileSMHandler)
	gd.POST("/surat_masuk", sm.NewCreateFileHandler().CreateFileHandler)
	gd.GET("/surat_masuk/:fID", sm.NewReadFileHandler().ReadFileHandler)
	gd.DELETE("/surat_masuk/:fID", sm.NewDeleteFileHandler().DeleteFile)
	// surat
	gd.GET("/surat_keluar", sk.NewListFileSKHandler().ListFileSKHandler)
	gd.POST("/surat_keluar", sk.NewCreateFileHandler().CreateFileHandler)
	gd.GET("/surat_keluar/:fID", sk.NewReadFileHandler().ReadFileHandler)
	gd.DELETE("/surat_keluar/:fID", sk.NewDeleteFileHandler().DeleteFile)

	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))

	return e
}
