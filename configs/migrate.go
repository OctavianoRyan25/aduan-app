package configs

import (
	"os/user"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/admin"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/complaint"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&complaint.Complaint{}, &complaint.Image{}, &complaint.Status{}, &user.User{}, &admin.Admin{}); err != nil {
		return err
	}
	return nil
}
