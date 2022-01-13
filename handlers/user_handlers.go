package handlers

import (
	"gin_webservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers is the handler of GET /user
// @Summary List  all users
// @Description list all the users
// @Tags user
// @Produce  json
// @Success 200 {object} []models.User
// @Router /users/user/ [get]
func GetAllUsers(c *gin.Context) {
	users, err := models.DB.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// GetUserByID is the handler of GET /user/:id
// @Summary Get single user by given id
// @Description get single user
// @Tags user
// @Produce  json
// @Param id path integer true "id"
// @Success 200 {object} models.User
// @Router /users/user/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := models.DB.GetUserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser is handler for POST /user
// @Summary create user
// @Description create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "user"
// @Success 201 {object} models.User
// @Failure 404 {int} http.StatusNotFound
// @Router /users/user/ [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.DB.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

// UpdateUser is handler for PUT /user
// @Summary update a user
// @Description update an existing user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "user"
// @Param id path integer true "id"
// @Success 200 {object} models.User
// @Failure 404 {int} http.StatusNotFound
// @Router /users/user/ [put]
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := models.DB.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	models.DB.UpdateUser(&user, id)

	c.JSON(http.StatusOK, user)
}

// DeleteUser is handler for DELETE /user
// @Summary delete user
// @Description delete an existing user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path integer true "id"
// @Success 200 {object} models.User
// @Failure 404 {integer} http.StatusNotFound
// @Router /users/user/ [delete]
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	models.DB.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
