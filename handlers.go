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
	info := `{
		"nome": "API de Conversão de Imagens",
		"descricao": "API para conversão de imagens em diferentes formatos",
		"versao": "1.0.0"
		"autor": "Fabricio Roney de Amorim"
		"linkedin": "https://www.linkedin.com/in/fabricio-roney/"
		"github": "https://github.com/fabricioramorim"
		"caminhos": {"info": "/info", "convert": "/convert"}
	}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(info))
}

func convertImage(w http.ResponseWriter, r *http.Request) {
	// Lê a imagem enviada
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Erro ao ler a imagem", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decodifica a imagem
	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Erro ao decodificar a imagem", http.StatusBadRequest)
		return
	}

	// Converte para PNG (ou outro formato desejado)
	newImg := imaging.Clone(img)
	tempFile, err := os.CreateTemp("", "converted_image.*.png")
	if err != nil {
		http.Error(w, "Erro ao criar arquivo temporário", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	err = png.Encode(tempFile, newImg)
	if err != nil {
		http.Error(w, "Erro ao converter a imagem", http.StatusInternalServerError)
		return
	}

	// Envia a imagem convertida como resposta
	tempFile.Seek(0, 0)
	w.Header().Set("Content-Type", "image/png")
	http.ServeContent(w, r, "converted_image.png", time.Now(), tempFile)
}
