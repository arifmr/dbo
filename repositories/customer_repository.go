package repositories

import (
	"github.com/arifmr/dbo/database"
	"github.com/arifmr/dbo/models"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: database.DB,
	}
}

func (r *CustomerRepository) GetCustomers(offset, limit int) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Offset(offset).Limit(limit).Select("id, name, email, address, created_at, updated_at").Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Select("id, name, email, address, created_at, updated_at").First(&customer, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	err := r.db.Create(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) UpdateCustomer(customer *models.Customer) error {
	err := r.db.Save(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) DeleteCustomer(id int) error {
	err := r.db.Delete(&models.Customer{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) SearchCustomers(criteria string) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Where("name LIKE ? OR email LIKE ?", "%"+criteria+"%", "%"+criteria+"%").
		Select("id, name, email, address, created_at, updated_at").
		Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) GetCustomerByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}
