package configs

import (
	"os"
)

var (
	Env = os.Getenv("ENV")
)
