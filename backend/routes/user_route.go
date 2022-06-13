package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/user", controllers.Login)
	app.Post("/auth/:token", controllers.User)
	app.Post("/insertproduct", controllers.CreateProduct)
	app.Get("/userproduct/:usermail", controllers.GetProduct)
	app.Post("/logout", controllers.Logout)
	app.Post("/edit/:ProductId", controllers.EditProduct)
	app.Post("/delete/:productId", controllers.DeleteProduct)

}
