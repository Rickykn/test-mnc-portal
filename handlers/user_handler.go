package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rickykn/buddyku-app.git/dtos"
	"github.com/Rickykn/buddyku-app.git/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var registerInput *dtos.UserRegisterDTO

	err := c.ShouldBindJSON(&registerInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	user, response, _ := h.userService.Register(registerInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		newUser := &dtos.UserRegisterResponse{
			Name:  user.Name,
			Email: user.Email,
		}

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        newUser,
		})
	}

}

func (h *Handler) LoginUser(c *gin.Context) {
	var loginInput *dtos.LoginUserDTO

	err := c.ShouldBindJSON(&loginInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	_, response, _ := h.userService.Login(loginInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	}

}

func (h *Handler) UserCreateNewPost(c *gin.Context) {
	var articleInput *dtos.ArticleInputDTO

	userContext := c.MustGet("user").(models.User)

	err := c.ShouldBindJSON(&articleInput)

	articleInput.Email = userContext.Email

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	article, response, _ := h.userService.CreateArticle(articleInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        article,
		})
	}

}

func (h *Handler) GetPointUser(c *gin.Context) {

	userContext := c.MustGet("user").(models.User)

	response, _ := h.userService.GetPointUser(userContext.Email)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	}
}

func (h *Handler) GetArticleDetail(c *gin.Context) {
	id := c.Param("id")

	convId, _ := strconv.Atoi(id)

	response, _ := h.userService.GetArticleDetail(convId)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	}
}
