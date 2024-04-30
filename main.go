package main

import (
    "github.com/labstack/echo/v4"
    "eulabs/db"
    "eulabs/products"
    
)

func main() {
    e := echo.New()

    if err := db.ConnectDB(); err != nil {
        e.Logger.Fatal(err.Error())
    }

    // routes
    e.GET("/products/:id", products.GetProduct())

    e.Logger.Fatal(e.Start(":4000"))
}
