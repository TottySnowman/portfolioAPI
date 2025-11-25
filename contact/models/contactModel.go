package contactModels

type ContactModel struct {
	From    string `json:"from" binding:"required,email"`
	Message string `json:"message" binding:"required,min=1"`
	Name    string `json:"name" binding:"required,min=1"`
}
