package repository

import (
	"gorm.io/gorm"
	"udemy-golang-react/model"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// フィールド定義
type userRepository struct {
	db *gorm.DB
}

// IUserRepositoryを宣言している場合はGetUserByEmail , CreateUserの二つのメソッドを実装する必要がある。

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// GetUserByEmail
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	//ur *userRepository →　const ur = userRepository的な意味　userRepositoryには
	//const db = gorm.DBのフィールドがあり、それを呼び出している
	//First　→ gormのメソッド
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	//新規作成だからとりあえず何も返さない
	return nil
}
