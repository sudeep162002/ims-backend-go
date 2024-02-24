package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	db "github.com/sudeep162002/ims-go-backend/db"
	"github.com/sudeep162002/ims-go-backend/middleware"
	user_routes "github.com/sudeep162002/ims-go-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
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

		fmt.Println(response)
		// Return the JSON response with a custom status code
		c.JSON(http.StatusOK, response)
	})

	// health check route
	r.GET("/health", func(c *gin.Context) {
		// Create a JSON object
		response := gin.H{
			"status": "ok",
		}
		fmt.Println(response)
		// Return the JSON response with a custom status code
		c.JSON(http.StatusOK, response)
	})

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
