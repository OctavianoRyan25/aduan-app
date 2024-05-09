package api

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Inisialisasi Cloudinary
	name := os.Getenv("CLOUDINARY_NAME")
	apikey := os.Getenv("CLOUDINARY_API_KEY")
	apisecreet := os.Getenv("CLOUDINARY_API_SECRET")
	cld, err := cloudinary.NewFromParams(name, apikey, apisecreet)
	if err != nil {
		return "", err
	}

	// Upload image to Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder:   "complaints", // Folder in Cloudinary where the image will be stored
		PublicID: filename,     // File name in Cloudinary
	})
	if err != nil {
		return "", err
	}

	// Return public URL of the uploaded image
	return uploadResult.SecureURL, nil
}
