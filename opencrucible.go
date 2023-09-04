package opencrucible

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/h2non/filetype"
	"github.com/lu4p/cat"
)

// Version exposes the current package version.
const Version = "0.0.2"

//Detects

func DetectFileTypeMIME(StreamToDetect []byte) (string, string, error) {
	// If no StreamToDetect was given, return an error with a message.
	if len(StreamToDetect) == 0 {
		return "", "", errors.New("stream to parse is empty")
	}
	mtype, err := mimetype.DetectReader(bytes.NewReader(StreamToDetect))
	return mtype.String(), mtype.Extension(), err
}

func DetectFileType(StreamToDetect []byte) (string, string, error) {
	kind, _ := filetype.Match(StreamToDetect)
	if kind == filetype.Unknown {
		return "", "", errors.New("unknown file type")
	} else {
		return kind.MIME.Value, kind.Extension, nil
	}
}

//Parsers

func TXTParseToString(StreamToParse []byte) (string, error) {
	// If no StreamToDetect was given, return an error with a message.
	if len(StreamToParse) == 0 {
		return "", errors.New("stream to parse is empty")
	}
	txt, err := cat.FromBytes(StreamToParse)
	return txt, err
}

func RTFParseToString(StreamToParse []byte) (string, error) {
	// If no StreamToDetect was given, return an error with a message.
	if len(StreamToParse) == 0 {
		return "", errors.New("stream to parse is empty")
	}
	rtf, err := cat.FromBytes(StreamToParse)
	return rtf, err
}

func DOCXFileParseToString(FileToParse string) (string, error) {
	// extract text from an XLSX file
	file, err := os.Open(FileToParse)
	if err != nil {
		return "", fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()
	stat, _ := file.Stat()
	docx, err := DOCX2Text(file, stat.Size())
	return docx, err
}

func ODTParseToString(StreamToParse []byte) (string, error) {
	// If no StreamToDetect was given, return an error with a message.
	if len(StreamToParse) == 0 {
		return "", errors.New("stream to parse is empty")
	}
	odt, err := cat.FromBytes(StreamToParse)
	return odt, err
}