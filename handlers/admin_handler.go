package handlers

import (
	"net/http"

	"github.com/Rickykn/buddyku-app.git/dtos"
	"github.com/Rickykn/buddyku-app.git/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterAdmin(c *gin.Context) {
	var registerAdminInput *dtos.AdminRegisterDTO

	err := c.ShouldBindJSON(&registerAdminInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	admin, response, _ := h.adminService.RegisterAdmin(registerAdminInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		newAdmin := &dtos.AdminRegisterResponse{
			Name: admin.Name,
			Role: admin.Role,
		}

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        newAdmin,
		})
	}

}

func (h *Handler) LoginAdmin(c *gin.Context) {
	var loginAdminInput *dtos.LoginAdminDTO

	err := c.ShouldBindJSON(&loginAdminInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	_, response, _ := h.adminService.LoginAdmin(loginAdminInput)

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

func (h *Handler) SetPointReward(c *gin.Context) {
	var pointInput *dtos.RequestPoint

	err := c.ShouldBindJSON(&pointInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	adminContext := c.MustGet("admin").(models.Admin)

	if adminContext.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You're not admin you cant create or change reward point",
		})
		return
	}

	newPoint, response, _ := h.adminService.SetPoint(pointInput)

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
			"data":        newPoint,
		})
	}
}

func (h *Handler) GetAllUser(c *gin.Context) {
	var userResponse []*dtos.ResponseAllUser

	users, response, _ := h.adminService.GetAllUser()

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		for _, val := range users {
			user := &dtos.ResponseAllUser{
				ID:          val.ID,
				Name:        val.Name,
				Email:       val.Email,
				PointReward: val.Point_reward,
			}

			userResponse = append(userResponse, user)
		}
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        userResponse,
		})
	}
}
