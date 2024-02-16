package routes

import (
	"log"
	config "mock-gm-home-service/configuration"
	"mock-gm-home-service/controllers"
	"mock-gm-home-service/database"
	"mock-gm-home-service/repositories"
	"mock-gm-home-service/services"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	route := gin.Default()

	env := config.LoadConfig()
	if env != nil {
		log.Fatalf("ERROR LOAD CONFIG FILE '%v'", env)
	}

	// Provide database
	db, err := database.NewPostgresqlDatabase()
	if err != nil {
		log.Fatalf("Cannot connect to database. '%v'", err)
	}

	// Provide repositories
	customerRepo := repositories.NewTbTCustomerRepo(db)
	serviceRepo := repositories.NewTbTServiceRepo(db)

	// Provide Services
	customerService := services.NewCustomerService(
		customerRepo,
		serviceRepo,
	)

	// Provide Controllers
	serviceController := controllers.NewCustomerController(customerService)

	// Provide Routes
	route.GET("/gin/api/get-service", serviceController.GetService)
	route.POST("/gin/api/get-customers", serviceController.GetCustomer)

	PORT := os.Getenv("PORT")

	// Config Port
	route.Run(PORT)

	return route
}
