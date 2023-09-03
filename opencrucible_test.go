package opencrucible

import (
	"os"
	"testing"
)

func TestTXTParse(t *testing.T) {
	txt, err := os.ReadFile("test_file/test_file_txt.txt")
	if err != nil {
		panic(err)
	}
	got, err := TXTParseToString(txt)
	if err != nil {
		panic(err)
	}
	want := "This is a test file to test library"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestODTarse(t *testing.T) {
	odt, err := os.ReadFile("test_file/test_file_odt.odt")
	if err != nil {
		panic(err)
	}
	got, err := ODTParseToString(odt)
	if err != nil {
		panic(err)
	}
	want := "This is a test for library\n"
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRTFarse(t *testing.T) {
	rtf, err := os.ReadFile("test_file/test_file_rtf.rtf")
	if err != nil {
		panic(err)
	}
	got, err := RTFParseToString(rtf)
	if err != nil {
		panic(err)
	}
	want := "Normalheading 1heading 2heading 3heading 4heading 5heading 6TitleSubtitleThis is a test file to test library "
	t.Logf("Parsed: %s", got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}