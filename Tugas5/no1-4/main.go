package main

import (
	"github.com/labstack/echo/v4"
	"tugas5/handler"
)

func main() {
	e := echo.New()

	e.GET("/datas", handler.GetDatas)
	e.GET("/datas:Id", handler.GetDataById)
	e.POST("/datas", handler.CreateData)
	e.DELETE("/datas:Id", handler.DeleteData)

	e.Start(":8000")
}
