package controller

import (
	"http-db-exp/model"
	"net/http"

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
