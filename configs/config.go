package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MySQL      MySQLConfig
	Cloudinary CloudinaryConfig
	ModelAPI   ModelAPIConfig
	Firebase   FirebaseConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type CloudinaryConfig struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}

type ModelAPIConfig struct {
	SoilPredictURL       string
	PlantRecommendation  string
	InformationRetrieval string
}

type FirebaseConfig struct {
	FirebaseKey string
	FirebaseURL string
}

func InitConfig() (*AppConfig, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	}

	return &AppConfig{
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
		Cloudinary: CloudinaryConfig{
			CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
			ApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
			ApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
		},
		ModelAPI: ModelAPIConfig{
			SoilPredictURL:       os.Getenv("SOIL_MODEL_API_URL"),
			PlantRecommendation:  os.Getenv("PLANT_MODEL_API_URL"),
			InformationRetrieval: os.Getenv("INFORMATION_RETRIEVAL_API_URL"),
		},
		Firebase: FirebaseConfig{
			FirebaseKey: os.Getenv("FIREBASE_ADMIN"),
			FirebaseURL: os.Getenv("FIREBASE_URL"),
		},
	}, nil
}
