package ui

import "github.com/gin-gonic/gin"

type StockController struct {
}

func NewStockController() StockController {
	return StockController{}
}

func (c StockController) GetStock(ctx *gin.Context) {

}
