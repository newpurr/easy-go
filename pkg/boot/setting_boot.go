package boot

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/pkg/setting"
	"time"
)

var DefaultSettingBootloader = NewSettingBootloader("configs/")

type SettingBootloader struct {
	Paths []string
}

func NewSettingBootloader(paths ...string) *SettingBootloader {
	return &SettingBootloader{Paths: paths}
}

func (sb SettingBootloader) Boot() error {
	s, err := setting.NewSetting(sb.Paths...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &application.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &application.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &application.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &application.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &application.EmailSetting)
	if err != nil {
		return err
	}

	application.Setting = s
	application.AppSetting.DefaultContextTimeout *= time.Second
	application.JWTSetting.Expire *= time.Second
	application.ServerSetting.ReadTimeout *= time.Second
	application.ServerSetting.WriteTimeout *= time.Second
	if application.ServerSetting.HttpPort == "" {
		application.ServerSetting.HttpPort = "8888"
	}
	if application.ServerSetting.RunMode == "" {
		application.ServerSetting.RunMode = "debug"
	}

	return nil
}
