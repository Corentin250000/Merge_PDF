package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// ID unique requis par Fyne pour Preferences API
	a := app.NewWithID("com.ceramaret.fusionpdf")
	w := a.NewWindow("Fusion PDF")
	w.Resize(fyne.NewSize(600, 400))

	files := []string{}
	selectedIndex := -1

	list := widget.NewList(
		func() int { return len(files) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(fmt.Sprintf("%d. %s", i+1, files[i]))
		},
	)

	list.OnSelected = func(id int) {
		selectedIndex = id
	}

	btnAdd := widget.NewButton("Ajouter PDF", func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {
			if r == nil {
				return
			}
			files = append(files, r.URI().Path())
			list.Refresh()
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fd.Show()
	})

	btnRemove := widget.NewButton("Supprimer", func() {
		if selectedIndex >= 0 && selectedIndex < len(files) {
			files = append(files[:selectedIndex], files[selectedIndex+1:]...)
			list.Refresh()
			selectedIndex = -1
		}
	})

	btnUp := widget.NewButton("Monter", func() {
		if selectedIndex > 0 {
			files[selectedIndex], files[selectedIndex-1] = files[selectedIndex-1], files[selectedIndex]
			list.Refresh()
			list.Select(selectedIndex - 1)
		}
	})

	btnDown := widget.NewButton("Descendre", func() {
		if selectedIndex >= 0 && selectedIndex < len(files)-1 {
			files[selectedIndex], files[selectedIndex+1] = files[selectedIndex+1], files[selectedIndex]
			list.Refresh()
			list.Select(selectedIndex + 1)
		}
	})

	btnClear := widget.NewButton("Effacer la liste", func() {
		files = []string{}
		list.Refresh()
		selectedIndex = -1
	})

	btnMerge := widget.NewButton("Fusionner", func() {
		if len(files) == 0 {
			dialog.ShowInformation("Erreur", "Aucun fichier à fusionner", w)
			return
		}
		save := dialog.NewFileSave(func(uc fyne.URIWriteCloser, _ error) {
			if uc == nil {
				return
			}
			defer uc.Close()
			err := api.MergeCreateFile(files, uc.URI().Path(), false, nil)
			if err != nil {
				dialog.ShowError(fmt.Errorf("échec fusion: %w", err), w)
				return
			}
			dialog.ShowInformation("Succès", "Fusion terminée : "+uc.URI().Path(), w)
		}, w)
		save.SetFileName("fusion.pdf")
		save.Show()
	})

	docText := `=== Documentation – Fusion PDF ===

1. Ajouter des fichiers PDF
   Cliquez sur "Ajouter PDF" et sélectionnez un ou plusieurs fichiers.
   Ils apparaissent dans la liste.

2. Organiser l'ordre
   Sélectionnez un fichier et utilisez "Monter" ou "Descendre".
   L'ordre affiché sera l'ordre final des pages.

3. Supprimer ou effacer
   - "Supprimer" enlève uniquement le fichier sélectionné.
   - "Effacer la liste" vide complètement la sélection.

4. Fusionner
   Cliquez sur "Fusionner" et choisissez un nom de fichier final.
   Le programme génère un PDF unique.

--- Messages d'erreur fréquents ---

- "Fichier verrouillé" :
   Le PDF est ouvert dans Acrobat Reader ou Edge.
   Fermez le fichier et recommencez.

- "Aucun fichier à fusionner" :
   Vous avez cliqué sur Fusionner sans rien ajouter.

--- ⚠️ Bug connu ---
Si vous choisissez comme fichier de sortie un PDF
qui est déjà présent dans la liste à fusionner :
- le fichier final sera vide et corrompu
- le programme indiquera que le fichier est vide

Solution :
Toujours donner un nom différent au fichier final,
par exemple "fusion_result.pdf".`

	btnDoc := widget.NewButton("Documentation", func() {
		dialog.ShowInformation("Documentation", docText, w)
	})

	// Organisation des boutons
	buttons := container.NewVBox(
		btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc,
	)
	content := container.NewHSplit(list, buttons)
	content.Offset = 0.7
	w.SetContent(content)
	w.ShowAndRun()
}
