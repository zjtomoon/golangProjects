package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)

// 流式响应
// 当数据产生的时候发送数据
// 使用分块传输编码（Chunked transfer encoding）的流式 JOSN 响应。

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		for _, l := range locations {
			if err := json.NewEncoder(c.Response()).Encode(l); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
