package main

import (
	"image"
	"image/png"
	"net/http"
	"os"
	"time"

	"github.com/disintegration/imaging"
)

func getInfoMe(w http.ResponseWriter, r *http.Request) {
	// json com as informações da api
	info := `
	{
		"name": "API de Conversão de Imagens",
		"description": "API para conversão de imagens em diferentes formatos",
		"version": "1.0.0"
		"autor": "Fabricio Roney de Amorim"
		"linkedin": "https://www.linkedin.com/in/fabricio-roney/"
		"github": "https://github.com/fabricioramorim"
		"handlers": 
			{
			"/info": "This handler, describe all about the API", 
			"/convert": 
				{
				"/webp": "Convert source image to WEBP"
				}
			}
	}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(info))
}

func convertImage(w http.ResponseWriter, r *http.Request) {
	info := `
	{
		"Error": "Invalid Call, format is missing",
	}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(info))
}

func convertImageWebp(w http.ResponseWriter, r *http.Request) {
	// Read
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Erro ao ler a imagem", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decoder
	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Erro ao decodificar a imagem", http.StatusBadRequest)
		return
	}

	// Convert
	newImg := imaging.Clone(img)
	tempFile, err := os.CreateTemp("", "converted_image.*.webp")
	if err != nil {
		http.Error(w, "Erro ao criar arquivo temporário", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	err = webp.Encode(tempFile, newImg)
	if err != nil {
		http.Error(w, "Erro ao converter a imagem", http.StatusInternalServerError)
		return
	}

	// Send
	tempFile.Seek(0, 0)
	w.Header().Set("Content-Type", "image/webp")
	http.ServeContent(w, r, "converted_image.webp", time.Now(), tempFile)
}
