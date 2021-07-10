package boot

type Bootloader interface {
	Boot() error
}
