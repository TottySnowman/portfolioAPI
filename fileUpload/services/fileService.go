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

func (service *FileService) HandleFileUpload(uploadPath string, file *multipart.FileHeader, language string) (string, error) {
	uploadedPath, err := service.uploader.Upload(uploadPath, file, language)
	if err != nil {
		return "", err
	}

	convertedPath := service.ConvertLocalPathToPublicPath(uploadedPath)

	return convertedPath, nil
}

func (service *FileService) ConvertLocalPathToPublicPath(localPath string) string {
	splitted := strings.Split(localPath, "/")

	convertedPath := strings.Join(splitted[1:], "/")

  convertedPath = "/" + convertedPath
	return convertedPath
}

func (service *FileService) HandleFileDelete(requestpath string, filePath string) (error) {
	convertedPath := service.ConvertPublicPathToLocalPath(filePath)
	err := service.deleter.Delete(requestpath, convertedPath)
	if err != nil {
		return err
	}

  return nil
}

func (service *FileService) ConvertPublicPathToLocalPath(publicPath string) string {
	convertedPath := "./public" + publicPath
	return convertedPath
}
