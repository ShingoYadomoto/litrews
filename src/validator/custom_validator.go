package validator

import (
	"github.com/go-playground/validator"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}

	ErrorMessage struct {
		Field   string
		Message string
	}
)

func New() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func CustomErrorMessages(err error) []ErrorMessage {
	errs := err.(validator.ValidationErrors)
	ems := make([]ErrorMessage, len(errs))

	for i, err := range errs {
		ems[i] = ErrorMessage{err.Field(), err.Tag()}
	}

	return ems
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// 標準のバリデーションエラーメッセージを日本語に対応してくれるやつっぽいけど時間かかりそうなので保留
// func (v *CustomValidator) GetTranslator() (trans ut.Translator) {
//	j := ja.New()
//	uni := ut.New(j, j)
//	trans, _ = uni.GetTranslator("ja")
//
//	return
//}
//func (v *CustomValidator) RegisterMessage(trans ut.Translator) {
//	vn := "required"
//	vm := "{0}は必須です。"
//
//	v.RegisterTranslation(vn, trans, func(ut ut.Translator) error {
//		return ut.Add(vn, vm, true)
//	}, func(ut ut.Translator, fe validator.FieldError) string {
//		t, _ := ut.T(vn, fe.Field())
//
//		return t
//	})
//}
