package fileServices

type FileUploader interface{
  Upload(path string) error
}
type FileDeleter interface{
  Delete(path string) error
}
