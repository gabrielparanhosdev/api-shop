package helpers

import (
	"eulabs/common"
	"eulabs/types"

	"github.com/labstack/echo/v4"
)

func GetProductDB(id int) (*types.ProductResponse, error) {
	log := echo.New()
	query := "SELECT _id, name, description, price FROM products WHERE _id = ?"

	var product types.ProductResponse
	
	err := db.DB.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
    if err != nil {
		log.Logger.Errorf("Error query select product", err)
        return nil, err
    }

	return &product, nil
}