package api

import (
	"CRUD_API/internal/handler"
	"CRUD_API/internal/repository"
	"CRUD_API/internal/service"
	"CRUD_API/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	dbpool, err := database.NewPostgresDb()
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer dbpool.Close()
	log.Println("Connected to database")

	router := gin.Default()

	userRepo := repository.NewUserRepository(dbpool)
	userService := service.NewUserService(userRepo)
	handler.NewUserHandler(router, userService)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
