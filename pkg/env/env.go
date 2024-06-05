package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Get(key string, args ...string) string {
	val := os.Getenv(key)
	if val == "" && len(args) > 0 {
		val = args[0]
	}
	return val
}

func GetInt(key string, args ...int) int {
	value := 0
	if len(args) > 0 {
		value = args[0]
	}
	val := os.Getenv(key)
	if val != "" {
		valInt, err := strconv.Atoi(val)
		if err == nil {
			value = valInt
		}
	}
	return value
}

func GetInt64(key string, args ...int64) int64 {
	var value int64 = 0
	if len(args) > 0 {
		value = args[0]
	}
	val := os.Getenv(key)
	if val != "" {
		valInt, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			value = valInt
		}
	}
	return value
}
