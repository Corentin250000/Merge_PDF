package main

import (
	"fyne.io/fyne/v2/app"
	"merge_pdf/src/i18n"
	"merge_pdf/src/os_wrappers"
	"merge_pdf/src/ui"
)

func main() {
	systemLang := os_wrappers.GetSystemLangCode()
	i18n.InitI18n(systemLang)

	a := app.NewWithID("com.ceramaret.fusionpdf")
	w := ui.CreateMainWindow(a)

	w.ShowAndRun()
}
