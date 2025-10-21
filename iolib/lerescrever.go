package iolib

import (
	"bufio"
	"fmt"
	"grafos/dsa"
	"os"
	"strconv"
	"strings"
)

// LerGrafoDoArquivo lê grafo de arquivo e já faz TrimSpace em tudo
func LerGrafoDoArquivo(nome string) (dsa.Grafo, dsa.No, dsa.No, error) {
	file, err := os.Open(nome)
	if err != nil {
		return nil, "", "", err
	}
	defer file.Close()

	grafo := make(dsa.Grafo)
	var inicio, objetivo dsa.No

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha == "" || strings.HasPrefix(linha, "#") {
			continue
		}

		if strings.HasPrefix(linha, "inicio=") {
			inicio = dsa.No(strings.TrimSpace(strings.TrimPrefix(linha, "inicio=")))
			continue
		}
		if strings.HasPrefix(linha, "fim=") || strings.HasPrefix(linha, "objetivo=") {
			objetivo = dsa.No(strings.TrimSpace(strings.TrimPrefix(linha, "fim=")))
			if strings.HasPrefix(linha, "objetivo=") {
				objetivo = dsa.No(strings.TrimSpace(strings.TrimPrefix(linha, "objetivo=")))
			}
			continue
		}

		// Linha de aresta: origem destino custo
		partes := strings.Fields(linha)
		if len(partes) != 3 {
			continue
		}

		custo, err := strconv.ParseFloat(strings.TrimSpace(partes[2]), 64)
		if err != nil {
			fmt.Println("⚠️  Erro ao converter custo:", partes[2])
			continue
		}

		origem := dsa.No(strings.TrimSpace(partes[0]))
		destino := dsa.No(strings.TrimSpace(partes[1]))

		grafo[origem] = append(grafo[origem], dsa.Aresta{Para: destino, Custo: custo})

		// Se quiser grafo não-direcionado, descomente:
		// grafo[destino] = append(grafo[destino], dsa.Aresta{Para: origem, Custo: custo})
	}

	return grafo, inicio, objetivo, scanner.Err()
}

// SalvarSaida salva um arquivo com o conteúdo passado
func SalvarSaida(caminho, conteudo string) error {
	return os.WriteFile(caminho, []byte(conteudo), 0644)
}

// GerarRelatorio gera saída em texto
func GerarRelatorio(nome string, caminho []dsa.No, custo float64) string {
	var sb strings.Builder
	sb.WriteString("════════════════════════════════════════════\n")
	sb.WriteString(fmt.Sprintf("Exemplo: %s\n", nome))
	sb.WriteString("────────────────────────────────────────────\n")
	if len(caminho) == 0 || custo == 0 || custoEsInfinito(custo) {
		sb.WriteString("Nenhum caminho encontrado.\n")
	} else {
		sb.WriteString(fmt.Sprintf("Caminho encontrado:\n  %v\n", caminho))
		sb.WriteString(fmt.Sprintf("Custo total: %.2f\n", custo))
	}
	sb.WriteString("════════════════════════════════════════════\n")
	return sb.String()
}
