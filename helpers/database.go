package helpers

import (
	"eulabs/common"
	"eulabs/types"
	"fmt"

	"github.com/labstack/echo/v4"
)

const (
	productsTable string = "products"
)

var log = echo.New()

func GetProductDB(id int) (*types.ProductResponse, error) {
	query := "SELECT _id, name, description, price FROM " + productsTable + " WHERE _id = ?"

	var product types.ProductResponse

	err := db.DB.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		log.Logger.Errorf("Error query select product", err)
		return nil, err
	}

	return &product, nil
}

func InsertProduct(product *types.ProductRequest) (*types.ProductResponse, error) {

	query := "INSERT INTO " + productsTable + " (name, description, price) VALUES (?, ?, ?)"

	result, err := db.DB.Exec(query, product.Name, product.Description, product.Price)

	if err != nil {
		log.Logger.Errorf("Error query insert product", err)
		return nil, err
	}

	productID, err := result.LastInsertId()

	insertedProduct := &types.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	if err == nil {
		insertedProduct.Id = int(productID)
	} else {
		log.Logger.Errorf("Error get insert product", err)
	}

	return insertedProduct, nil
}


func UpdateProductDB(id int, product *types.ProductRequest) (*types.ProductResponse, error) {

	query := "UPDATE " + productsTable + " SET name = ?, description = ?, price = ? WHERE _id = ?"

    result, err := db.DB.Exec(query, product.Name, product.Description, product.Price, id)
    if err != nil {
        return nil, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }
    if rowsAffected == 0 {
		msg := "dont update product"
		log.Logger.Errorf(msg)
        return nil, fmt.Errorf(msg)
    }

    updatedProduct := &types.ProductResponse{
        Id:          id,
        Name:        product.Name,
        Description: product.Description,
        Price:       product.Price,
    }

    return updatedProduct, nil
}

func DeleteProductDB(id int) error {
    query := "DELETE FROM " + productsTable + " WHERE _id = ?"

    result, err := db.DB.Exec(query, id)
    if err != nil {
		log.Logger.Error(err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
		log.Logger.Error(err)
        return err
    }
    if rowsAffected == 0 {
		msg := "404"
		log.Logger.Errorf(msg)
        return fmt.Errorf(msg)
    }

    return nil
}

func GetAllProducts() ([]*types.ProductResponse, error) {
    query := "SELECT _id, name, description, price FROM " + productsTable

    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*types.ProductResponse
    for rows.Next() {
        var product types.ProductResponse
        if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price); err != nil {
            return nil, err
        }
        products = append(products, &product)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}
