# 🧩 MergePDF

MergePDF is a simple graphical application that allows you to **merge multiple PDF files** into a single document.  
It is developed in **Go**, using [Fyne](https://fyne.io) for the user interface and [pdfcpu](https://github.com/pdfcpu/pdfcpu) for PDF processing.  
The application produces a **standalone, portable binary** that runs without external dependencies.

---

## ✨ Features

- Add one or more PDF files (multi-selection supported).  
- Reorder the list (Move Up / Move Down).  
- Delete one or several files, or clear the entire list.  
- Merge selected files into a single PDF.  
- Integrated documentation and multilingual support.  
- Fully portable executable — no installer needed.

> ⚠️ **Note**:  
> The software is fully portable on **Windows**.  
> If you need to share it on **Linux**, users need to **build it** using the `build.sh` script.  

---

## ⚙️ Installation

### 1. Clone the project
```bash
git clone https://github.com/Ceramaret-SA/Merge_PDF.git
cd Merge_PDF
```

### 2. Build

#### 1. Windows

##### Build automatically
Use the provided PowerShell script:

```powershell
.\build.ps1
```

This script will:
- Check if Go 1.25.1 is installed (install it if missing),
- Install all Go dependencies,
- Compile `MergePDF.exe` as a GUI app (no console window),
- Launch it automatically.

---

#### 2. Linux (Fedora / Ubuntu / Debian)

##### Build automatically
Use the provided Bash script:

```bash
.\build.sh
```

This script will:
- Checks for Go 1.25.1 and installs it if needed (via APT or DNF),
- Installs Fyne dependencies (`libX11-dev`, `libGL`, etc.),
- Compiles the `MergePDF` binary,
- Runs it automatically.

> ⚠️ **Warning**:  
> You need to be connected as an **administrator** (`sudo`) to launch the build script.   

> ⚠️ **Note**:  
> The program is fully functional on **Windows**.  
> On **Linux**, it works correctly on Fedora and Ubuntu, but has not yet been extensively tested on other distributions.  
> Some GUI behaviors (e.g., file dialogs) may vary slightly depending on your desktop environment.  

---

## 🖥️ Usage

Launch **MergePDF.exe** or **MergePDF**.

Add PDF files using **“Add PDF”**.

Reorder them with **“Move Up”** or **“Move Down”**.

Remove a file or use **“Clear List”** to start fresh.

Click **“Merge”**, choose an output file name, and confirm.

Check the generated PDF file.

---

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

---

## 📚 Built-in documentation

A **“Documentation”** button in the interface opens a user guide containing:  
- usage steps,  
- common errors and their solutions,  
- a warning about the known problem. 

---

## 🧩 Internal Structure

```txt
Merge_PDF/
├── README.md
├── go.mod
├── build.ps1
├── build.sh
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
| ├── os_wrappers_linux.go
│ └── os_wrappers_windows.go
├── ui/
│ └── ui.go
└── utils/
  └── utils.go
```

---

## 🧠 Platform Compatibility

This project has been **fully tested and confirmed to work on Windows**.

The application **should also work on Linux**, as it uses cross-platform  
Go and Fyne libraries — however, **Linux compatibility has not yet been tested**.

If you plan to run or build on **Linux**:
- The PowerShell-based file dialogs **are automatically** replaced with native Fyne dialogs.
- Language detection will use the `LANG` environment variable.

---

## 🧱 Compatibility

| Platform | Status | Notes |
|-----------|---------|-------|
| Windows 10/11 | ✅ Stable | Fully portable (`MergePDF.exe`) |
| Linux (Fedora / Ubuntu) | ⚠️ Working | Tested; minor GUI variations possible |

---

## 📜 License

- [Fyne](https://fyne.io) – BSD  
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) – Apache 2.0  
- [go-i18n](https://github.com/nicksnyder/go-i18n) – MIT
- This project may be used for **commercial purposes**. 
