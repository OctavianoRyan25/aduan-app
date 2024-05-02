// internal/complaint/image_util.go
package complaint

import (
	"io"

	storage "github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/storage"
)

func SaveImage(path string, data io.Reader) error {
	return storage.SaveImage(path, data)
}
