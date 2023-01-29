package repositories

import (
	"fmt"

	"github.com/Rickykn/buddyku-app.git/dtos"
	"github.com/Rickykn/buddyku-app.git/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(email, name, password string) (*models.User, error)
	FindOneUser(email string) (*models.User, int, error)
	FindPointReward() (*models.Point, int, error)
	CreateNewArticle(inputArticle *dtos.ArticleInputDTO, user_id int) (*models.Article, error)
	UpdatePointUser(user_id, newPoint int) (int, error)
	GetDetailArticle(id int) (*models.Article, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (u *userRepository) CreateUser(email, name, password string) (*models.User, error) {
	newUser := &models.User{
		Name:         name,
		Email:        email,
		Password:     password,
		Point_reward: 0,
	}

	result := u.db.Create(&newUser)

	return newUser, result.Error
}
func (u *userRepository) FindOneUser(email string) (*models.User, int, error) {
	var user *models.User

	result := u.db.Where("email = ?", email).First(&user)

	return user, int(result.RowsAffected), result.Error
}

func (u *userRepository) FindPointReward() (*models.Point, int, error) {
	var point *models.Point

	result := u.db.Where("status = ?", "Active").First(&point)

	return point, int(result.RowsAffected), result.Error
}

func (u *userRepository) CreateNewArticle(inputArticle *dtos.ArticleInputDTO, user_id int) (*models.Article, error) {
	newArticle := &models.Article{
		Content_article: inputArticle.Content_article,
		Source:          inputArticle.Source,
		User_id:         &user_id,
	}
	result := u.db.Create(&newArticle)

	return newArticle, result.Error
}

func (u *userRepository) UpdatePointUser(user_id, newPoint int) (int, error) {
	result := u.db.Model(&models.User{}).Where("id = ?", user_id).Update("point_reward", newPoint)
	return int(result.RowsAffected), result.Error
}

func (u *userRepository) GetDetailArticle(id int) (*models.Article, error) {
	var articleDetail *models.Article
	fmt.Println(id)

	result := u.db.Where("id = ?", id).Preload("User").First(&articleDetail)

	return articleDetail, result.Error
}
