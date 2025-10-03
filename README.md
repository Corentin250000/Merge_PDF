# FusionPDF

FusionPDF est une application graphique simple permettant de **fusionner plusieurs fichiers PDF** en un seul document.  
Elle est développée en **Go**, avec [Fyne](https://fyne.io) pour l’interface et [pdfcpu](https://github.com/pdfcpu/pdfcpu) pour la gestion des fichiers PDF.  

L’application génère un **binaire unique et portable**, utilisable sans dépendances externes.

---

## ✨ Fonctionnalités

- Ajouter un ou plusieurs fichiers PDF.  
- Réorganiser l’ordre (Monter / Descendre).  
- Supprimer un fichier ou effacer toute la liste.  
- Fusionner les fichiers sélectionnés en un seul PDF.  
- Documentation intégrée accessible via un bouton.  

---

## 📥 Installation

### 1. Prérequis
- [Go 1.22+](https://go.dev/dl/) installé.  
- Les modules Go seront téléchargés automatiquement (`fyne`, `pdfcpu`).  

### 2. Cloner le projet
```bash
https://github.com/Ceramaret-SA/Merge_PDF.git
cd FusionPDF
```

### 3. Initialiser les dépendances
```bash
go mod tidy
```

### 4. Compilation

    ```bash
    go build -ldflags "-s -w -H=windowsgui" -o FusionPDF.exe main.go
    ```

Le binaire obtenu (`FusionPDF.exe`) est portable et peut être utilisé **sans installer Go ni d’autres dépendances**.

## 🖥️ Utilisation

1. Lancer **FusionPDF.exe**.  
2. Ajouter des fichiers PDF avec le bouton **« Ajouter PDF »**.  
3. Réorganiser l’ordre avec **« Monter »** et **« Descendre »**.  
4. Supprimer un fichier ou utiliser **« Effacer la liste »** si nécessaire.  
5. Cliquer sur **« Fusionner »** et choisir le nom du fichier de sortie.  
6. Ouvrir le fichier final pour vérifier le résultat.  

---

## ⚠️ Bug connu

Si vous choisissez comme **fichier de sortie** un PDF qui est déjà présent dans la liste des fichiers à fusionner :  
- le fichier généré sera **vide et corrompu**,  
- le programme indiquera que le fichier est vide.  

👉 **Solution :** toujours donner un **nouveau nom** au fichier de sortie (exemple : `fusion_result.pdf`).  

---

## 📚 Documentation intégrée

Un bouton **« Documentation »** dans l’interface ouvre une aide utilisateur avec :  
- les étapes d’utilisation,  
- les erreurs fréquentes et leur solution,  
- l’avertissement concernant le bug connu.  

---

## 📜 Licence

- [Fyne](https://fyne.io) – BSD  
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) – Apache 2.0  
- Ce projet peut être utilisé dans un cadre commercial.  
