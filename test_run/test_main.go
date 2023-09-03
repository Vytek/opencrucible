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
	//Ninth test
	odt, err := os.ReadFile("../test_file/test_file_odt.odt")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(odt)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(odt)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Tenth test
	docx, err := os.ReadFile("../test_file/test_file_docx.docx")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(docx)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(docx)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Eleventh test
	docx_ms, err := os.ReadFile("../test_file/test_file_docx_ms.docx")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(docx_ms)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(docx_ms)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	//Twelfth test
	pdf, err := os.ReadFile("../test_file/test_file_pdf.pdf")
	if err != nil {
		panic(err)
	}
	MIMEType, Ext, _ = opencrucible.DetectFileType(pdf)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
	MIMEType, Ext, _ = opencrucible.DetectFileTypeMIME(pdf)
	fmt.Println(MIMEType)
	fmt.Println(Ext)
}
