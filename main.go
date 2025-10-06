package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"golang.org/x/text/language"
)

// Windows API for system language detection
var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLang = kernel32.NewProc("GetUserDefaultUILanguage")
)

// GetWindowsLangCode returns a two-letter ISO code (e.g., "en", "fr")
func GetWindowsLangCode() string {
	r1, _, _ := procGetUserDefaultLang.Call()
	langID := uint16(r1)
	primaryLangID := langID & 0x3FF

	switch primaryLangID {
	case 0x09:
		return "en"
	case 0x0C, 0x3C, 0x40, 0x80:
		return "fr"
	// case 0x07:
	// 	return "de"
	// case 0x0A:
	// 	return "es"
	// case 0x10:
	// 	return "it"
	default:
		return "en" // fallback
	}
}

var localizer *i18n.Localizer
var currentLang string

// -------- Load translations --------
func initI18n(lang string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("lang/active.en.json")
	bundle.LoadMessageFile("lang/active.fr.json")

	if lang != "fr" && lang != "en" {
		lang = "en"
	}

	localizer = i18n.NewLocalizer(bundle, lang)
	currentLang = lang
}

func t(id string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: id})
}

// -------- Windows dialogs --------
func selectMultipleFilesWindows() ([]string, error) {
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

func selectSaveFileWindows() (string, error) {
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

// -------- Main --------
func main() {
	systemLang := GetWindowsLangCode()
	initI18n(systemLang)

	a := app.NewWithID("com.ceramaret.fusionpdf")
	w := a.NewWindow(t("WindowTitle"))
	w.Resize(fyne.NewSize(700, 500))

	files := []string{}
	fileChecks := []*widget.Check{}
	listContainer := container.NewVBox()

	refreshList := func() {
		listContainer.Objects = nil
		fileChecks = []*widget.Check{}
		for _, f := range files {
			chk := widget.NewCheck(f, nil)
			fileChecks = append(fileChecks, chk)
			listContainer.Add(chk)
		}
		listContainer.Refresh()
	}

	btnAdd := widget.NewButton(t("AddPDF"), func() {
		selectedFiles, err := selectMultipleFilesWindows()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if selectedFiles != nil {
			files = append(files, selectedFiles...)
			refreshList()
		}
	})

	btnRemove := widget.NewButton(t("Remove"), func() {
		newFiles := []string{}
		for i, chk := range fileChecks {
			if !chk.Checked {
				newFiles = append(newFiles, files[i])
			}
		}
		files = newFiles
		refreshList()
	})

	btnUp := widget.NewButton(t("MoveUp"), func() {
		for i := 1; i < len(files); i++ {
			if fileChecks[i].Checked && !fileChecks[i-1].Checked {
				files[i], files[i-1] = files[i-1], files[i]
				fileChecks[i].Checked, fileChecks[i-1].Checked = false, true
			}
		}
		refreshList()
	})

	btnDown := widget.NewButton(t("MoveDown"), func() {
		for i := len(files) - 2; i >= 0; i-- {
			if fileChecks[i].Checked && !fileChecks[i+1].Checked {
				files[i], files[i+1] = files[i+1], files[i]
				fileChecks[i].Checked, fileChecks[i+1].Checked = false, true
			}
		}
		refreshList()
	})

	btnClear := widget.NewButton(t("ClearList"), func() {
		files = []string{}
		refreshList()
	})

	btnMerge := widget.NewButton(t("Merge"), func() {
		if len(files) == 0 {
			dialog.ShowInformation("Error", t("ErrorNoFiles"), w)
			return
		}
		outFile, err := selectSaveFileWindows()
		if err != nil || outFile == "" {
			return
		}
		err = api.MergeCreateFile(files, outFile, false, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Merge failed: %w", err), w)
			return
		}
		dialog.ShowInformation("Success", t("SuccessMerged")+" "+outFile, w)
	})

	btnDoc := widget.NewButton(t("Documentation"), func() {
		dialog.ShowInformation(t("DocTitle"), t("DocText"), w)
	})

	langSelect := widget.NewSelect([]string{"English", "Français"}, func(choice string) {
		switch choice {
		case "Français":
			initI18n("fr")
		default:
			initI18n("en")
		}
		w.SetTitle(t("WindowTitle"))
		btnAdd.SetText(t("AddPDF"))
		btnRemove.SetText(t("Remove"))
		btnUp.SetText(t("MoveUp"))
		btnDown.SetText(t("MoveDown"))
		btnClear.SetText(t("ClearList"))
		btnMerge.SetText(t("Merge"))
		btnDoc.SetText(t("Documentation"))
		w.Content().Refresh()
	})
	if systemLang == "fr" {
		langSelect.SetSelected("Français")
	} else {
		langSelect.SetSelected("English")
	}

	scroll := container.NewVScroll(listContainer)
	scroll.SetMinSize(fyne.NewSize(450, 400))
	buttons := container.NewVBox(langSelect, btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc)
	content := container.NewHSplit(scroll, buttons)
	content.Offset = 0.75
	w.SetContent(content)

	refreshList()
	w.ShowAndRun()
}
