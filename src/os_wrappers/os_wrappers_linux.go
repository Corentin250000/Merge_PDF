//go:build linux
// +build linux

package os_wrappers

import (
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
)

// Detect system language
func GetSystemLangCode() string {
	lang := os.Getenv("LANG")
	if len(lang) >= 2 {
		return lang[:2]
	}
	return "en"
}

// Use Zenity for native multi-selection dialog
func SelectMultipleFiles(main fyne.Window) ([]string, error) {
	cmd := exec.Command("zenity",
		"--file-selection",
		"--multiple",
		"--title=Select PDF files",
		"--file-filter=*.pdf")

	out, err := cmd.Output()
	if err != nil {
		// Zenity cancelled -> no file selected
		return nil, nil
	}

	// Zenity separates multiple paths with "|"
	files := strings.Split(strings.TrimSpace(string(out)), "|")
	return files, nil
}

// Native save dialog using Zenity
func SelectSaveFile(main fyne.Window) (string, error) {
	cmd := exec.Command("zenity",
		"--file-selection",
		"--save",
		"--confirm-overwrite",
		"--title=Save merged PDF as",
		"--file-filter=*.pdf")

	out, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(out)), nil
}
