package dtos

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenDTO struct {
	IDToken string `json:"idToken"`
}

type ResponseTokenDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRegisterDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserGetPointDTO struct {
	Point_reward int `json:"point_reward"`
}
