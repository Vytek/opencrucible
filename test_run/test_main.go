package main

import (
    "fmt"

    "github.com/vytek/opencrucible"
)

func main() {
	testBytes := []byte("This random text has a MIME type of text/plain; charset=utf-8.")
	MIMEType, Ext, _ := opencrucible.DetectFileTypeMIME(testBytes)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
}