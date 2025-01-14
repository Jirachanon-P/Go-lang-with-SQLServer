package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @summary Get All Users
// @description Get All Users
// @tags User
// @id GetAllUsers
// @accept json
// @produce json
// @Router /user [get]
// @security ApiKeyAuth
func GetAllUsers(c *gin.Context) {
	var reqUser Models.User
	err := Models.GetUserById(&reqUser, GetUserFromToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Cant Find Request User",
			Data:   err.Error(),
		})
	}

	if reqUser.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, Models.Response{
			Code:   403,
			Status: "No Access To this Api",
			Data:   nil,
		})
		return
	}

	var users []Models.User
	err = Models.GetAllUsers(&users)
	if err != nil {
		c.JSON(http.StatusNotFound, Models.Response{
			Code:   404,
			Status: "No User Found",
			Data:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Models.Response{
			Code:   200,
			Status: "Success",
			Data:   users,
		})
	}
}

// CreateUser godoc
// @summary Create User
// @description Create User
// @tags User
// @id CreateUser
// @accept json
// @produce json
// @param User body Models.User true "User data to create"
// @Router /user [post]
// @security ApiKeyAuth
func CreateUser(c *gin.Context) {
	var reqUser Models.User
	err := Models.GetUserById(&reqUser, GetUserFromToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Cant Find Request User",
			Data:   err.Error(),
		})
	}

	if reqUser.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, Models.Response{
			Code:   403,
			Status: "No Access To this Api",
			Data:   nil,
		})
		return
	}

	var user Models.User
	c.BindJSON(&user)
	fmt.Println(user)
	err = Models.CreatUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, Models.Response{
			Code:   201,
			Status: ("User Created"),
			Data:   user,
		})
	}
}

// GetUserById godoc
// @summary Get UserBy Id
// @description Get UserBy Id
// @tags User
// @id GetUserById
// @accept json
// @produce json
// @param id path string true "id of user"
// @Router /user/{id} [get]
// @security ApiKeyAuth
func GetUserById(c *gin.Context) {
	var reqUser Models.User
	err := Models.GetUserById(&reqUser, GetUserFromToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Cant Find Request User",
			Data:   err.Error(),
		})
	}

	var user Models.User
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Code:   400,
			Status: "Invalid Id",
			Data:   err.Error(),
		})
		return
	}

	if reqUser.Id != id && reqUser.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, Models.Response{
			Code:   403,
			Status: "No Acces to other user information",
			Data:   nil,
		})
		return
	}

	err = Models.GetUserById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Models.Response{
			Code:   404,
			Status: "No User Found",
			Data:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Models.Response{
			Code:   200,
			Status: "Success",
			Data:   user,
		})
	}
}

// UpdateUser godoc
// @summary Update User
// @description Update User
// @tags User
// @id UpdateUser
// @accept json
// @produce json
// @param id path string true "id of user to be update"
// @param User body Models.User true "User data to update"
// @Router /user/{id} [put]
// @security ApiKeyAuth
func UpdateUser(c *gin.Context) {
	var reqUser Models.User
	err := Models.GetUserById(&reqUser, GetUserFromToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Cant Find Request User",
			Data:   err.Error(),
		})
	}
	var user Models.User
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Code:   400,
			Status: "Invalid Id",
			Data:   err.Error(),
		})
	}

	if reqUser.Id != id && reqUser.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, Models.Response{
			Code:   403,
			Status: "No Acces to other user information",
			Data:   nil,
		})
		return
	}

	err = Models.GetUserById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Models.Response{
			Code:   404,
			Status: "No User Found",
			Data:   err.Error(),
		})
	}

	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Models.Response{
			Code:   200,
			Status: "Updated user " + user.Username,
			Data:   user,
		})
	}
}

// DeleteUser godoc
// @summary Delete User
// @description Delete User
// @tags User
// @id DeleteUser
// @accept json
// @produce json
// @param id path string true "id of user to be delete"
// @Router /user/{id} [delete]
// @security ApiKeyAuth
func DeleteUser(c *gin.Context) {
	var reqUser Models.User
	err := Models.GetUserById(&reqUser, GetUserFromToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Cant Find Request User",
			Data:   err.Error(),
		})
	}
	var user Models.User
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Code:   400,
			Status: "Invalid Id",
			Data:   err.Error(),
		})
	}

	if reqUser.Id != id && reqUser.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, Models.Response{
			Code:   403,
			Status: "No Acces to other user information",
			Data:   nil,
		})
		return
	}

	err = Models.GetUserById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Models.Response{
			Code:   404,
			Status: "No User Found",
			Data:   err.Error(),
		})
	}

	err = Models.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Models.Response{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Models.Response{
			Code:   200,
			Status: "user " + user.Username + " is deleted",
			Data:   user,
		})
	}
}
