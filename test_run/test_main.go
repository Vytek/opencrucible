package main

import (
	"fmt"
	"os"

	"github.com/vytek/opencrucible"
)

func main() {
	//First test
	testBytes := []byte("This random text has a MIME type of text/plain; charset=utf-8.")
	MIMEType, Ext, _ := opencrucible.DetectFileTypeMIME(testBytes)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Second test
	MIMEType, Ext, _ = opencrucible.DetectFileType(testBytes)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Third test
	dat, err := os.ReadFile("../test_file/test_file_txt.txt")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(dat)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Fourth test
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(dat)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
}
