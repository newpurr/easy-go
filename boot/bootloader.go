package boot

type Bootloader interface {
	Boot() error
}

type BootloaderWrapperBootloader struct {
	f func() Bootloader
}

func BootloaderWrapper(f func() Bootloader) *BootloaderWrapperBootloader {
	return &BootloaderWrapperBootloader{f}
}

func (b *BootloaderWrapperBootloader) Boot() error {
	return b.f().Boot()
}

type FuncWrapperBootloader struct {
	// callable
	fn func() error
}

func FuncWrapper(f func() error) *FuncWrapperBootloader {
	return &FuncWrapperBootloader{f}
}

func (b *FuncWrapperBootloader) Boot() error {
	return b.fn()
}
