# FusionPDF

FusionPDF is a simple graphical tool that allows you to **merge multiple PDF files** into a single document.  
It is written in **Go**, using [Fyne](https://fyne.io) for the GUI and [pdfcpu](https://github.com/pdfcpu/pdfcpu) for PDF processing.  
The application is fully **portable** - no installation or dependencies required.

## ✨ Features

- Add one or more PDF files to the list.  
- Reorder the list (Move Up / Move Down).  
- Remove selected files or clear the entire list.  
- Merge selected files into one combined PDF.  
- Built-in multilingual interface (English, French, German, Spanish, Italian).  
- Integrated documentation.  

## ⚙️ Installation

### 1. Prerequisites
- [Go 1.22+](https://go.dev/dl/) installed.  
- Go will automatically download the required modules on build.

### 2. Clone the project
```bash
git clone https://github.com/Ceramaret-SA/Merge_PDF.git
cd Merge_PDF
```

### 3. Initialize dependencies
```bash
go mod tidy
```

### 4. Build

To create a **Windows portable executable** without the console window:

```powershell
go build -ldflags "-s -w -H=windowsgui" -o MergePDF.exe .
```

To build for **Linux**:

```bash
go build -ldflags "-s -w" -o MergePDF .
```

## 🖥️ Usage

Launch **MergePDF.exe** or **MergePDF**.

Add PDF files using **“Add PDF”**.

Reorder them with **“Move Up”** or **“Move Down”**.

Remove a file or use **“Clear List”** to start fresh.

Click **“Merge”**, choose an output file name, and confirm.

Check the generated PDF file.

## ⚠️ Known Issue

If the **output file** is also present in the list of PDFs to merge:

- The generated file will be **empty and corrupted**.  
- The program will display an **“empty file”** error.

**Solution:** Always specify a **new file name** for the output  
(for example: `merged_result.pdf`).

## 🌍 Language Support

Automatic **system language detection** (Windows and Linux).

**Supported languages:**
- 🇬🇧 English  
- 🇫🇷 French  
- 🇩🇪 German  
- 🇪🇸 Spanish  
- 🇮🇹 Italian  

You can manually switch the language directly from within the application.  
All translations are stored **inside the binary** via Go’s [`embed`](https://pkg.go.dev/embed).

## 📚 Built-in documentation

A **“Documentation”** button in the interface opens a user guide containing:  
- usage steps,  
- common errors and their solutions,  
- a warning about the known problem. 

## 🧩 Internal Structure

```txt
Merge_PDF/
├── README.md
├── go.mod
├── go.sum
├── main.go
└── src/
├── i18n/
│ ├── i18n.go
│ └── lang/
│ │ ├── active.en.json
│ │ ├── active.fr.json
│ │ ├── active.de.json
│ │ ├── active.es.json
│ │ └── active.it.json
├── os_wrappers/
│ └── os_wrappers.go
├── ui/
│ └── ui.go
└── utils/
  └── utils.go
```

## 🧠 Platform Compatibility

This project has been **fully tested and confirmed to work on Windows**.

The application **should also work on Linux**, as it uses cross-platform  
Go and Fyne libraries — however, **Linux compatibility has not yet been tested**.

If you plan to run or build on Linux:
- The PowerShell-based file dialogs are automatically replaced with native Fyne dialogs.
- Language detection will use the `LANG` environment variable.


## 📜 License

- [Fyne](https://fyne.io) – BSD  
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) – Apache 2.0  
- [go-i18n](https://github.com/nicksnyder/go-i18n) – MIT
- This project may be used for **commercial purposes**. 
