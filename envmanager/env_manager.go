package envmanager

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	ProjectName string
	ProjectId   string
}

func GetEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return EnvVariables{
		ProjectName: os.Getenv("PROJECT_NAME"),
		ProjectId:   os.Getenv("PROJECT_ID"),
	}
}
