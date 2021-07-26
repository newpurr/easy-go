package boot

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/validator"
)

type ValidatorBootloader struct {
}

func NewValidatorBootloader() *ValidatorBootloader {
	return &ValidatorBootloader{}
}

func (sb ValidatorBootloader) Boot() error {
	application.Validator = validator.NewCustomValidator()
	application.Validator.Engine()
	binding.Validator = application.Validator
	return nil
}
