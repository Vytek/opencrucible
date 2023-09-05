package opencrucible

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTXTParser(t *testing.T) {
	txt, err := os.ReadFile(filepath.Join("test_file", "test_file_txt.txt"))
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
	got, err := TXTFileParseToString(filepath.Join("test_file", "test_file_txt.txt"))
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
	odt, err := os.ReadFile(filepath.Join("test_file", "test_file_odt.odt"))
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
	got, err := ODTFileParseToString(filepath.Join("test_file", "test_file_odt.odt"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test for library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestOpenOfficeMetadata(t *testing.T) {
	got, err := OpenOfficeFileMetadata(filepath.Join("test_file", "test_file_odt.odt"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "LibreOffice/7.0.5.2$Windows_X86_64 LibreOffice_project/64390860c6cd0aca4beafafcfd84613dd9dfb63a"
	t.Logf("Parsed: %s", got.Meta.Generator)
	t.Logf("Parsed: %s", got.Meta.Title)
	t.Logf("Parsed: %s", got.Meta.CreationDate)
	if got.Meta.Generator != want {
		t.Errorf("got %q, wanted %q", got.Meta.Generator, want)
	}
}

func TestRTFParser(t *testing.T) {
	rtf, err := os.ReadFile(filepath.Join("test_file", "test_file_rtf.rtf"))
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
	got, err := RTFFileParseToString(filepath.Join("test_file", "test_file_rtf.rtf"))
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
	got, err := DOCXFileParseToString(filepath.Join("test_file", "test_file_docx.docx"))
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
	got, err := DOCXFileParseToString(filepath.Join("test_file", "test_file_docx_ms.docx"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDOCFileParser(t *testing.T) {
	got, err := DOCFileParseToString(filepath.Join("test_file", "test_file_doc.doc"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library\r"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// This�is�a�test�file�to�test�library
// 54 68 69 73 00 69 73 00 61 00 74 65 73 74 00 66 69 6C 65 00 74 6F 00 74 65 73 74 00 6C 69 62 72 61 72 79 20
// Must be 20 NOT 00

func TestPDFFileParser(t *testing.T) {
	got, err := PDFFileParseToString(filepath.Join("test_file", "test_file_pdf.pdf"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "This is a test file to test library"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPPTXFileParser(t *testing.T) {
	got, err := PPTXFileParseToString(filepath.Join("test_file", "test_file_pptx.pptx"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "Slide 1:\nTest\nThis \nis\n a \ntest\n \nfile\n \nto\n \ntest\n \nlibrary"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestOfficeMetadata(t *testing.T) {
	got, err := OfficeFileMetadata(filepath.Join("test_file", "test_file_docx_ms.docx"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "Enrico Speranza"
	t.Logf("Parsed: %s", got.Title)
	t.Logf("Parsed: %s", got.Created)
	t.Logf("Parsed: %s", got.LastModifiedBy)
	if got.Creator != want {
		t.Errorf("got %q, wanted %q", got.Creator, want)
	}
}

func TestOfficePPTXMetadata(t *testing.T) {
	got, err := OfficeFileMetadata(filepath.Join("test_file", "test_file_pptx.pptx"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "Presentazione standard di PowerPoint"
	t.Logf("Parsed: %s", got.Title)
	t.Logf("Parsed: %s", got.Created)
	t.Logf("Parsed: %s", got.LastModifiedBy)
	if got.Title != want {
		t.Errorf("got %q, wanted %q", got.Title, want)
	}
}

func TestPDFMetdata(t *testing.T) {
	got, err := PDFFileMetadata(filepath.Join("test_file", "test_file_pdf.pdf"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	want := "1.4"
	t.Logf("Parsed: %s", got.PdfVersion)
	t.Logf("Parsed: %s", got.GetAuthor())
	t.Logf("Parsed: %s", got.GetDate())
	t.Logf("Parsed: %s", got.GetCreator())
	t.Logf("Parsed: %s", got.GetTitle())
	if got.PdfVersion != want {
		t.Errorf("got %q, wanted %q", got.PdfVersion, want)
	}
}

func TestPDFDetect(t *testing.T) {
	pdf, err := os.ReadFile(filepath.Join("test_file", "test_file_pdf_1000.pdf"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, _, err := DetectFileTypeMIME(pdf)
	if err != nil {
		t.Errorf("unable to detect file \n %s", err)
	}
	want := "application/pdf"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPPTXDetect(t *testing.T) {
	pdf, err := os.ReadFile(filepath.Join("test_file", "test_file_pptx.pptx"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, _, err := DetectFileTypeMIME(pdf)
	if err != nil {
		t.Errorf("unable to detect file \n %s", err)
	}
	want := "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDOCDetect(t *testing.T) {
	pdf, err := os.ReadFile(filepath.Join("test_file", "test_file_doc.doc"))
	if err != nil {
		t.Errorf("error loading file \n %s", err)
	}
	got, _, err := DetectFileTypeMIME(pdf)
	if err != nil {
		t.Errorf("unable to detect file \n %s", err)
	}
	want := "application/x-ole-storage"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
