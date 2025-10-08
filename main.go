package main

import (
	"fyne.io/fyne/v2/app"
	"merge_pdf/src/i18n"
	"merge_pdf/src/os_wrappers"
	"merge_pdf/src/ui"
)

func main() {
	// Detect system language and initialize localization
	systemLang := os_wrappers.GetSystemLangCode()
	i18n.InitI18n(systemLang)

	// Create main app instance
	a := app.NewWithID("com.ceramaret.fusionpdf")

	// Create and show main window
	w := ui.CreateMainWindow(a)
	w.ShowAndRun()
}