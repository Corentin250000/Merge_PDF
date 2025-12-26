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

	a := app.NewWithID("com.corentin250000.mergepdf")
	w := ui.CreateMainWindow(a)
	w.ShowAndRun()
}