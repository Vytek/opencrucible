package opencrucible

import (
	"bytes"
	"errors"

	"github.com/gabriel-vasile/mimetype"
	"github.com/h2non/filetype"
)

func DetectFileTypeMIME(FileToDetect []byte) (string, string, error) {
	// If no FileToDetect was given, return an error with a message.
	if len(FileToDetect) == 0 {
		return "", "", errors.New("file to detect is empty")
	}
	mtype, err := mimetype.DetectReader(bytes.NewReader(FileToDetect))
	return mtype.String(), mtype.Extension(), err
}

func DetectFileType(FileToDetect []byte) (string, string, error) {
	kind, _ := filetype.Match(FileToDetect)
	if kind == filetype.Unknown {
		return "", "", errors.New("unknown file type")
	} else {
		return kind.MIME.Value, kind.Extension, nil
	}
}
