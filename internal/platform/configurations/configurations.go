package configurations

import "os"

// Port : Getter for port
func Port() string {
	if value, ok := os.LookupEnv("CONFIG_PORT"); ok {
		return value
	}
	return ":3000"
}
