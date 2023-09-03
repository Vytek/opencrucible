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
	txt, err := os.ReadFile("../test_file/test_file_txt.txt")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(txt)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Fourth test
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(txt)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Fifth test
	rtf, err := os.ReadFile("../test_file/test_file_rtf.rtf")
	if err != nil {
		panic(err)
	}
	//Sixth and seventh test
	MIMEType, Ext, _ = opencrucible.DetectFileType(rtf)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(rtf)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
}
