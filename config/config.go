package config

import (
	"fmt"
	"strconv"

	"os"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version     string
	Port        int
	ServiceName string
	JwtSecret   string
}

func loadConfig() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("faild to load env", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	port := os.Getenv("PORT")
	serviceName := os.Getenv("SERVICE_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}

	if port == "" {
		fmt.Println("port is required")
		os.Exit(1)
	}

	if serviceName == "" {
		fmt.Println("service name is required")
		os.Exit(1)
	}

	if jwtSecret == "" {
		fmt.Println("jwt secret is required")
		os.Exit(1)
	}

	converted_port, err := strconv.ParseInt(port, 10, 64)

	if err != nil {
		fmt.Println("port need must number", err)
		os.Exit(1)
	}

	configurations = &Config{
		Version:     version,
		Port:        int(converted_port),
		ServiceName: serviceName,
		JwtSecret:   jwtSecret,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
