package products

import (
	"database/sql"
	"net/http"
	"strconv"

	"eulabs/helpers"
	"eulabs/types"

	"github.com/labstack/echo/v4"
)

func UpdateProduct() echo.HandlerFunc {
    return func(c echo.Context) error {
        var productRequest types.ProductRequest
        if err := c.Bind(&productRequest); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "invalid params")
        }

        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "invalid Id")
        }

        product, err := helpers.GetProductDB(id) 

		if product == nil || err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "Product not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Error select product for update")
		}

        updatedProduct, err := helpers.UpdateProductDB(id, &productRequest)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Error update product")
        }
	
		return c.JSON(http.StatusOK, updatedProduct)
    }
}
