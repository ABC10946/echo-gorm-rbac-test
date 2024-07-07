package controller

import (
	"http-db-exp/model"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (con Controller) ItemPost(c echo.Context) error {
	type ItemReceived struct {
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
	}

	var itemReceived ItemReceived
	errBind := c.Bind(&itemReceived)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, errBind.Error())
	}

	item := model.Item{
		Name:     itemReceived.Name,
		Price:    itemReceived.Price,
		Quantity: itemReceived.Quantity,
	}

	result := con.DB.Create(&item)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusCreated, item)
}

func (con Controller) ItemList(c echo.Context) error {
	var items []model.Item
	result := con.DB.Find(&items)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func (con *Controller) ItemListWithoutDB(c echo.Context) error {
	con.Mutex.Lock()
	slog.Debug("items", con.Items)
	con.Mutex.Unlock()
	return c.JSON(http.StatusOK, con.Items)
}

func (con *Controller) ItemUpdateVariable(c echo.Context) error {
	go func(conn *Controller) {
		time.Sleep(1 * time.Second)
		slog.Debug("Start updating item variable")
		for {
			conn.Mutex.Lock()
			var items []model.Item
			result := conn.DB.Find(&items)
			if result.Error != nil {
				slog.Error(result.Error.Error())
				c.Set("error", result.Error.Error())
				return
			}
			slog.Debug("items", items)
			conn.Items = items
			slog.Debug("Finish updating item variable")
			time.Sleep(1 * time.Second)
			conn.Mutex.Unlock()
		}
	}(con)
	return c.JSON(http.StatusOK, "Start updating item variable")
}
