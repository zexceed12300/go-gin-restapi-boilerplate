package initializers

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func EmbedEnv(env string) {
	s := strings.Split(env, "\n")
	for _, v := range s {
		if strings.Contains(v, "=") {
			vS := strings.Split(v, "=")
			os.Setenv(strings.Trim(vS[0], "\r "), strings.Trim(vS[1], "\r "))
		}
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Cannot load .env file")
	}
}
