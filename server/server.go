package server

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

// createItem godoc
// @Summary Create a new item
// @Description Create a new item with the specified name and price
// @Accept json
// @Produce json
// @Param params body Item true "Item parameters"
// @Success 201 {object} Item
// @Router /items [post]
func createItem(c echo.Context) error {
	var newItem Item
	if err := c.Bind(&newItem); err != nil {
		logrus.WithError(err).Error("Failed to bind request body")
		return err
	}
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	logrus.WithField("item", newItem).Info("Created new item")
	return c.JSON(http.StatusCreated, newItem)
}

// getItems godoc
// @Summary Get all items
// @Description Get a list of all items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func getItems(c echo.Context) error {
	logrus.Info("Getting all items")
	return c.JSON(http.StatusOK, items)
}

// getItem godoc
// @Summary Get an item by ID
// @Description Get an item with the specified ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Router /items/{id} [get]
func getItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse item ID")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for _, item := range items {
		if item.ID == id {
			logrus.WithField("item", item).Info("Found item")
			return c.JSON(http.StatusOK, item)
		}
	}
	logrus.WithField("id", id).Error("Item not found")
	return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

// updateItem godoc
// @Summary Update an item by ID
// @Description Update an item with the specified ID and new name and price
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body Item true "Updated item object"
// @Success 200 {object} Item
// @Router /items/{id} [put]
func updateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse item ID")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var updatedItem Item
	if err := c.Bind(&updatedItem); err != nil {
		logrus.WithError(err).Error("Failed to bind request body")
		return err
	}
	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			logrus.WithField("item", updatedItem).Info("Updated item")
			return c.JSON(http.StatusOK, updatedItem)
		}
	}
	logrus.WithField("id", id).Error("Item not found")
	return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

// deleteItem godoc
// @Summary Delete an item by ID
// @Description Delete an item with the specified ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 204
// @Router /items/{id} [delete]
func deleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse item ID")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			logrus.WithField("id", id).Info("Deleted item")
			break
		}
	}
	return c.JSON(http.StatusOK, items)
}

func Start() error {
	logrus.SetLevel(logrus.DebugLevel)
	e := echo.New()
	e.POST("/items", createItem)
	e.GET("/items", getItems)
	e.GET("/items/:id", getItem)
	e.PUT("/items/:id", updateItem)
	e.DELETE("/items/:id", deleteItem)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Start(":8000")

	return nil
}
