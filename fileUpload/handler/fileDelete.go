package fileHandler

import (
	"errors"
	"os"
)

type FileDeleteHandler struct{}

func (handler *FileDeleteHandler) Delete(requestPath string, filePath string) error {
	switch requestPath {
	case "/logo":
		return handler.handleLogoDelete(filePath)

	default:
		return errors.New("Invalid path")
	}
}

func (handler *FileDeleteHandler) handleLogoDelete(filePath string) error {
  return os.Remove(filePath)
}
