// products/products.go
package products

import (
	"eulabs/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProducts() echo.HandlerFunc {
    return func(c echo.Context) error {
        products, err := helpers.GetAllProductsDB()
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Error get all products")
        }

        return c.JSON(http.StatusOK, products)
    }
}
