package routes

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/middlewares"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/admin"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/bot"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo, complaintController *complaint.ComplaintController, userController *user.UserController, adminController *admin.AdminController, storage *storage.Storage) {
	// Complaint routes

	e.Use(middleware.Logger())

	// Menggunakan middleware CORS dengan konfigurasi default
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.POST("/complaints", complaintController.CreateComplaint, middlewares.Authentication())
	e.GET("/complaints", complaintController.GetAllComplaint, middlewares.Authentication())
	e.GET("/complaints/:id", complaintController.GetComplaintByID, middlewares.Authentication())
	e.PUT("/complaints/:id", complaintController.UpdateComplaint, middlewares.Authentication())
	e.DELETE("/complaints/:id", complaintController.DeleteComplaint, middlewares.Authentication())

	// User routes
	e.POST("/register", userController.RegisterUser)
	e.POST("/login", userController.LoginUser)
	e.POST("/inactive/:id", userController.InactiveUser, middlewares.Authentication())

	// Admin routes
	e.POST("/register-admin", adminController.RegisterAdmin)
	e.POST("/login-admin", adminController.LoginAdmin)
	e.PUT("/admin/complaints/:id", adminController.UpdateStatusComplaint, middlewares.Authentication())
	e.GET("/admin/complaints", adminController.GetAllComplaint, middlewares.Authentication())
	e.GET("/admin/complaints-paginate", adminController.GetAllComplaintWithPaginate, middlewares.Authentication())
	e.GET("/admin/users", adminController.GetAllUser, middlewares.Authentication())
	e.PUT("/admin/users/:id", adminController.UpdatePasswordUser, middlewares.Authentication())
	e.POST("/reactived/:id", adminController.ActivateUser, middlewares.Authentication())

	// Other routes
	e.Static("/public", "public")
	e.POST("/chatbot", bot.ClassifyEnvironmentalIssue)
	// Add other routes as needed
}
