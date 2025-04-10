package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde un archivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
}

// GetEnv obtiene una variable de entorno por su clave
func GetEnv(key string) string {
	return os.Getenv(key)
}