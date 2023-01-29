package handlers

import "github.com/Rickykn/buddyku-app.git/services"

type Handler struct {
	userService  services.UserService
	adminService services.AdminService
}

type HandlerConfig struct {
	UserService  services.UserService
	AdminService services.AdminService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService:  c.UserService,
		adminService: c.AdminService,
	}
}
