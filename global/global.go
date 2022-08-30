package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	Trans    ut.Translator
	Validate *validator.Validate
	DB       *gorm.DB
	Esclient *elastic.Client
)
