package dsa

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type No string

type Aresta struct {
	Para  No
	Custo float64
}

type Grafo map[No][]Aresta

// gerarSaidaVisual cria um arquivo .txt no caminho especificado com o resultado do algoritmo A*.
// Ela mostra o grafo, o caminho encontrado (se existir) e o custo total do trajeto.
func GerarSaidaVisual(nome string, grafo Grafo, inicio, objetivo No, caminho []No, custo float64) error {
	var b strings.Builder

	b.WriteString("========================================\n")
	b.WriteString(" ALGORITMO A* - RESULTADO DO PROCESSO\n")
	b.WriteString("========================================\n\n")

	b.WriteString(fmt.Sprintf("Arquivo de saída: %s\n", nome))
	b.WriteString(fmt.Sprintf("Nó inicial: %s\n", inicio))
	b.WriteString(fmt.Sprintf("Nó objetivo: %s\n\n", objetivo))

	b.WriteString("Grafo:\n")
	for origem, arestas := range grafo {
		for _, a := range arestas {
			b.WriteString(fmt.Sprintf("  %s -> %s (custo: %.2f)\n", origem, a.Para, a.Custo))
		}
	}
	b.WriteString("\n----------------------------------------\n")

	if math.IsInf(custo, 1) {
		b.WriteString("Nenhum caminho encontrado entre os nós.\n")
	} else {
		b.WriteString("Caminho encontrado:\n")
		for i, n := range caminho {
			if i == len(caminho)-1 {
				b.WriteString(fmt.Sprintf("  %s\n", n))
			} else {
				b.WriteString(fmt.Sprintf("  %s → ", n))
			}
		}
		b.WriteString(fmt.Sprintf("\n\nCusto total: %.2f\n", custo))
	}

	b.WriteString("----------------------------------------\n")

	return os.WriteFile(nome, []byte(b.String()), 0644)
}
