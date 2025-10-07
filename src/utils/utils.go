package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"merge_pdf/src/i18n"
)

// ChangeLanguage updates the application language and refreshes all UI elements.
func ChangeLanguage(choice string, w fyne.Window, btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc *widget.Button, langSelect *widget.Select) {
	// Select the corresponding language
	switch choice {
	case "Français":
		i18n.InitI18n("fr")
	case "Deutsch":
		i18n.InitI18n("de")
	case "Español":
		i18n.InitI18n("es")
	case "Italiano":
		i18n.InitI18n("it")
	default:
		i18n.InitI18n("en")
	}

	// Update window title
	w.SetTitle(i18n.T("WindowTitle"))

	// Update all UI button labels
	btnAdd.SetText(i18n.T("AddPDF"))
	btnRemove.SetText(i18n.T("Remove"))
	btnUp.SetText(i18n.T("MoveUp"))
	btnDown.SetText(i18n.T("MoveDown"))
	btnClear.SetText(i18n.T("ClearList"))
	btnMerge.SetText(i18n.T("Merge"))
	btnDoc.SetText(i18n.T("Documentation"))

	// Refresh language selector options
	langSelect.Options = []string{"English", "Français", "Deutsch", "Español", "Italiano"}
	langSelect.Refresh()
}
