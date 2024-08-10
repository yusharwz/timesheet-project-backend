package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmailActivatedAccount(email, code, unique string) error {
	host := os.Getenv("HOST_FRONTEND")
	t := fmt.Sprintf("e=%s&unique=%s", email, unique)
	t = base64.StdEncoding.EncodeToString([]byte(t))
	url := fmt.Sprintf("%s/accounts/activation?t=%s", host, t)

	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Activation  Account")
	m.SetBody("text/plain", "Click link to activated your account: \n"+url+"\n \nThis is information about your account for Login after activation: \nEmail: "+email+"\nPassword: "+code+"\n\nnote: For the security of your account, we recommend that you change your password immediately after logging in")

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendNewPassword(email, newPassword string) error {
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "New Password")
	m.SetBody("text/plain", "Your new password: "+newPassword+"\n\nnote: For the security of your account, we recommend that you change your password immediately after logging in")

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendNotificationToTrainer(email, name, status, statusBy string) error {
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	link := fmt.Sprintf("%s/approvals/on-progress", os.Getenv("HOST_FRONTEND"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Approval Notification")
	m.SetBody("text/plain", "Hai, "+name+"\nPengajuan timesheet anda sudah di "+status+" oleh "+statusBy+". Silahkan buka aplikasi untuk melihat update dengan klik link berikut "+link)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendNotificationToManager(email, name string) error {
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	link := fmt.Sprintf("%s/approvals/on-progress", os.Getenv("HOST_FRONTEND"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Approval Notification")
	m.SetBody("text/plain", "Ada pengajuan timesheet baru oleh trainer "+name+". Silahkan buka aplikasi untuk melihat update dengan klik link berikut "+link)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendNotificationToBenefit(email, name string) error {
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	link := fmt.Sprintf("%s/approvals/on-progress", os.Getenv("HOST_FRONTEND"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Approval Notification")
	m.SetBody("text/plain", "Ada pengajuan timesheet oleh trainer "+name+" yang sudah disetujui oleh Manager IT. Silahkan buka aplikasi untuk melihat update dengan klik link berikut "+link)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
