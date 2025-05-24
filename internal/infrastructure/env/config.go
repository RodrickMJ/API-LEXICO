package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadDBEnv() string {
	godotenv.Load(".env")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
}
