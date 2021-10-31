package config

var (
	debugMode = false
)

func SetDebugMode(mode bool) {
	debugMode = mode
}

func IsDebugMode() bool {
	return debugMode
}
