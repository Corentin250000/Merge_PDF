package ui

import (
	"fmt"

	"merge_pdf/src/i18n"
	"merge_pdf/src/os_wrappers"
	"merge_pdf/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func CreateMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow(i18n.T("WindowTitle"))
	w.Resize(fyne.NewSize(700, 500))

	files := []string{}
	fileChecks := []*widget.Check{}
	listContainer := container.NewVBox()

	// --- Fonction de mise à jour de la liste ---
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

	// --- Boutons principaux ---
	btnAdd := widget.NewButton(i18n.T("AddPDF"), func() {
		selectedFiles, err := os_wrappers.SelectMultipleFiles(w)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if selectedFiles != nil {
			files = append(files, selectedFiles...)
			refreshList()
		}
	})

	btnRemove := widget.NewButton(i18n.T("Remove"), func() {
		newFiles := []string{}
		for i, chk := range fileChecks {
			if !chk.Checked {
				newFiles = append(newFiles, files[i])
			}
		}
		files = newFiles
		refreshList()
	})

	btnUp := widget.NewButton(i18n.T("MoveUp"), func() {
		for i := 1; i < len(files); i++ {
			if fileChecks[i].Checked && !fileChecks[i-1].Checked {
				files[i], files[i-1] = files[i-1], files[i]
				fileChecks[i].Checked, fileChecks[i-1].Checked = false, true
			}
		}
		refreshList()
	})

	btnDown := widget.NewButton(i18n.T("MoveDown"), func() {
		for i := len(files) - 2; i >= 0; i-- {
			if fileChecks[i].Checked && !fileChecks[i+1].Checked {
				files[i], files[i+1] = files[i+1], files[i]
				fileChecks[i].Checked, fileChecks[i+1].Checked = false, true
			}
		}
		refreshList()
	})

	btnClear := widget.NewButton(i18n.T("ClearList"), func() {
		files = []string{}
		refreshList()
	})

	btnMerge := widget.NewButton(i18n.T("Merge"), func() {
		if len(files) == 0 {
			dialog.ShowInformation("Error", i18n.T("ErrorNoFiles"), w)
			return
		}
		outFile, err := os_wrappers.SelectSaveFile(w)
		if err != nil || outFile == "" {
			return
		}
		err = api.MergeCreateFile(files, outFile, false, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Merge failed: %w", err), w)
			return
		}
		dialog.ShowInformation("Success", i18n.T("SuccessMerged")+" "+outFile, w)
	})

	// --- Documentation dynamique ---
	btnDoc := widget.NewButton(i18n.T("Documentation"), func() {
		dialog.ShowInformation(i18n.T("DocTitle"), i18n.T("DocText"), w)
	})

	// --- Sélecteur de langue ---
	var langSelect *widget.Select
	langSelect = widget.NewSelect(
		[]string{"English", "Français", "Deutsch", "Español", "Italiano"},
		func(choice string) {
			utils.ChangeLanguage(choice, w, btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc, langSelect)
		},
	)

	// --- Layout général ---
	scroll := container.NewVScroll(listContainer)
	scroll.SetMinSize(fyne.NewSize(450, 400))

	buttons := container.NewVBox(
		langSelect,
		btnAdd,
		btnRemove,
		btnUp,
		btnDown,
		btnClear,
		btnMerge,
		btnDoc,
	)

	content := container.NewHSplit(scroll, buttons)
	content.Offset = 0.75
	w.SetContent(content)
	refreshList()

	return w
}
