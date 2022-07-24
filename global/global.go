package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Trans    ut.Translator
	Validate *validator.Validate
)
