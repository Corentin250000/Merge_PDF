//go:build linux
// +build linux

package os_wrappers

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

// Detects system language on Linux
func GetSystemLangCode() string {
	lang := os.Getenv("LANG")
	if len(lang) >= 2 {
		return lang[:2]
	}
	return "en"
}

// Opens a non-blocking file dialog (single selection, reliable on Linux)
func SelectMultipleFiles(main fyne.Window) ([]string, error) {
	var files []string
	d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if reader != nil {
			files = append(files, reader.URI().Path())
		}
	}, main)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	d.Show()
	return files, nil
}

// Opens a save dialog for the output PDF
func SelectSaveFile(main fyne.Window) (string, error) {
	var out string
	d := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer != nil {
			out = writer.URI().Path()
		}
	}, main)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	d.Show()
	return out, nil
}
