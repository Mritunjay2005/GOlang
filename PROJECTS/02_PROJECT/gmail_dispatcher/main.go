package main

import (
	"bytes"
	"html/template"
	"sync"
	"log"
	"github.com/codersgyan/email-dispatcher/config"
	
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	config.LoadEnv()
	
	recipientChannel := make(chan Recipient)

	go func() {
		err:= loadRecipient("./emails.csv", recipientChannel)
		if err != nil {
        log.Printf("Fatal: could not load recipients: %v", err)
        // channel is already closed via defer, workers will exit cleanly
    }
	}()
	var wg sync.WaitGroup
	workerCount := 5
// using the comncept of multi threading
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg, )
	}

	wg.Wait()
}
//passing through tempalte 
func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		 log.Fatal("Failed to parse template:", err)
	}
	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
