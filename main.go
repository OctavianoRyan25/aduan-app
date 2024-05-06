// cmd/main.go
package main

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/middlewares"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/admin"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Initialize DB
	// Konfigurasi koneksi ke database MySQL
	// DB_USER := os.Getenv("DB_USER")
	dsn := "root:@tcp(localhost:3306)/minpro?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate schema
	db.AutoMigrate(&complaint.Complaint{}, &complaint.Image{}, &complaint.Status{}, &user.User{}, &admin.Admin{})

	// Initialize storage
	storage := storage.NewStorage()

	// Memasangkan module complaint
	complaintRepo := complaint.NewComplaintRepository(db)
	complaintUC := complaint.NewComplaintUseCase(complaintRepo)
	complaintController := complaint.NewComplaintController(complaintUC, storage)

	// Memasangkan module user
	userRepo := user.NewUserRepository(db)
	userUC := user.NewUserUseCase(userRepo)
	userController := user.NewUserController(userUC)

	//Memasangkan module admin
	adminRepo := admin.NewAdminRepository(db)
	adminUC := admin.NewAdminUseCase(adminRepo)
	adminController := admin.NewAdminController(adminUC)

	// Routes
	e.Static("/public", "public")
	e.POST("/complaints", complaintController.CreateComplaint, middlewares.Authentication())
	e.GET("/complaints", complaintController.GetAllComplaint, middlewares.Authentication())
	e.GET("/complaints/:id", complaintController.GetComplaintByID, middlewares.Authentication())
	e.PUT("/complaints/:id", complaintController.UpdateComplaint, middlewares.Authentication())
	e.DELETE("/complaints/:id", complaintController.DeleteComplaint, middlewares.Authentication())
	e.POST("/inactive/:id", userController.InactiveUser, middlewares.Authentication())

	e.POST("/register", userController.RegisterUser)
	e.POST("/login", userController.LoginUser)

	//Admin
	e.POST("/register-admin", adminController.RegisterAdmin)
	e.POST("/login-admin", adminController.LoginAdmin)
	e.PUT("/admin/complaints/:id", adminController.UpdateStatusComplaint, middlewares.Authentication())
	e.GET("/admin/complaints", adminController.GetAllComplaint, middlewares.Authentication())
	e.GET("/admin/users", adminController.GetAllUser, middlewares.Authentication())
	e.PUT("/admin/users/:id", adminController.UpdatePasswordUser, middlewares.Authentication())
	e.POST("/reactived/:id", adminController.ActivateUser, middlewares.Authentication())
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
