package contactService

import (
	"net/smtp"
	"os"
	contactModels "portfolioAPI/contact/models"
)

type ContactService struct{}

func NewContactService() *ContactService {
	return &ContactService{}
}

func (service *ContactService) SendMail(messageModel contactModels.ContactModel) error {
	smtpClient := service.getSmtpClient()
	adminMail := os.Getenv("ADMIN_EMAIL")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	to := []string{adminMail}
	msg := []byte(
		"Subject: New Message from Portfolio" +
			"\r\n" +
			messageModel.Message + "\r\n" +
    "Message from: " + messageModel.From + "\r\n" + 
    "Name: " + messageModel.Name + "\r\n")

  err := smtp.SendMail(smtpHost + ":" + smtpPort, smtpClient, smtpUsername, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func (service *ContactService) getSmtpClient() smtp.Auth {
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")

	return smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
}
