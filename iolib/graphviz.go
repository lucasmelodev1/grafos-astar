package iolib

import (
	"fmt"
	"grafos/dsa"
	"strings"
)

func custoEsInfinito(c float64) bool {
	return c != c || c > 1e9 // NaN ou infinito
}

func GerarGraphviz(grafo dsa.Grafo, caminho []dsa.No, inicio, objetivo dsa.No) string {
	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	sb.WriteString("  rankdir=LR;\n")
	sb.WriteString("  node [shape=circle];\n")

	// Destaca início e objetivo
	sb.WriteString(fmt.Sprintf("  \"%s\" [style=filled, fillcolor=green];\n", inicio))
	sb.WriteString(fmt.Sprintf("  \"%s\" [style=filled, fillcolor=red];\n", objetivo))

	// Cria um mapa rápido das arestas do caminho
	caminhoArestas := make(map[string]bool)
	for i := 0; i < len(caminho)-1; i++ {
		key := string(caminho[i]) + "->" + string(caminho[i+1])
		caminhoArestas[key] = true
	}

	// Cria arestas
	for origem, arestas := range grafo {
		for _, a := range arestas {
			key := string(origem) + "->" + string(a.Para)
			style := "solid"
			color := "black"

			if caminhoArestas[key] {
				style = "bold"
				color = "green"
			}

			sb.WriteString(fmt.Sprintf(
				"  \"%s\" -> \"%s\" [label=\"%.2f\", style=%s, color=%s];\n",
				origem, a.Para, a.Custo, style, color,
			))
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}
