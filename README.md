# FusionPDF

FusionPDF is a simple graphical application that allows you to **merge multiple PDF files** into a single document.  
It is developed in **Go**, using [Fyne](https://fyne.io) for the interface and [pdfcpu](https://github.com/pdfcpu/pdfcpu) for PDF file processing.  

The application produces a **single, portable binary** that can run without external dependencies.

---

## ✨ Features

- Add one or more PDF files.  
- Reorder the list (Move Up / Move Down).  
- Delete a file or clear the entire list.  
- Merge selected files into a single PDF.  
- Built-in documentation accessible via a button.  

---

## 📥 Installation

### 1. Prerequisites
- [Go 1.22+](https://go.dev/dl/) installed.  
- Go modules (`fyne`, `pdfcpu`) will be downloaded automatically.  

### 2. Clone the project
```bash
https://github.com/Ceramaret-SA/Merge_PDF.git
cd FusionPDF
```

### 3. Initialize dependencies
```bash
go mod tidy
```

### 4. Build

```bash
go build -ldflags "-s -w -H=windowsgui" -o FusionPDF.exe main.go
```

The resulting binary (`FusionPDF.exe`) is portable and can be used **without installing Go or any other dependencies**.

## 🖥️ Usage

1. Launch **FusionPDF.exe**.  
2. Add PDF files using the **“Add PDF”** button.  
3. Reorder the list with **“Move Up”** and **“Move Down”**.  
4. Remove a file or click **“Clear List”** if needed.  
5. Click **“Merge”** and choose the name of the output file.  
6. Open the resulting file to verify the result.  

---

## ⚠️ Known problem

If you select as **output file** a PDF that is already in the list of files to be merged:  
- the generated file will be **empty and corrupted**,  
- the program will indicate that the file is empty.  

👉 **Solution:** always provide a **new file name** for the output file (e.g. `fusion_result.pdf`).  

---

## ⚠️ Language notice

The application interface is currently available **only in French**.  
An **English version** will be added in a future release.  

---

## 📚 Built-in documentation

A **“Documentation”** button in the interface opens a user guide containing:  
- usage steps,  
- common errors and their solutions,  
- a warning about the known problem.  

---

## 📜 License

- [Fyne](https://fyne.io) – BSD  
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) – Apache 2.0  
- This project may be used for commercial purposes. 


