package opencrucible

import (
	"os"
	"testing"
)

func TestTXTParser(t *testing.T) {
	txt, err := os.ReadFile("test_file/test_file_txt.txt")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, err := TXTParseToString(txt)
	if err != nil {
		t.Errorf("unable to parse file \n %s", err)
	}
	want := "This is a test file to test library"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTXTFileParser(t *testing.T) {
	got, err := TXTFileParseToString("test_file/test_file_txt.txt")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestODTParser(t *testing.T) {
	odt, err := os.ReadFile("test_file/test_file_odt.odt")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, err := ODTParseToString(odt)
	if err != nil {
		t.Errorf("unable to parse file \n %s", err)
	}
	want := "This is a test for library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestODTFileParser(t *testing.T) {
	got, err := ODTFileParseToString("test_file/test_file_odt.odt")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test for library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRTFParser(t *testing.T) {
	rtf, err := os.ReadFile("test_file/test_file_rtf.rtf")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, err := RTFParseToString(rtf)
	if err != nil {
		t.Errorf("unable to parse file \n %s", err)
	}
	want := "Normalheading 1heading 2heading 3heading 4heading 5heading 6TitleSubtitleThis is a test file to test library "
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRTFFileParser(t *testing.T) {
	got, err := RTFFileParseToString("test_file/test_file_rtf.rtf")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "Normalheading 1heading 2heading 3heading 4heading 5heading 6TitleSubtitleThis is a test file to test library "
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDOCXFileParser(t *testing.T) {
	got, err := DOCXFileParseToString("test_file/test_file_docx.docx")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDOCXMSFileParser(t *testing.T) {
	got, err := DOCXFileParseToString("test_file/test_file_docx_ms.docx")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// This�is�a�test�file�to�test�library
// 54 68 69 73 00 69 73 00 61 00 74 65 73 74 00 66 69 6C 65 00 74 6F 00 74 65 73 74 00 6C 69 62 72 61 72 79 20
// Must be 20 NOT 00

func TestPDFFileParser(t *testing.T) {
	got, err := PDFFileParseToString("test_file/test_file_pdf.pdf")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
