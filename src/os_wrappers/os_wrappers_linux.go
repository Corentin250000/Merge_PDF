//go:build linux
// +build linux

package os_wrappers

import (
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func GetSystemLangCode() string {
	lang := os.Getenv("LANG")
	if len(lang) >= 2 {
		return lang[:2]
	}
	return "en"
}

func SelectMultipleFiles(main fyne.Window) ([]string, error) {
	var files []string
	done := make(chan struct{})
	dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if reader != nil {
			files = append(files, reader.URI().Path())
		}
		close(done)
	}, main).Show()
	<-done
	return files, nil
}

func SelectSaveFile(main fyne.Window) (string, error) {
	var out string
	done := make(chan struct{})
	dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer != nil {
			out = writer.URI().Path()
		}
		close(done)
	}, main).Show()
	<-done
	return out, nil
}
