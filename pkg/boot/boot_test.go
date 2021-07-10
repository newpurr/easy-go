package boot

import (
	"fmt"
	"sync"
	"testing"
)

var (
	l           sync.RWMutex
	bootloaders []Bootloader
)

func AddBootloader(b Bootloader) {
	l.Lock()
	defer l.Unlock()
	bootloaders = append(bootloaders, b)
}

func Start(s []Bootloader) {
	l.Lock()
	defer l.Unlock()

	for _, b := range s {
		AddBootloader(b)
	}
	MustBoot()
}

func Boot() error {
	l.RLock()
	defer l.RUnlock()

	return Run(bootloaders)
}

func MustBoot() {
	err := Boot()
	if err != nil {
		panic(err)
	}
}

func Reset() {
	bootloaders = []Bootloader{}
}

type TestBootloader struct {
}

func (t TestBootloader) Boot() error {
	fmt.Println("init")
	//return errors.New("test")
	return nil
}

func TestAddBootloader(t *testing.T) {
	type args struct {
		bootloaders []Bootloader
	}
	testLoader := TestBootloader{}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				[]Bootloader{testLoader},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, l := range tt.args.bootloaders {
				AddBootloader(l)
			}

			if len(bootloaders) != 1 || bootloaders[0] != testLoader {
				t.Error("error.")
			}
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()

			MustBoot()
		})
	}
}
