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

// Non-blocking file open for Linux
func SelectMultipleFiles(main fyne.Window) ([]string, error) {
	results := []string{}

	dlg := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		results = append(results, reader.URI().Path())
	}, main)

	dlg.SetFilter(&fyne.FileFilter{
		Extensions: []string{".pdf"},
	})
	dlg.Show()

	// Return immediately; results will be updated asynchronously
	return results, nil
}

func SelectSaveFile(main fyne.Window) (string, error) {
	var out string
	dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			out = writer.URI().Path()
		}
	}, main).Show()
	return out, nil
}
