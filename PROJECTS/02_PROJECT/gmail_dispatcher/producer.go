package main

import (
	"encoding/csv"
	"os"
	"io"
	"log"
)

func loadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch)

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

//reading the file 
	r := csv.NewReader(f)
	r.Read() // skip header row

for {
        record, err := r.Read()
        if err == io.EOF {
            break  // end of file, stop
        }
        if err != nil {
            log.Printf("Error reading row: %v", err)
            continue  // skip bad row, don't crash
        }
        ch <- Recipient{
            Name:  record[0],
            Email: record[1],
        }
    }
    return nil
}