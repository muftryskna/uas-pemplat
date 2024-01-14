// handlers/user_handler.go
package handlers

import (
	"net/http"
	"uas-pemplat/database"
	"uas-pemplat/models"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Accept json
// @Produce json
// @Param user body models.User true "User object to be created"
// @Success 201 {object} models.User
// @Router /api/users/ [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// GetUserHandler handles the retrieval of a user by ID
// @Summary Get a user by ID
// @Description Get a user by their ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /api/users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler handles the update of a user by ID
// @Summary Update a user by ID
// @Description Update a user with the provided details
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User object to be updated"
// @Success 200 {object} models.User
// @Router /api/users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler handles the deletion of a user by ID
// @Summary Delete a user by ID
// @Description Delete a user by their ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string
// @Router /api/users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// LoginHandler handles the user login operation
// @Summary User login
// @Description Authenticate user
// @Accept json
// @Produce json
// @Param loginUser body models.User true "User credentials for login"
// @Success 200 {object} models.User
// @Router /api/login [post]
func LoginHandler(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	if err := database.DB.Where("email = ?", loginUser.Email).First(&userFromDB).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// TODO: Implementasi autentikasi login (misalnya, menggunakan bcrypt)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// GetAllUsersHandler handles the retrieval of all users
// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users/ [get]
func GetAllUsersHandler(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}
