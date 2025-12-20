package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaByAPI(otp string, email string, purpose string) error {
	// URL api
	postURL := "http://localhost:8080/email/send_text"

	// data json
	mailReq := MailRequest{
		ToEmail:     email,
		Subject:     "Verify OTP " + purpose,
		MessageBody: "OTP is " + otp,
		Attachment:  "path/to/email",
	}

	// convert struct to json
	requestBody, err := json.Marshal(mailReq)
	if err != nil {
		return err
	}

	fmt.Println("Request body: ", requestBody)
	// create request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// PUT header
	req.Header.Set("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check response status
	fmt.Sprintln("Response status: ", resp.Status)
	return nil
}
