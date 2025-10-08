//go:build windows
// +build windows

package os_wrappers

import (
	"os/exec"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
)

// Windows-specific language detection
var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLang = kernel32.NewProc("GetUserDefaultUILanguage")
)

func GetSystemLangCode() string {
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

func SelectMultipleFiles(main fyne.Window) ([]string, error) {
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

func SelectSaveFile(main fyne.Window) (string, error) {
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
