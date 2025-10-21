package main

import (
	"fmt"
	"grafos/dsa"
	"grafos/iolib"
	"os"
	"path/filepath"
)

func main() {
	dirBase := "exemplos"
	pastas, err := os.ReadDir(dirBase)
	if err != nil {
		fmt.Println("Erro ao ler diretório base:", err)
		return
	}

	for _, pasta := range pastas {
		if !pasta.IsDir() {
			continue
		}

		caminhoPasta := filepath.Join(dirBase, pasta.Name())
		entrada := filepath.Join(caminhoPasta, "entrada.txt")
		saidaTxt := filepath.Join(caminhoPasta, "saida.txt")
		saidaDot := filepath.Join(caminhoPasta, "saida.dot")

		grafo, inicio, objetivo, err := iolib.LerGrafoDoArquivo(entrada)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		if inicio == "" || objetivo == "" {
			fmt.Printf("Falha: início ou objetivo vazio na pasta '%s'\n", pasta.Name())
			continue
		}

		fmt.Printf("Processando '%s' → Início: %s | Objetivo: %s\n", pasta.Name(), inicio, objetivo)

		caminho, custo := dsa.AEstrela(grafo, inicio, objetivo, func(a, b dsa.No) float64 { return 0 })

		err = iolib.SalvarSaida(saidaTxt, iolib.GerarRelatorio(pasta.Name(), caminho, custo))
		if err != nil {
			fmt.Println("Erro ao salvar TXT:", err)
		}

		err = iolib.SalvarSaida(saidaDot, iolib.GerarGraphviz(grafo, caminho, inicio, objetivo))
		if err != nil {
			fmt.Println("Erro ao salvar DOT:", err)
		}

		fmt.Printf("Exemplo '%s' processado.\n", pasta.Name())
	}
}
