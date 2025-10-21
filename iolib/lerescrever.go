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
func LerGrafoDoArquivo(nome string) (dsa.Grafo, dsa.No, dsa.No, map[dsa.No]bool, error) {
	file, err := os.Open(nome)
	if err != nil {
		return nil, "", "", nil, err
	}
	defer file.Close()

	grafo := make(dsa.Grafo)
	var inicio, objetivo dsa.No
	obstaculos := make(map[dsa.No]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha == "" || strings.HasPrefix(linha, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(linha, "inicio="):
			inicio = dsa.No(strings.TrimSpace(strings.TrimPrefix(linha, "inicio=")))
		case strings.HasPrefix(linha, "fim=") || strings.HasPrefix(linha, "objetivo="):
			val := strings.SplitN(linha, "=", 2)[1]
			objetivo = dsa.No(strings.TrimSpace(val))
		case strings.HasPrefix(linha, "obstaculo="):
			val := strings.SplitN(linha, "=", 2)[1]
			obstaculos[dsa.No(strings.TrimSpace(val))] = true
		default:
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
		}
	}

	return grafo, inicio, objetivo, obstaculos, scanner.Err()
}

func SalvarSaida(caminho, conteudo string) error {
	return os.WriteFile(caminho, []byte(conteudo), 0644)
}

func GerarRelatorio(nome string, caminho []dsa.No, custo float64) string {
	var sb strings.Builder
	sb.WriteString("════════════════════════════════════════════\n")
	sb.WriteString(fmt.Sprintf("Exemplo: %s\n", nome))
	sb.WriteString("────────────────────────────────────────────\n")
	if len(caminho) == 0 || custo == 0 || custoEInfinito(custo) {
		sb.WriteString("Nenhum caminho encontrado.\n")
	} else {
		sb.WriteString(fmt.Sprintf("Caminho encontrado:\n  %v\n", caminho))
		sb.WriteString(fmt.Sprintf("Custo total: %.2f\n", custo))
	}
	sb.WriteString("════════════════════════════════════════════\n")
	return sb.String()
}
