package configs

import (
	"capstone/repositories/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default value")
	}
}

func InitConfigMySQL() mysql.Config {
	return mysql.Config{
		DBName: os.Getenv("DBName"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBHost: os.Getenv("DBHost"),
		DBPort: os.Getenv("DBPort"),
	}
}

func GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "https://dev-capstone.practiceproject.tech/v1/users/auth/google/callback",
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func InitConfigJWT() string {
	return os.Getenv("SECRET_JWT")
}

func InitConfigCloudinary() string {
	return os.Getenv("CLOUDINARY_URL")
}

func InitConfigKeyChatbot() string {
	return os.Getenv("AI_KEY")
}
