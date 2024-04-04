package utils

import (
	"encoding/json"
	"fmt"
	"net"
)

type Mensagem struct {
	L1        string `json:"l1"`
	L1Fnt     string `json:"l1_fnt"`
	L1FntSize string `json:"l1_fnt_size"`
	L1Y       string `json:"l1_y"`
	L2        string `json:"l2"`
	L2Fnt     string `json:"l2_fnt"`
	L2FntSize string `json:"l2_fnt_size"`
	L2Y       string `json:"l2_y"`
	L3        string `json:"l3"`
	L3Fnt     string `json:"l3_fnt"`
	L3FntSize string `json:"l3_fnt_size"`
	L3Y       string `json:"l3_y"`
	L4        string `json:"l4"`
	L4Fnt     string `json:"l4_fnt"`
	L4FntSize string `json:"l4_fnt_size"`
	L4Y       string `json:"l4_y"`
	Cluster   string `json:"cluster"`
	Porta     string `json:"porta"`
}

func SefazMs() {
	// Dados para o JSON
	dadosJSON := Mensagem{
		L1:        "IMP-2773",
		L1Fnt:     "5",
		L1FntSize: "16",
		L1Y:       "1",
		L2:        "PARADA NO",
		L2Fnt:     "5",
		L2FntSize: "16",
		L2Y:       "15",
		L3:        "  PÁTIO     .",
		L3Fnt:     "5",
		L3FntSize: "16",
		L3Y:       "29",
		L4:        "     T",
		L4Fnt:     "8",
		L4FntSize: "20",
		L4Y:       "29",
		Cluster:   "10.5.28.49",
		Porta:     "4001",
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(dadosJSON)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	// Endereço de destino
	destinoAddr, err := net.ResolveUDPAddr("udp", "192.168.53.236:1111")
	if err != nil {
		fmt.Println("Erro ao resolver endereço:", err)
		return
	}

	// Cria uma conexão UDP
	conn, err := net.DialUDP("udp", nil, destinoAddr)
	if err != nil {
		fmt.Println("Erro ao criar conexão UDP:", err)
		return
	}
	defer conn.Close()

	// Envia os dados
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Erro ao enviar dados:", err)
		return
	}

	fmt.Println("JSON enviado com sucesso para 192.168.53.236:1111")
}
