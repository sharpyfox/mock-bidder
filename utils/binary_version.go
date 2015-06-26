package utils

import (
	"fmt"
	"runtime"
)

const BINARY_VERSION = "0.0.1"

func Version(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, BINARY_VERSION, runtime.Version())
}
