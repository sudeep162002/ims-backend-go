package user_routes

import (
	"github.com/gin-gonic/gin"
	user_controllers "github.com/sudeep162002/ims-go-backend/controller"
	"github.com/sudeep162002/ims-go-backend/middleware"
)

// app.get('/get-users',authenticateToken, UsersHandler.getUsers);
//
//	app.get('/get-users/:id',authenticateToken, UsersHandler.getUsersById);
//	app.post('/insert-user',authenticateToken, UsersHandler.insertUser);
//	app.put('/update-user/:id',authenticateToken,UsersHandler.updateUser);

func SetupHelloRoutes(r *gin.Engine) {

	// Define routes related to hello
	userroutes := r.Group("/")
	{

		// r.Use(middleware.Auth())
		// userroutes.GET("/", user_controllers.HelloUser)
		userroutes.GET("/get-users", middleware.Auth(), user_controllers.GetUsers)
		userroutes.GET("/get-users/:id", middleware.Auth(), user_controllers.GetUsersById)
		userroutes.POST("/insert-user", middleware.Auth(), user_controllers.InsertUser)
		userroutes.PUT("/update-user/:id", middleware.Auth(), user_controllers.UpdateUser)
	}
}
