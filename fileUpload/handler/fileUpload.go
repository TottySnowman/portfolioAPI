package fileHandler

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileUploadHandler struct{}

func (handler *FileUploadHandler) Upload(filePath string, file *multipart.FileHeader) (string, error) {
	switch filePath {
	case "/logo":
		return handler.logoUpload(file)

	case "/cv":
		return handler.cvUpload(file)

	default:
		return "", errors.New("Invalid path")
	}
}

func (handler *FileUploadHandler) logoUpload(file *multipart.FileHeader) (string, error) {
	directory := filepath.Dir("./public/logo/")
	allowedMimeTypes := []string{"image/png", "image/webp", "image/svg+xml", "image/jpeg"}

	return handler.handleFileUpload(file, directory, allowedMimeTypes)
}

func (handler *FileUploadHandler) cvUpload(file *multipart.FileHeader) (string, error) {
	directory := filepath.Dir("./public/cv/")
	allowedMimeTypes := []string{"application/pdf"}

	return handler.handleFileUpload(file, directory, allowedMimeTypes)
}

func (handler *FileUploadHandler) handleFileUpload(file *multipart.FileHeader, filePath string, allowedMimeTypes []string) (string, error) {
	outputPath := filepath.Join(filePath, filepath.Base(file.Filename))

	if err := handler.ensureDirectoryExists(filePath); err != nil {
		return "", err
	}

	openFile, err := handler.openFile(file)
	if err != nil {
		return "", err
	}

	if !handler.validateFileMimeType(openFile, allowedMimeTypes) {
    print("AYO NOT ALLOWED")
		return "", errors.New("Invalid mime type!")
	}

	destinationFile, err := handler.createDestinationFile(outputPath)
	if err != nil {
		return "", err
	}

	err = handler.copyFile(destinationFile, openFile)

	return outputPath, nil
}

func (handler *FileUploadHandler) validateFileMimeType(openFile multipart.File, allowedMimeTypes []string) bool {
	fileType, err := handler.getFileType(openFile)
	if err != nil {
		return false
	}

	for _, allowedMimeType := range allowedMimeTypes {
		if allowedMimeType == fileType {
			return true
		}
	}

	return false
}

func (handler *FileUploadHandler) getFileType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	file.Seek(0, 0)

	return http.DetectContentType(buffer), nil
}

func (handler *FileUploadHandler) ensureDirectoryExists(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return os.MkdirAll(directoryPath, os.ModePerm)
	}
	return nil
}

func (handler *FileUploadHandler) openFile(file *multipart.FileHeader) (multipart.File, error) {
	srcFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	return srcFile, nil
}

func (handler *FileUploadHandler) createDestinationFile(filePath string) (*os.File, error) {
	destinationFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer destinationFile.Close()

	return destinationFile, nil
}

func (handler *FileUploadHandler) copyFile(destinationFile *os.File, openFile multipart.File) error {
	_, err := io.Copy(destinationFile, openFile)
	if err != nil {
		return err
	}

	return nil
}
