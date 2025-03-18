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
			validation.RuneLength(6, 50).Error("メールアドレスは6から50文字までで入力してください"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードは必須項目です"),
			validation.RuneLength(6, 30).Error("パスワードは1から10文字で入力して下さい。"),
		),
	)
}
