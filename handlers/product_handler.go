// handlers/product_handler.go
package handlers

import (
	"fmt"
	"net/http"
	"uas-pemplat/database"
	"uas-pemplat/models"

	"github.com/gin-gonic/gin"
)

// CreateProductHandler handles the creation of a new product
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Accept json
// @Produce json
// @Param product body models.Product true "Product object to be created"
// @Success 201 {object} models.Product
// @Router /api/products/ [post]
func CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan produk ke database
	if err := database.DB.Create(&product).Error; err != nil {
		fmt.Println("Error creating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Kirim respons berhasil dengan informasi produk yang baru saja dibuat
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

// GetAllProductsHandler handles the retrieval of all products
// @Summary Get all products
// @Description Get a list of all products
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products/ [get]
func GetAllProductsHandler(c *gin.Context) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		fmt.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProductHandler handles the update of a product by ID
// @Summary Update a product by ID
// @Description Update a product with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product object to be updated"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [put]
func UpdateProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		fmt.Println("Error finding product:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan produk ke database
	if err := database.DB.Save(&product).Error; err != nil {
		fmt.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductHandler handles the retrieval of a product by ID
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [get]
func GetProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		fmt.Println("Error finding product:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProductHandler handles the deletion of a product by ID
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} string
// @Router /api/products/{id} [delete]
func DeleteProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		fmt.Println("Error finding product:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Hapus produk dari database
	if err := database.DB.Delete(&product).Error; err != nil {
		fmt.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
