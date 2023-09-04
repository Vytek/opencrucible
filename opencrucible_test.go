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
	got, err := TXTFileParserToString("test_file/test_file_txt.txt")
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
	got, err := ODTFileParserToString("test_file/test_file_odt.odt")
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
	got, err := RTFFileParserToString("test_file/test_file_rtf.rtf")
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

func TestPDFMSFileParser(t *testing.T) {
	got, err := PDFFileParserToString("test_file/test_file_pdf.pdf")
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}