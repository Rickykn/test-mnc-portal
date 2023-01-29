package services

import (
	"github.com/Rickykn/buddyku-app.git/dtos"
	help "github.com/Rickykn/buddyku-app.git/helpers"
	"github.com/Rickykn/buddyku-app.git/models"
	r "github.com/Rickykn/buddyku-app.git/repositories"
)

type AdminService interface {
	SetPoint(pointInput *dtos.RequestPoint) (*models.Point, *help.JsonResponse, error)
	RegisterAdmin(registerAdminInput *dtos.AdminRegisterDTO) (*models.Admin, *help.JsonResponse, error)
	LoginAdmin(loginAdminInput *dtos.LoginAdminDTO) (*dtos.TokenAdminDTO, *help.JsonResponse, error)
	GetAllUser() ([]*models.User, *help.JsonResponse, error)
}

type adminService struct {
	adminRepository r.AdminRepository
}

type ASConfig struct {
	AdminRespository r.AdminRepository
}

func NewAdminService(c *ASConfig) AdminService {
	return &adminService{
		adminRepository: c.AdminRespository,
	}
}

func (a *adminService) RegisterAdmin(registerAdminInput *dtos.AdminRegisterDTO) (*models.Admin, *help.JsonResponse, error) {
	_, row, err := a.adminRepository.FindOneAdmin(registerAdminInput.Name)

	if row == 1 {
		return nil, help.HandlerError(400, "Email has been Taken", nil), err
	}

	hasingPassword, _ := help.HashPassword(registerAdminInput.Password)

	newAdmin, err := a.adminRepository.CreateAdmin(registerAdminInput.Name, hasingPassword)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newAdmin, help.HandlerSuccess(201, "Success register admin", newAdmin), nil
}

func (a *adminService) LoginAdmin(loginAdminInput *dtos.LoginAdminDTO) (*dtos.TokenAdminDTO, *help.JsonResponse, error) {
	findAdmin, row, err := a.adminRepository.FindOneAdmin(loginAdminInput.Name)
	if row == 0 || err != nil {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	isPasswordCorrect := help.CheckPasswordHash(loginAdminInput.Password, findAdmin.Password)

	if !isPasswordCorrect {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	dataToken := &dtos.ResponseTokenAdminDTO{
		ID:   findAdmin.ID,
		Name: findAdmin.Name,
		Role: findAdmin.Role,
	}

	tokenString, err := help.CreateJwtAdmin(dataToken)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return &dtos.TokenAdminDTO{IDToken: tokenString}, help.HandlerSuccess(200, "Login success", &dtos.TokenDTO{IDToken: tokenString}), nil
}

func (a *adminService) SetPoint(pointInput *dtos.RequestPoint) (*models.Point, *help.JsonResponse, error) {

	row, err := a.adminRepository.UpdateStatusReward()

	if err != nil || row == 0 {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	newPoint, err := a.adminRepository.CreatePoint(pointInput.Value_point)
	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newPoint, help.HandlerSuccess(201, "Success Create New Point", newPoint), nil
}
func (a *adminService) GetAllUser() ([]*models.User, *help.JsonResponse, error) {
	users, err := a.adminRepository.FindAllUser()

	if err != nil {
		return nil, help.HandlerError(404, "User Not found", nil), err
	}

	return users, help.HandlerSuccess(200, "Get All User Success", users), nil
}
