package v1

import (
	"admin-api-gateway/api/models"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func RandomGenerator() int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}

func SendCodeGmail(client models.Client) (string, error) {
	email := "abdulazizxoshimov22@gmail.com"
	password := "hxytgczqprxfsltu "

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("test", email, password, smtpHost)

	randomNumber := RandomGenerator()
	randomNumberString := strconv.Itoa(randomNumber)

	to := []string{client.Email}
	msg := []byte(randomNumberString)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, to, msg)
	if err != nil {
		return "", err
	}
	return randomNumberString, nil
}
