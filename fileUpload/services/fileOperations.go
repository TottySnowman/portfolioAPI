package fileServices

import "mime/multipart"

type FileUploader interface {
	Upload(path string, file *multipart.FileHeader) (string, error)
}
type FileDeleter interface {
	Delete(requestPath string, filePath string) error
}
