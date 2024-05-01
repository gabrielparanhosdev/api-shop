package products

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"eulabs/helpers"
)

func GetProduct() echo.HandlerFunc {
    return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "invalid Id")
        }

        product, err := helpers.GetProductDB(id)
		if product == nil || err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "Product not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Error select product")
		}
	
		return c.JSON(http.StatusOK, product)
    }
}
