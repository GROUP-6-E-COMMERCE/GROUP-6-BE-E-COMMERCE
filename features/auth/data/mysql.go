package data

import (
	"construct-week1/features/auth"
	"construct-week1/middlewares"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// SelectLogin implements auth.Data
func (repo *mysqlUserRepository) SelectLogin(data auth.User) (interface{}, error) {
	var userData auth.User

	// Error disini
	res := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	if res != nil {
		return nil, errors.New("failed")
	}

	result := repo.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&userData)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		return nil, fmt.Errorf("error")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return nil, errToken
	}

	return token, nil
}

// Dependency Injection
func NewUserRepository(conn *gorm.DB) auth.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}
