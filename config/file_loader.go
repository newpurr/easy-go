package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func MustLoadFile(filename string) *Config {
	c, err := LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	return c
}

func LoadConfig(filename string) (*Config, error) {
	path := filepath.Dir(filename)
	ext := filepath.Ext(filename)
	f := strings.Replace(filepath.Base(filename), ext, "", -1)

	vp := viper.New()
	vp.SetConfigName(f)
	vp.AddConfigPath(path)
	vp.SetConfigType(strings.Replace(ext, ".", "", 1))
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := newConfig(vp)

	go s.watchDog()

	return s, nil
}
