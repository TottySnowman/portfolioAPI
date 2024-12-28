package fileController

import (
	"net/http"
	fileService "portfolioAPI/fileUpload/services"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService *fileService.FileService
}

func NewFileController(fileService *fileService.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

func (con *FileController) UploadFile(context *gin.Context) {
	_, err := context.FormFile("file")
	//file, err := context.FormFile("file")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}
	path, pathExists := context.GetPostForm("uploadPath")
	if !pathExists {

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	println(path)
}

func (con *FileController) DeleteFile(context *gin.Context) {
	path, pathExists := context.GetPostForm("deletePath")
	if !pathExists {

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	println(path)
}
