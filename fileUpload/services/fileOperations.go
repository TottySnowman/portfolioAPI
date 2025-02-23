package fileServices

import (
	"mime/multipart"
)

type FileUploader interface {
	Upload(uploadPath string, file *multipart.FileHeader, language string) (string, error)
}
type FileDeleter interface {
	Delete(requestPath string, filePath string) error
}
