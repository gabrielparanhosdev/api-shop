package products

import (
    "net/http"
	"strconv"
	
    "github.com/labstack/echo/v4"
    // "eulabs/db"
	"eulabs/products/types"
)

func GetProduct() echo.HandlerFunc {
    return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
        }

        product := products.ProductResponse{
			Id:          id,
			Name:        "Produto qualquer",
			Description: "Descrição do Produto qualquer",
			Price:       29.99,
		}

        return c.JSON(http.StatusOK, product)
    }
}
