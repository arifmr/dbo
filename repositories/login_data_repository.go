package repositories

import (
	"github.com/arifmr/dbo/database"
	"github.com/arifmr/dbo/models"
	"gorm.io/gorm"
)

type LoginDataRepository struct {
	db *gorm.DB
}

func NewLoginDataRepository() *LoginDataRepository {
	return &LoginDataRepository{
		db: database.DB,
	}
}

func (l *LoginDataRepository) GetLoginDataByCustomerID(customerID int64) (loginData []*models.LoginData, err error) {
	err = database.DB.Where("customer_id = ?", customerID).Order("login_time desc").Find(&loginData).Error
	if err != nil {
		return nil, err
	}
	return loginData, nil
}

func (l *LoginDataRepository) InsertLoginData(loginData *models.LoginData) error {
	err := database.DB.Create(loginData).Error
	if err != nil {
		return err
	}
	return nil
}
