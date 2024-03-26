package examples

//
// import (
// 	"bytes"
// 	"compress/gzip"
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"
// )
//
// const path = "/var/data/output.json.gz" // Altere a extensão do arquivo para .gz
// const imagePath = "/var/data/image.png"
//
// // JSONData representa a estrutura dos dados JSON
// type JSONData struct {
// 	Result struct {
// 		Lum    int    `json:"lum"`
// 		Dthr   string `json:"dthr"`
// 		Octeto string `json:"octeto"`
// 		Image  []byte `json:"image"`
// 	} `json:"result"`
// 	Error interface{} `json:"error"`
// 	ID    int         `json:"id"`
// }
//
// func TestSnmp2() {
// 	generateAndSaveFile()
// 	fmt.Println("Aplicação rodando em segundo plano...")
//
// }
//
// func generateAndSaveFile() {
// 	lum := rand.Intn(101)
// 	imageData, err := getImage()
// 	if err != nil {
// 		fmt.Printf("Erro ao obter a imagem: %v\n", err)
// 		return
// 	}
//
// 	jsonData := JSONData{
// 		Result: struct {
// 			Lum    int    `json:"lum"`
// 			Dthr   string `json:"dthr"`
// 			Octeto string `json:"octeto"`
// 			Image  []byte `json:"image"`
// 		}{
// 			Lum:    lum,
// 			Dthr:   time.Now().Format("2006-01-02 15:04:05"),
// 			Octeto: "ff",
// 			Image:  imageData,
// 		},
// 		Error: nil,
// 		ID:    rand.Intn(899999) + 100000,
// 	}
//
// 	// Converte a estrutura para JSON
// 	jsonBytes, err := json.Marshal(jsonData)
// 	if err != nil {
// 		fmt.Printf("Erro ao converter para JSON: %v\n", err)
// 		return
// 	}
//
// 	// Compacta o JSON usando gzip
// 	compressedData, err := compressData(jsonBytes)
// 	if err != nil {
// 		fmt.Printf("Erro ao compactar os dados: %v\n", err)
// 		return
// 	}
//
// 	// Salva o JSON compactado em um arquivo
// 	err = saveDataToFile(compressedData, path)
// 	if err != nil {
// 		fmt.Printf("Erro ao salvar o arquivo: %v\n", err)
// 	}
//
// 	fmt.Println("Arquivo salvo com sucesso!")
// }
//
// func getImage() ([]byte, error) {
// 	imageData, err := os.ReadFile(imagePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return imageData, nil
// }
//
// func compressData(data []byte) ([]byte, error) {
// 	var buf bytes.Buffer
// 	writer := gzip.NewWriter(&buf)
// 	_, err := writer.Write(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }
//
// func saveDataToFile(data []byte, filename string) error {
// 	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
//
// 	_, err = file.Write(data)
// 	return err
// }
// //
