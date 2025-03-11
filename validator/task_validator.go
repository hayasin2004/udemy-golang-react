package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"udemy-golang-react/model"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct {
}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("タスクのタイトルは必須です"),
			validation.RuneLength(1, 20).Error("最大で20文字しか打てません"),
		),
	)
}
