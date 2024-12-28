package fileServices

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
