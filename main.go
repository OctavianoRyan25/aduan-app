// cmd/main.go
package main

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/configs"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/admin"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := configs.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	// Auto migrate schema
	err = configs.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

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
	// Start server
	routes.RegisterRoutes(e, complaintController, userController, adminController, storage)
	e.Logger.Fatal(e.Start(":8080"))
}
