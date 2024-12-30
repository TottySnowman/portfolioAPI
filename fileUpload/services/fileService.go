package fileServices

import (
	"mime/multipart"
	"strings"
)

type FileService struct {
	uploader FileUploader
	deleter  FileDeleter
}

func NewFileService(uploader FileUploader, deleter FileDeleter) *FileService {
	return &FileService{
		uploader: uploader,
		deleter:  deleter,
	}
}

func (service *FileService) HandleFileUpload(path string, file *multipart.FileHeader) (string, error) {
	uploadedPath, err := service.uploader.Upload(path, file)
	if err != nil {
		return "", err
	}

	convertedPath := service.ConvertLocalPathToPublicPath(uploadedPath)

	print(convertedPath)
	return convertedPath, nil
}

func (service *FileService) ConvertLocalPathToPublicPath(localPath string) string {
	splitted := strings.Split(localPath, "/")

	convertedPath := strings.Join(splitted[1:], "/")

	return convertedPath
}
