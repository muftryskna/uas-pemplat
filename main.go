package main

import (
	"uas-pemplat/database"
	"uas-pemplat/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "uas-pemplat/docs" // Import generated docs
)

// @title UAS Pemrograman Berbasis Platform API
// @description API CRUD dengan menggunakan bahasa GO
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http https

func main() {
	// Inisialisasi koneksi database
	database.InitDB()

	// Membuat router baru
	r := gin.Default()

	// Middleware untuk mendukung CORS
	r.Use(cors.Default())

	// Grouping untuk API Users
	usersGroup := r.Group("/api/users")
	{
		// @Summary Update a user by ID
		// @Description Update a user with JSON payload
		// @Accept json
		// @Produce json
		// @Param id path int true "User ID"
		// @Param user body User true "User object to be updated"
		// @Success 200 {object} User
		// @Router /api/users/{id} [put]
		usersGroup.PUT("/{id}", handlers.UpdateUserHandler)

		// @Summary Delete a user by ID
		// @Description Delete a user by their ID
		// @Produce json
		// @Param id path int true "User ID"
		// @Success 200 {string} string
		// @Router /api/users/{id} [delete]
		usersGroup.DELETE("/{id}", handlers.DeleteUserHandler)

		// @Summary Create a new user
		// @Description Create a new user with JSON payload
		// @Accept json
		// @Produce json
		// @Param user body User true "User object to be created"
		// @Success 201 {object} User
		// @Router /api/users/ [post]
		usersGroup.POST("/", handlers.CreateUserHandler)

		// @Summary Get a user by ID
		// @Description Get a user by their ID
		// @Produce json
		// @Param id path int true "User ID"
		// @Success 200 {object} User
		// @Router /api/users/{id} [get]
		usersGroup.GET("/{id}", handlers.GetUserHandler)
		// ...

		// @Summary Get all users
		// @Description Get a list of all users
		// @Produce json
		// @Success 200 {array} User
		// @Router /api/users/ [get]
		usersGroup.GET("/", handlers.GetAllUsersHandler)
	}

	// Grouping untuk API Products
	productsGroup := r.Group("/api/products")
	{
		// @Summary Get a products by id
		// @Description Get a list of products by its id
		// @Produce json
		// @Success 200 {array} Product
		// @Router /api/products/{id} [get]
		productsGroup.GET("/{id}", handlers.GetProductHandler)
		// @Summary Update a product by ID
		// @Description Update a product with JSON payload
		// @Accept json
		// @Produce json
		// @Param id path int true "Product ID"
		// @Param product body Product true "Product object to be updated"
		// @Success 200 {object} Product
		// @Router /api/products/{id} [put]
		productsGroup.PUT("/{id}", handlers.UpdateProductHandler)

		// @Summary Delete a product by ID
		// @Description Delete a product by its ID
		// @Produce json
		// @Param id path int true "Product ID"
		// @Success 200 {string} string
		// @Router /api/products/{id} [delete]
		productsGroup.DELETE("/{id}", handlers.DeleteProductHandler)

		// @Summary Create a new product
		// @Description Create a new product with JSON payload
		// @Accept json
		// @Produce json
		// @Param product body Product true "Product object to be created"
		// @Success 201 {object} Product
		// @Router /api/products/ [post]
		productsGroup.POST("/", handlers.CreateProductHandler)
		// ...

		// @Summary Get all products
		// @Description Get a list of all products
		// @Produce json
		// @Success 200 {array} Product
		// @Router /api/products/ [get]
		productsGroup.GET("/", handlers.GetAllProductsHandler)
	}

	// @Summary User login
	// @Description Perform user login with JSON payload
	// @Accept json
	// @Produce json
	// @Param user body User true "User login credentials"
	// @Success 200 {object} LoginResponse
	// @Router /api/login [post]
	r.POST("/api/login", handlers.LoginHandler)

	// @Summary Swagger endpoint
	// @Description Swagger UI endpoint
	// @Produce html
	// @Router /swagger/*any [get]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Menjalankan server
	r.Run(":8080")
}
