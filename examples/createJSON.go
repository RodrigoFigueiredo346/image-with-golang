package examples

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const path = "/var/data/output.json"
const imagePath = "/var/data/image.png"

// JSONData representa a estrutura dos dados JSON
type JSONData struct {
	Result struct {
		Lum    int    `json:"lum"`
		Dthr   string `json:"dthr"`
		Octeto string `json:"octeto"`
		Image  []byte `json:"image"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}

func TestSnmp2() {

	generateAndSaveFile()

	fmt.Println("Aplicação rodando em segundo plano...")

	select {}

}

func generateAndSaveFile() {

	lum := rand.Intn(101)

	imageData, err := getImage()
	if err != nil {
		fmt.Printf("Erro ao obter a imagem: %v\n", err)

	}

	jsonData := JSONData{
		Result: struct {
			Lum    int    `json:"lum"`
			Dthr   string `json:"dthr"`
			Octeto string `json:"octeto"`
			Image  []byte `json:"image"`
		}{
			Lum:    lum,
			Dthr:   time.Now().Format("2006-01-02 15:04:05"),
			Octeto: "ff",
			Image:  imageData,
		},
		Error: nil,
		ID:    rand.Intn(899999) + 100000,
	}

	// Converte a estrutura para JSON
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v\n", err)

	}

	// Salva o JSON em um arquivo, substituindo o conteúdo existente
	err = saveJSONToFile(jsonBytes, path)
	if err != nil {
		fmt.Printf("Erro ao salvar o arquivo: %v\n", err)
	}

	// Aguarda 15 segundos
	//time.Sleep(15 * time.Second)

}

func getImage() ([]byte, error) {
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	//fmt.Println("imageData == ", imageData)
	return imageData, nil
}

func saveJSONToFile(data []byte, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
