package boot

import "github.com/pkg/errors"

func MustBoot(s []Bootloader) {
	err := Do(s)
	if err != nil {
		panic(err)
	}
}

func Do(s []Bootloader) (err error) {
	for _, b := range s {
		err = b.Boot()
		if err == nil {
			continue
		}

		err = errors.Wrap(err, "Boot failed")

		goto er
	}

er:
	return err
}
