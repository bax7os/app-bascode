package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl = ""
	Porta  = 0
	// HashKey is used to authenticate the cookies
	HashKey []byte
	// BlockKey is used to crypt the data in the cookies
	BlockKey []byte
)

// Carregar inicializes the ambient variables
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		return
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		Porta = 8080
	}
	ApiUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
