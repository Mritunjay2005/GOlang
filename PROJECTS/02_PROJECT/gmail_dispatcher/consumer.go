package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
	"os"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		
	 email := os.Getenv("SMTP_EMAIL")
	 password := os.Getenv("SMTP_PASSWORD")
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := "587"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Just testing our email campaign.")
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker :%d Error parsing template for %s", id, recipient.Email)
			// we can add dlq here
			continue
		}

		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)

		// err = smtp.SendMail(smtpHost+":"+smtpPort, nil, email, []string{recipient.Email}, []byte(msg))
		// if err != nil {
		// 	log.Fatal(err)
		// }

		

auth := smtp.PlainAuth(
	"",
	email,
	password,
	smtpHost,
)

err = smtp.SendMail(
	smtpHost+":"+smtpPort,
	auth,
	email,
	[]string{recipient.Email},
	[]byte(msg),
)

if err != nil {
	log.Printf("Worker %d: Failed sending to %s: %v",
		id,
		recipient.Email,
		err,
	)
	continue
}

// delay so we don't overwhelm the server
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Sent email to %s \n", id, recipient.Email)

	}
}
