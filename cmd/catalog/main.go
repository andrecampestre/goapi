package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/devfullcycle/imersao17/goapi/internal/database" // Import the database package
	"github.com/devfullcycle/imersao17/goapi/internal/service"
	"github.com/devfullcycle/imersao17/goapi/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/categories/{id}", WebCategoryHandler.GetCategory)
	c.Get("/categories", WebCategoryHandler.GetCategories)
	c.Post("/categories", WebCategoryHandler.CreateCategory)

	c.Get("/products/{id}", WebProductHandler.GetProduct)
	c.Get("/products", WebProductHandler.GetProducts)
	c.Get("/products/category/{categoryID}", WebProductHandler.GetProductByCategoryID)
	c.Post("/products", WebProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
