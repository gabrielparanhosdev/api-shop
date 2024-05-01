package products

import (
	"net/http"

	"eulabs/helpers"
	"eulabs/types"

	"github.com/labstack/echo/v4"
)

func CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product types.ProductRequest
		if err := c.Bind(&product); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error payload")
		}

		insertedProduct, err := helpers.InsertProduct(&product)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error insert product")
		}

		return c.JSON(http.StatusCreated, insertedProduct)
	}
}
