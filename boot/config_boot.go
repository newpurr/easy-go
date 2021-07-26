package boot

import (
	"github.com/newpurr/easy-go/config"
)

type ConfigBootloader struct {
	filename string
	readfns  []func(*config.Config) error
}

func NewConfigBootloader(filename string, fns ...func(*config.Config) error) *ConfigBootloader {
	return &ConfigBootloader{filename: filename, readfns: fns}
}

func (c ConfigBootloader) Boot() error {
	conf, err := config.LoadConfig(c.filename)
	if err != nil {
		return err
	}

	for _, readfn := range c.readfns {
		err := readfn(conf)
		if err != nil {
			return err
		}
	}

	return nil
}
