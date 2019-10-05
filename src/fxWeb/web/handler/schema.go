package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *WebHandler) CreateTable(ctx echo.Context) error {
	if err := handler.service.CreateTable(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "create_success")
}

func (handler *WebHandler) DeleteTable(ctx echo.Context) error {
	if err := handler.service.DeleteTable(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "delete_success")
}
