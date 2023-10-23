package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo-app/app/internal/handlers"
	"todo-app/app/internal/repositories"
	"todo-app/app/internal/servers"
	"todo-app/app/internal/services"
)

func init() {
	//godotenv initialization
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	//viper initialization
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Viper init error: %s", err.Error())
	}
}

func main() {
	db, err := repositories.NewPostgresDB(&repositories.DBConfigs{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.db-name"),
		SSLMode:  viper.GetString("db.ssl-mode"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB-PING", db.Ping())
	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)
	server := servers.NewHTTPServer(viper.GetString("http.port"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
