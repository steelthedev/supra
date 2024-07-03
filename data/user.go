package data

import (
	"github.com/steelthedev/supra-paints/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	hashedPwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPwd)
	return nil
}

func (us *UserService) CreateUser(user User) error {
	if err := us.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
