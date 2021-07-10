package boot

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/internal/model"
)

type MysqlBootloader struct {
}

func NewMysqlBootloader() *MysqlBootloader {
	return &MysqlBootloader{}
}

func (sb MysqlBootloader) Boot() error {
	var err error
	application.DBEngine, err = model.NewDBEngine(application.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
