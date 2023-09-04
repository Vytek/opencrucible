# opencrucible
Detect type and extract text from different file type. Similar to Tika Project but in Golang.

[![Go Reference](https://pkg.go.dev/badge/github.com/vytek/opencrucible.svg)](https://pkg.go.dev/github.com/vytek/opencrucible)

![Logo OpenCrucible](https://github.com/Vytek/opencrucible/blob/main/doc/OpenCrucibleLogoResize.png)

List of formats read:

| Format  | FileParser | MIME Type |
| ------------- | ------------- | ------------- |
| TXT | X  | text/plain; charset=utf-8 |
| RTF | X | text/rtf |
| ODT | X | application/vnd.oasis.opendocument.text |
| DOCX  | X  | application/vnd.openxmlformats-officedocument.wordprocessingml.document |
