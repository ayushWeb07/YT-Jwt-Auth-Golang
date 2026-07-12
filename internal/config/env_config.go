package config

import (
	"fmt"
	"os"
	"strconv"
)

func LoadSingleEnvVar[T string | int | float32 | bool](key string, defaultVal T) T {
	val, exists := os.LookupEnv(key)

	if !exists {
		return defaultVal
	}

	switch any(defaultVal).(type) {
	case string:
		return any(val).(T)

	case int:
		i, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultVal
		}

		return any(i).(T)

	case float32:
		f, err := strconv.ParseFloat(val, 32)

		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultVal
		}

		return any(float32(f)).(T)

	case bool:
		b, err := strconv.ParseBool(val)

		if err != nil {
			fmt.Println("Something went wrong while parsing this env variable: " + err.Error())
			return defaultVal
		}

		return any(b).(T)

	default:
		fmt.Println("Unsupported datatype has been passed for the env var: " + key)
		return defaultVal
	}

}
