package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"merge_pdf/src/i18n"
)

// ChangeLanguage met à jour la langue et tout le texte de l'interface.
func ChangeLanguage(choice string, w fyne.Window, btnAdd, btnRemove, btnUp, btnDown, btnClear, btnMerge, btnDoc *widget.Button, langSelect *widget.Select) {
	// --- Sélection de la langue ---
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

	// --- Mise à jour du titre de la fenêtre ---
	w.SetTitle(i18n.T("WindowTitle"))

	// --- Mise à jour des libellés de tous les boutons ---
	btnAdd.SetText(i18n.T("AddPDF"))
	btnRemove.SetText(i18n.T("Remove"))
	btnUp.SetText(i18n.T("MoveUp"))
	btnDown.SetText(i18n.T("MoveDown"))
	btnClear.SetText(i18n.T("ClearList"))
	btnMerge.SetText(i18n.T("Merge"))
	btnDoc.SetText(i18n.T("Documentation"))

	// --- Mise à jour des options du menu de langue ---
	langSelect.Options = []string{"English", "Français", "Deutsch", "Español", "Italiano"}
	langSelect.Refresh()
}
