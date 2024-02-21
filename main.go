package main

import (
	"net/http"
	"os"

	db "github.com/sudeep162002/ims-go-backend/db"
	"github.com/sudeep162002/ims-go-backend/middleware"
	user_routes "github.com/sudeep162002/ims-go-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// make db connection
	db.Initialize()

	// cors for attack protection

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowCredentials: true,          // Access-Control-Allow-Credentials: true
	}))

	//global  Middleware
	r.Use(middleware.LoggerMiddleware())

	// Setup routes
	user_routes.SetupHelloRoutes(r)

	//welcome route
	r.GET("/", func(c *gin.Context) {
		// Create a JSON object
		response := gin.H{
			"message": "welcome to ims backend go version",
		}

		// Return the JSON response with a custom status code
		c.JSON(http.StatusOK, response)
	})

	// health check route
	r.GET("/health", func(c *gin.Context) {
		// Create a JSON object
		response := gin.H{
			"status": "ok",
		}

		// Return the JSON response with a custom status code
		c.JSON(http.StatusOK, response)
	})

	// Run the server on port 3000
	r.Run(":" + os.Getenv("PORT"))
}
