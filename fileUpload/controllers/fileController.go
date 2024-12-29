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
	file, err := context.FormFile("file")
	if err != nil {
		println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}
	uploadPath, err := con.fileService.HandleFileUpload(context.FullPath(), file)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"filePath": uploadPath})
	return
}

func (con *FileController) DeleteFile(context *gin.Context) {
	path, pathExists := context.GetPostForm("deletePath")
	if !pathExists {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	println(path)
}
