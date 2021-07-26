package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	vp        *viper.Viper
	reference map[string]interface{}
}

func (s *Config) Viper() *viper.Viper {
	return s.vp
}

func (s *Config) Read(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Config) ReadWithRefresh(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := s.reference[k]; !ok {
		s.reference[k] = v
	}

	return nil
}

func newConfig(vp *viper.Viper) *Config {
	return &Config{vp, make(map[string]interface{})}
}

func (s *Config) watchDog() {
	s.vp.WatchConfig()
	s.vp.OnConfigChange(func(in fsnotify.Event) {
		_ = s.refreshReference()
	})
}

func (s *Config) refreshReference() error {
	for k, v := range s.reference {
		err := s.Read(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
