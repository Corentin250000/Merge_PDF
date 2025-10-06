package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// ---------------------------
// Windows file picker (multi-selection)
// ---------------------------
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

// ---------------------------
// Windows save file dialog
// ---------------------------
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

func main() {
	a := app.NewWithID("com.ceramaret.fusionpdf")
	w := a.NewWindow("PDF Merger")
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

	// --- Buttons ---
	btnAdd := widget.NewButton("Add PDF(s)", func() {
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

	btnRemove := widget.NewButton("Remove", func() {
		newFiles := []string{}
		for i, chk := range fileChecks {
			if !chk.Checked {
				newFiles = append(newFiles, files[i])
			}
		}
		files = newFiles
		refreshList()
	})

	btnUp := widget.NewButton("Move Up", func() {
		for i := 1; i < len(files); i++ {
			if fileChecks[i].Checked && !fileChecks[i-1].Checked {
				files[i], files[i-1] = files[i-1], files[i]
				fileChecks[i].Checked, fileChecks[i-1].Checked = false, true
			}
		}
		refreshList()
	})

	btnDown := widget.NewButton("Move Down", func() {
		for i := len(files) - 2; i >= 0; i-- {
			if fileChecks[i].Checked && !fileChecks[i+1].Checked {
				files[i], files[i+1] = files[i+1], files[i]
				fileChecks[i].Checked, fileChecks[i+1].Checked = false, true
			}
		}
		refreshList()
	})

	btnClear := widget.NewButton("Clear List", func() {
		files = []string{}
		refreshList()
	})

	btnMerge := widget.NewButton("Merge", func() {
		if len(files) == 0 {
			dialog.ShowInformation("Error", "No files to merge.", w)
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
		dialog.ShowInformation("Success", "Merged successfully:\n"+outFile, w)
	})

	docText := `=== PDF Merger - Documentation ===

1. Add PDF files
   Click "Add PDF(s)" to select one or more PDF files.

2. Arrange the order
   Check one or more files, then click "Move Up" or "Move Down".
   The order shown will be the order of pages in the final PDF.

3. Remove or clear
   - "Remove" deletes only the checked files.
   - "Clear List" removes all files from the list.

4. Merge
   Click "Merge" and choose a destination file name.
   The software will create a single merged PDF file.

--- ⚠️ Known Issue ---
If you choose as output file a PDF that is already in the list:
- the resulting file will be empty and corrupted
- the program will report it as an empty file

Solution:
Always use a different name for the output file (e.g. merged_result.pdf).`

	btnDoc := widget.NewButton("Documentation", func() {
		dialog.ShowInformation("Documentation", docText, w)
	})

	// --- Layout ---
	scroll := container.NewVScroll(listContainer)
	scroll.SetMinSize(fyne.NewSize(450, 400))

	buttons := container.NewVBox(btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc)
	content := container.NewHSplit(scroll, buttons)
	content.Offset = 0.75
	w.SetContent(content)

	refreshList()
	w.ShowAndRun()
}
