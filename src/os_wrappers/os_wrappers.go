package os_wrappers

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// Windows kernel32 DLL for user language detection
var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLang = kernel32.NewProc("GetUserDefaultUILanguage")
)

// GetWindowsLangCode detects the Windows UI language
func GetWindowsLangCode() string {
	r1, _, _ := procGetUserDefaultLang.Call()
	langID := uint16(r1)
	primaryLangID := langID & 0x3FF
	switch primaryLangID {
	case 0x09:
		return "en"
	case 0x0C, 0x3C, 0x40, 0x80:
		return "fr"
	case 0x07:
		return "de"
	case 0x0A:
		return "es"
	case 0x10:
		return "it"
	default:
		return "en"
	}
}

// GetSystemLangCode returns the OS language code (Windows or Linux)
func GetSystemLangCode() string {
	if runtime.GOOS == "windows" {
		return GetWindowsLangCode()
	}
	lang := os.Getenv("LANG")
	if len(lang) >= 2 {
		return lang[:2]
	}
	return "en"
}

// SelectMultipleFiles opens a dialog allowing users to select multiple PDF files
func SelectMultipleFiles(main fyne.Window) ([]string, error) {
	if runtime.GOOS == "windows" {
		return SelectMultipleFilesWindows()
	}
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

// SelectMultipleFilesWindows uses PowerShell to open a multi-file dialog on Windows
func SelectMultipleFilesWindows() ([]string, error) {
	ps := `[System.Reflection.Assembly]::LoadWithPartialName("System.windows.forms") | Out-Null
$fd = New-Object System.Windows.Forms.OpenFileDialog
$fd.Filter = "PDF files (*.pdf)|*.pdf"
$fd.Multiselect = $true
if($fd.ShowDialog() -eq "OK"){ $fd.FileNames }`
	cmd := exec.Command("powershell", "-Command", ps)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, nil
	}
	files := strings.Split(strings.TrimSpace(string(out)), "\r\n")
	return files, nil
}

// SelectSaveFile opens a dialog to choose a destination file path
func SelectSaveFile(main fyne.Window) (string, error) {
	if runtime.GOOS == "windows" {
		return SelectSaveFileWindows()
	}
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

// SelectSaveFileWindows uses PowerShell to open a save dialog on Windows
func SelectSaveFileWindows() (string, error) {
	ps := `[System.Reflection.Assembly]::LoadWithPartialName("System.windows.forms") | Out-Null
$fd = New-Object System.Windows.Forms.SaveFileDialog
$fd.Filter = "PDF files (*.pdf)|*.pdf"
$fd.FileName = "merged.pdf"
if($fd.ShowDialog() -eq "OK"){ $fd.FileName }`
	cmd := exec.Command("powershell", "-Command", ps)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
