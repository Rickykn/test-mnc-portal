package services

import (
	"github.com/Rickykn/buddyku-app.git/dtos"

	help "github.com/Rickykn/buddyku-app.git/helpers"
	"github.com/Rickykn/buddyku-app.git/models"
	r "github.com/Rickykn/buddyku-app.git/repositories"
)

type UserService interface {
	Login(loginInput *dtos.LoginUserDTO) (*dtos.TokenDTO, *help.JsonResponse, error)
	Register(registerInput *dtos.UserRegisterDTO) (*models.User, *help.JsonResponse, error)
	CreateArticle(articleInput *dtos.ArticleInputDTO) (*models.Article, *help.JsonResponse, error)
	GetPointUser(email string) (*help.JsonResponse, error)
	GetArticleDetail(id int) (*help.JsonResponse, error)
}

type userService struct {
	userRepository r.UserRepository
}

type USConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(c *USConfig) UserService {

	return &userService{
		userRepository: c.UserRepository,
	}
}

func (u *userService) Register(registerInput *dtos.UserRegisterDTO) (*models.User, *help.JsonResponse, error) {

	_, row, err := u.userRepository.FindOneUser(registerInput.Email)

	if row == 1 {
		return nil, help.HandlerError(400, "Email has been Taken", nil), err
	}

	hasingPassword, _ := help.HashPassword(registerInput.Password)

	newUser, err := u.userRepository.CreateUser(registerInput.Email, registerInput.Name, hasingPassword)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newUser, help.HandlerSuccess(201, "Success register account", newUser), nil
}

func (u *userService) Login(loginInput *dtos.LoginUserDTO) (*dtos.TokenDTO, *help.JsonResponse, error) {
	findUser, row, err := u.userRepository.FindOneUser(loginInput.Email)
	if row == 0 || err != nil {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	isPasswordCorrect := help.CheckPasswordHash(loginInput.Password, findUser.Password)

	if !isPasswordCorrect {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	dataToken := &dtos.ResponseTokenDTO{
		ID:    findUser.ID,
		Name:  findUser.Name,
		Email: findUser.Email,
	}

	tokenString, err := help.CreateJwtUser(dataToken)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return &dtos.TokenDTO{IDToken: tokenString}, help.HandlerSuccess(200, "Login success", &dtos.TokenDTO{IDToken: tokenString}), nil
}

func (u *userService) CreateArticle(articleInput *dtos.ArticleInputDTO) (*models.Article, *help.JsonResponse, error) {
	//find point reward
	point, row, err := u.userRepository.FindPointReward()

	if row == 0 || err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	findUser, row, err := u.userRepository.FindOneUser(articleInput.Email)
	if row == 0 || err != nil {
		return nil, help.HandlerError(404, "Wrong email or password", nil), err
	}

	//create new post
	newArticle, err := u.userRepository.CreateNewArticle(articleInput, findUser.ID)

	if err != nil {
		return nil, help.HandlerError(404, "Failed Crete new Aricle", nil), err
	}

	newPoint := findUser.Point_reward + point.Value_point

	rowUpdated, err := u.userRepository.UpdatePointUser(findUser.ID, newPoint)

	if rowUpdated == 0 || err != nil {
		return nil, help.HandlerError(404, "Server Error", nil), err
	}

	//return

	return newArticle, help.HandlerSuccess(201, "Success Create New Post", newArticle), nil
}

func (u *userService) GetPointUser(email string) (*help.JsonResponse, error) {

	findUser, row, err := u.userRepository.FindOneUser(email)
	if row == 0 || err != nil {
		return help.HandlerError(404, "User not found", nil), err
	}

	pointResponse := &dtos.UserGetPointDTO{
		Point_reward: findUser.Point_reward,
	}

	return help.HandlerSuccess(200, "Success get point user", pointResponse), err

}

func (u *userService) GetArticleDetail(id int) (*help.JsonResponse, error) {
	article, err := u.userRepository.GetDetailArticle(id)
	if err != nil {
		return help.HandlerError(500, "Article Not found", nil), err
	}

	return help.HandlerSuccess(200, "Success Get Article Detail", article), err
}
