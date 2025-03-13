package routes

import (
	"fmt"
	"go-tutorial/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	// User routes
	fmt.Println("Routes")
	app.Get("/", controllers.GetHome)               // Display users
	app.Get("/users", controllers.GetUsers)         // Display users
	app.Get("/api/users", controllers.GetUsersJSON) // Display users
	// app.Post("/users", controllers.AddUser)   // Add a new user
	// ✅ Allow POST requests to add a user
	app.Post("/users", controllers.AddUser)
	app.Get("/users/edit/:id", controllers.EditUserPage) // Edit form
	// app.Post("/users/edit/:id", controllers.UpdateUser)  // Update user
}
