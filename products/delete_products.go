package products

import (
	"database/sql"
	"net/http"
	"strconv"

	"eulabs/helpers"

	"github.com/labstack/echo/v4"
)

func DeleteProduct() echo.HandlerFunc {
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

        if err := helpers.DeleteProductDB(id); err != nil {
            if err.Error() == "404" {
                return echo.NewHTTPError(http.StatusNotFound, "product not found")
            }
            return echo.NewHTTPError(http.StatusInternalServerError, "error delete product")
        }

        return c.NoContent(http.StatusNoContent)
        
    }
}
