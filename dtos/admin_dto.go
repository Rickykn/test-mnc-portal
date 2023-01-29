package dtos

type AdminRegisterDTO struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminRegisterResponse struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type ResponseTokenAdminDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type TokenAdminDTO struct {
	IDToken string `json:"idToken"`
}

type LoginAdminDTO struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseAllUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PointReward int    `json:"point_reward"`
}
