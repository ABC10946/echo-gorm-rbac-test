package main

import (
	"fmt"
	"http-db-exp/controller"
	"http-db-exp/model"
	"log/slog"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, World!")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	errMigrate := db.AutoMigrate(&model.Item{}, &model.User{})

	if errMigrate != nil {
		slog.Error(errMigrate.Error())
		return
	}

	con := controller.Controller{DB: db, Items: []model.Item{}}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/items", con.ItemPost)
	e.GET("/items", con.ItemList)
	e.GET("/itemsWithoutDB", con.ItemListWithoutDB)
	e.GET("/startUpdateItemVariable", con.ItemUpdateVariable)

	e.Use(middleware.Logger())
	e.Logger.Debug(e.Start(":1234"))
}
