package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// ============================================================
// ✏️  CHANGE THESE VALUES TO MATCH YOUR SETUP
// ============================================================
const inputFolder  = "../images"     // 📁 folder containing your images
const outputFolder = "output"        // 📁 folder where .txt files will be saved (auto-created)
const width        = 100             // output width in characters
const mode         = "blend"         // "blend" | "edge" | "brightness"
const gamma        = 1.8             // contrast boost (try 1.5–2.5)
const threshold    = 0.04            // noise cut-off (try 0.02–0.10)
const ramp         = " _*!~)(+^#&$%@" // character ramp light→dark
// ============================================================

// Supported image extensions
var supportedExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

// Sobel kernels for edge detection
var sobelX = [3][3]float64{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}
var sobelY = [3][3]float64{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
}

func main() {
	// Create output folder if it doesn't exist
	if err := os.MkdirAll(outputFolder, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Cannot create output folder: %v\n", err)
		os.Exit(1)
	}

	// Read all files in input folder
	entries, err := os.ReadDir(inputFolder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Cannot open input folder '%s': %v\n", inputFolder, err)
		os.Exit(1)
	}

	// Filter only supported image files
	var imageFiles []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if supportedExts[ext] {
			imageFiles = append(imageFiles, e.Name())
		}
	}

	if len(imageFiles) == 0 {
		fmt.Fprintf(os.Stderr, "⚠️  No images found in '%s'\n", inputFolder)
		fmt.Fprintf(os.Stderr, "    Supported formats: .jpg .jpeg .png .gif\n")
		os.Exit(1)
	}

	fmt.Printf("🖼️  Found %d image(s) in '%s'\n", len(imageFiles), inputFolder)
	fmt.Printf("📂 Output will be saved to '%s/'\n\n", outputFolder)

	success := 0
	failed  := 0

	for i, filename := range imageFiles {
		inputPath  := filepath.Join(inputFolder, filename)
		baseName   := strings.TrimSuffix(filename, filepath.Ext(filename))
		outputPath := filepath.Join(outputFolder, baseName+".txt")

		fmt.Printf("[%d/%d] Converting: %-30s → %s ... ",
			i+1, len(imageFiles), filename, outputPath)

		err := convertImage(inputPath, outputPath)
		if err != nil {
			fmt.Printf("❌ FAILED (%v)\n", err)
			failed++
		} else {
			fmt.Printf("✅ Done\n")
			success++
		}
	}

	fmt.Printf("\n========================================\n")
	fmt.Printf("✅ Success : %d\n", success)
	if failed > 0 {
		fmt.Printf("❌ Failed  : %d\n", failed)
	}
	fmt.Printf("📁 All outputs saved in: ./%s/\n", outputFolder)
}

// convertImage converts a single image file to ASCII and saves as .txt
func convertImage(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("cannot open: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("cannot decode: %w", err)
	}

	bounds := img.Bounds()
	imgW   := bounds.Max.X - bounds.Min.X
	imgH   := bounds.Max.Y - bounds.Min.Y

	outW := width
	outH := int(math.Round(float64(outW) * float64(imgH) / float64(imgW) * 0.45))
	if outH < 1 {
		outH = 1
	}

	gray := sampleGray(img, outW, outH)

	var intensityMap [][]float64
	switch mode {
	case "edge":
		intensityMap = edgeDetect(gray, outW, outH)
	case "brightness":
		intensityMap = brightnessMap(gray, outW, outH)
	default: // blend
		edges := edgeDetect(gray, outW, outH)
		bright := brightnessMap(gray, outW, outH)
		intensityMap = blend(edges, bright, outW, outH, 0.7)
	}

	applyGamma(intensityMap, outW, outH, gamma)
	result := mapToChars(intensityMap, outW, outH, ramp, threshold)

	return os.WriteFile(outputPath, []byte(result), 0644)
}

// sampleGray samples the image at outW x outH resolution → [row][col] in [0,1]
func sampleGray(img image.Image, outW, outH int) [][]float64 {
	bounds := img.Bounds()
	imgW   := bounds.Max.X - bounds.Min.X
	imgH   := bounds.Max.Y - bounds.Min.Y

	g := make([][]float64, outH)
	for row := range g {
		g[row] = make([]float64, outW)
		for col := range g[row] {
			srcX := bounds.Min.X + int(float64(col)/float64(outW)*float64(imgW))
			srcY := bounds.Min.Y + int(float64(row)/float64(outH)*float64(imgH))
			if srcX >= bounds.Max.X { srcX = bounds.Max.X - 1 }
			if srcY >= bounds.Max.Y { srcY = bounds.Max.Y - 1 }
			g[row][col] = toLuma(img.At(srcX, srcY))
		}
	}
	return g
}

// toLuma converts a color to perceptual luminance [0,1]
func toLuma(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	rf := float64(r>>8) / 255.0
	gf := float64(g>>8) / 255.0
	bf := float64(b>>8) / 255.0
	return 0.299*rf + 0.587*gf + 0.114*bf
}

// edgeDetect applies a Sobel operator; returns normalized edge magnitude [0,1]
func edgeDetect(gray [][]float64, w, h int) [][]float64 {
	raw := make([][]float64, h)
	for i := range raw {
		raw[i] = make([]float64, w)
	}

	maxVal := 0.0
	for row := 1; row < h-1; row++ {
		for col := 1; col < w-1; col++ {
			var gx, gy float64
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					px := gray[row+ky][col+kx]
					gx += sobelX[ky+1][kx+1] * px
					gy += sobelY[ky+1][kx+1] * px
				}
			}
			mag := math.Sqrt(gx*gx + gy*gy)
			raw[row][col] = mag
			if mag > maxVal {
				maxVal = mag
			}
		}
	}

	out := make([][]float64, h)
	for row := range raw {
		out[row] = make([]float64, w)
		for col, v := range raw[row] {
			if maxVal > 0 {
				out[row][col] = v / maxVal
			}
		}
	}
	return out
}

// brightnessMap: dark pixel = high value (dark lines → characters)
func brightnessMap(gray [][]float64, w, h int) [][]float64 {
	out := make([][]float64, h)
	for row := range gray {
		out[row] = make([]float64, w)
		for col, v := range gray[row] {
			out[row][col] = 1.0 - v
		}
	}
	return out
}

// blend mixes edge + brightness maps (70% edge, 30% brightness)
func blend(edges, bright [][]float64, w, h int, edgeWeight float64) [][]float64 {
	out := make([][]float64, h)
	bw  := 1.0 - edgeWeight
	for row := range edges {
		out[row] = make([]float64, w)
		for col := range edges[row] {
			out[row][col] = edgeWeight*edges[row][col] + bw*bright[row][col]
		}
	}
	return out
}

// applyGamma boosts contrast using power-law curve
func applyGamma(m [][]float64, w, h int, gamma float64) {
	for row := range m {
		for col, v := range m[row] {
			m[row][col] = math.Pow(v, 1.0/gamma)
		}
	}
}

// mapToChars maps intensity values to ASCII characters
func mapToChars(m [][]float64, w, h int, ramp string, threshold float64) string {
	n  := len(ramp)
	var sb strings.Builder
	sb.Grow(w*h + h)

	for _, row := range m {
		for _, v := range row {
			if v < threshold {
				sb.WriteByte(' ')
				continue
			}
			idx := int(v * float64(n-1))
			if idx >= n { idx = n - 1 }
			sb.WriteByte(ramp[idx])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}