package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"udemy-golang-react/model"
)

type IUserValidator interface {
	UserValidator(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidator(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスは必須項目です"),
			validation.RuneLength(1, 50).Error("メールアドレスは40文字までです"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードは必須項目です"),
			validation.RuneLength(6, 30).Error("パスワードは10文字までです。"),
		),
	)
}
