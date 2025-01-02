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

	acceptLanguage := con.getLanguageFromRequest(context)
	uploadPath, err := con.fileService.HandleFileUpload(context.FullPath(), file, acceptLanguage)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"filePath": uploadPath})
}

func (con *FileController) getLanguageFromRequest(context *gin.Context) string {
	acceptLanguage := context.GetHeader("Accept-Language")

	return acceptLanguage
}

func (con *FileController) DeleteFile(context *gin.Context) {
	path := context.PostForm("filePath")
	if path == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	if err := con.fileService.HandleFileDelete(context.FullPath(), path); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}
