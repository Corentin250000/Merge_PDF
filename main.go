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
// Sélecteur d'entrées : multisélection de fichiers
// ---------------------------
func selectMultipleFilesWindows() ([]string, error) {
	ps := `[System.Reflection.Assembly]::LoadWithPartialName("System.windows.forms") | Out-Null
$fd = New-Object System.Windows.Forms.OpenFileDialog
$fd.Filter = "PDF files (*.pdf)|*.pdf"
$fd.Multiselect = $true
if($fd.ShowDialog() -eq "OK"){ $fd.FileNames }`

	cmd := exec.Command("powershell", "-Command", ps)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // masque la fenêtre PowerShell

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
// Sélecteur de sortie : fichier unique
// ---------------------------
func selectSaveFileWindows() (string, error) {
	ps := `[System.Reflection.Assembly]::LoadWithPartialName("System.windows.forms") | Out-Null
$fd = New-Object System.Windows.Forms.SaveFileDialog
$fd.Filter = "PDF files (*.pdf)|*.pdf"
$fd.FileName = "fusion.pdf"
if($fd.ShowDialog() -eq "OK"){ $fd.FileName }`

	cmd := exec.Command("powershell", "-Command", ps)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // masque la fenêtre PowerShell

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func main() {
	a := app.NewWithID("com.ceramaret.fusionpdf")
	w := a.NewWindow("Fusion PDF")
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

	// --- Boutons ---
	btnAdd := widget.NewButton("Ajouter PDF(s)", func() {
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

	btnRemove := widget.NewButton("Supprimer", func() {
		newFiles := []string{}
		for i, chk := range fileChecks {
			if !chk.Checked {
				newFiles = append(newFiles, files[i])
			}
		}
		files = newFiles
		refreshList()
	})

	btnUp := widget.NewButton("Monter", func() {
		for i := 1; i < len(files); i++ {
			if fileChecks[i].Checked && !fileChecks[i-1].Checked {
				files[i], files[i-1] = files[i-1], files[i]
				fileChecks[i].Checked, fileChecks[i-1].Checked = false, true
			}
		}
		refreshList()
	})

	btnDown := widget.NewButton("Descendre", func() {
		for i := len(files) - 2; i >= 0; i-- {
			if fileChecks[i].Checked && !fileChecks[i+1].Checked {
				files[i], files[i+1] = files[i+1], files[i]
				fileChecks[i].Checked, fileChecks[i+1].Checked = false, true
			}
		}
		refreshList()
	})

	btnClear := widget.NewButton("Effacer la liste", func() {
		files = []string{}
		refreshList()
	})

	btnMerge := widget.NewButton("Fusionner", func() {
		if len(files) == 0 {
			dialog.ShowInformation("Erreur", "Aucun fichier à fusionner", w)
			return
		}
		outFile, err := selectSaveFileWindows()
		if err != nil || outFile == "" {
			return
		}
		err = api.MergeCreateFile(files, outFile, false, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("échec fusion: %w", err), w)
			return
		}
		dialog.ShowInformation("Succès", "Fusion terminée : "+outFile, w)
	})

	docText := `=== Documentation - Fusion PDF ===

1. Ajouter des fichiers PDF
   Cliquez sur "Ajouter PDF(s)" et sélectionnez un ou plusieurs fichiers.

2. Organiser l'ordre
   Cochez un ou plusieurs fichiers puis utilisez "Monter" ou "Descendre".
   L'ordre affiché sera l'ordre final des pages.

3. Supprimer ou effacer
   - "Supprimer" enlève uniquement les fichiers cochés.
   - "Effacer la liste" vide complètement la sélection.

4. Fusionner
   Cliquez sur "Fusionner" et choisissez un nom de fichier final.
   Le programme génère un PDF unique.

--- ⚠️ Bug connu ---
Si vous choisissez comme fichier de sortie un PDF déjà présent dans la liste :
- le fichier final sera vide et corrompu
- le programme indiquera que le fichier est vide

Solution :
Toujours donner un nom différent au fichier final (ex: fusion_result.pdf).`

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
