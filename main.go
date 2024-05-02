// cmd/main.go
package main

import (
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize DB
	// Konfigurasi koneksi ke database MySQL
	// DB_USER := os.Getenv("DB_USER")
	dsn := "root:@tcp(localhost:3306)/minpro?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuat koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate schema
	db.AutoMigrate(&complaint.Complaint{}, &complaint.Image{})

	// Initialize storage
	storage := storage.NewStorage()

	// Initialize complaint module
	complaintRepo := complaint.NewComplaintRepository(db)
	complaintUC := complaint.NewComplaintUseCase(complaintRepo)
	complaintController := complaint.NewComplaintController(complaintUC, storage)

	// Routes
	e.Static("/public", "public")
	e.POST("/complaints", complaintController.CreateComplaint)
	e.GET("/complaints", complaintController.GetAllComplaint)
	e.GET("/complaints/:id", complaintController.GetComplaintByID)
	e.PUT("/complaints/:id", complaintController.UpdateComplaint)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
