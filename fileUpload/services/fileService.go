package fileServices

import "mime/multipart"

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

	return uploadedPath, nil
}
