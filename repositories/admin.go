package repositories

import (
	"github.com/Rickykn/buddyku-app.git/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(name, password string) (*models.Admin, error)
	FindOneAdmin(name string) (*models.Admin, int, error)
	CreatePoint(point int) (*models.Point, error)
	FindAllUser() ([]*models.User, error)
	UpdateStatusReward() (int, error)
}

type adminRepository struct {
	db *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func NewAdminRepository(c *ARConfig) AdminRepository {
	return &adminRepository{
		db: c.DB,
	}
}

func (a *adminRepository) CreateAdmin(name, password string) (*models.Admin, error) {
	newAdmin := &models.Admin{
		Name:     name,
		Password: password,
		Role:     "admin",
	}

	result := a.db.Create(&newAdmin)

	return newAdmin, result.Error
}
func (a *adminRepository) FindOneAdmin(name string) (*models.Admin, int, error) {
	var admin *models.Admin

	result := a.db.Where("name = ?", name).First(&admin)

	return admin, int(result.RowsAffected), result.Error
}

func (a *adminRepository) CreatePoint(point int) (*models.Point, error) {
	newPoint := &models.Point{
		Value_point: point,
		Status:      "Active",
	}

	result := a.db.Create(&newPoint)

	return newPoint, result.Error
}

func (a *adminRepository) FindAllUser() ([]*models.User, error) {
	var users []*models.User

	result := a.db.Find(&users)

	return users, result.Error
}

func (a *adminRepository) UpdateStatusReward() (int, error) {
	result := a.db.Model(&models.Point{}).Where("status = ?", "Active").Update("status", "No Active")
	return int(result.RowsAffected), result.Error
}
