# 🎨 Go-ASCII-Converter

Convert a bunch of images into **ASCII Art** using **Go (Golang)**.

This project scans a folder containing images and converts them into `.txt` ASCII art files automatically. It supports multiple image formats and processes all images in one go.

🚀 This is **Project #1** under my **Golang Projects** learning journey.

---

## 📌 Features

✅ Convert multiple images to ASCII art automatically  
✅ Batch processing for all images in a folder  
✅ Supports multiple image formats  
✅ Auto-generates output `.txt` files  
✅ Clean terminal progress tracking  
✅ Beginner-friendly Go project  

---

## 🖼️ Supported Formats

The converter supports:

- `.jpg`
- `.jpeg`
- `.png`
- `.gif`

---

## 📁 Folder Structure

```bash
📁 your-project/
   ├── image_to_ascii.go
   ├── images/               ← put ALL your images here
   │   ├── photo1.jpg
   │   ├── photo2.png
   │   ├── photo3.jpeg
   │   └── logo.gif
   └── output/               ← auto-created, results appear here
       ├── photo1.txt
       ├── photo2.txt
       ├── photo3.txt
       └── logo.txt
exit
```
---

## ⚙️ Configuration

Only change these 2 lines at the top if you want to use custom folders:

```bash
const inputFolder  = "images"    // 📁 folder with all your images
const outputFolder = "output"    // 📁 where all .txt files get saved
exit
```

---

## ▶️ Run the Project

Navigate to the project folder and run:

```bash
go run image_to_ascii.go
exit
```

---

## 📟 Example Terminal Output

```bash
🖼️  Found 4 image(s) in 'images'
📂 Output will be saved to 'output/'

[1/4] Converting: photo1.jpg     → output/photo1.txt ... ✅ Done
[2/4] Converting: photo2.png     → output/photo2.txt ... ✅ Done
[3/4] Converting: photo3.jpeg    → output/photo3.txt ... ✅ Done
[4/4] Converting: logo.gif       → output/logo.txt   ... ✅ Done

========================================
✅ Success : 4
📁 All outputs saved in: ./output/
exit
```

---

##🧩 How It Works

Put all your images inside the images/ folder
Run the Go file
The program automatically scans all images
Converts each image into ASCII art
Saves the output as .txt files inside the output/ folder

---

##🛠️ Tech Stack

Language: Go (Golang)
Libraries Used: Standard Go libraries only
External Packages: None

--- 

##🎯 Learning Goal

This project is part of my Golang Learning Journey, where I build practical beginner-friendly projects to strengthen my understanding of Go through hands-on development.

---

## ⭐ Support

If you found this project helpful, consider giving it a star ⭐ on GitHub.