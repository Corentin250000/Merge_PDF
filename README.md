# MergePDF

MergePDF is a lightweight graphical application that allows you to **merge multiple PDF files** into a single document.  
It is developed in **Go**, using [Fyne](https://fyne.io) for the interface and [pdfcpu](https://github.com/pdfcpu/pdfcpu) for PDF processing.  

The application produces a **single, portable executable** — no installation or external files required.

## ✨ Features

- Add one or more PDF files at once.  
- Reorder the list with **Move Up** / **Move Down**.  
- Remove selected files or **Clear the entire list**.  
- Merge all selected files into a single PDF document.  
- Choose the **output filename** freely.  
- Built-in **documentation** button for help.  
- **Automatic language detection** (English, French, German, Spanish, Italian).  
- **Manual language selector** in the interface.  
- All language files are **embedded inside the executable** (no external dependencies).

## 📥 Installation

### 1. Prerequisites
- [Go 1.22+](https://go.dev/dl/) installed if you want to build the binary yourself.  
- No prerequisites required to **run** the `.exe` file.

### 2. Clone the project
```bash
git clone https://github.com/Ceramaret-SA/Merge_PDF.git
cd MergePDF
```

### 3. Initialize dependencies
```bash
go mod tidy
```

### 4. Build

```bash
go build -ldflags "-s -w -H=windowsgui" -o MergePDF.exe main.go
```

This will generate a **standalone `MergePDF.exe`** containing all resources (translations and interface).  
It can be copied and executed on any Windows system — **no `lang/` folder needed**.

## 🖥️ Usage

1. Launch **MergePDF.exe**.  
2. Add one or more PDF files using **“Add PDF(s)”**.  
3. Reorder them with **“Move Up”** and **“Move Down”**.  
4. Remove selected files or click **“Clear List”** to start over.  
5. Click **“Merge”** and choose a name for the merged file.  
6. Check the resulting file to confirm the merge.

## ⚠️ Known issue

If you select as **output file** a PDF that is already present in the merge list:

- the generated file will be **empty and corrupted**,  
- the program will report that the file is empty.

👉 **Solution:** Always choose a **different name** for the output file (e.g. `merged_result.pdf`).

## 🌍 Language support

- Automatic detection of your system language (Windows).  
- Supported languages:  
  - English 🇬🇧  
  - French 🇫🇷  
  - German 🇩🇪  
  - Spanish 🇪🇸  
  - Italian 🇮🇹  
- A **language selector** in the interface lets you switch manually at any time.  
- All translations are stored **inside the binary** via Go’s [`embed`](https://pkg.go.dev/embed).

## 📚 Built-in documentation

A **“Documentation”** button in the interface opens a user guide containing:  
- usage steps,  
- common errors and their solutions,  
- a warning about the known problem.  


## 📜 License

- [Fyne](https://fyne.io) – BSD  
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) – Apache 2.0  
- [https://github.com/nicksnyder/go-i18n](go-i18n) – MIT
- This project may be used for **commercial purposes**. 
