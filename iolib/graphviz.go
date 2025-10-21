package iolib

import (
	"fmt"
	"grafos/dsa"
	"strings"
)

func custoEInfinito(c float64) bool {
	return c != c || c > 1e9
}

func GerarGraphviz(grafo dsa.Grafo, caminho []dsa.No, inicio, objetivo dsa.No, obstaculos map[dsa.No]bool) string {
	var sb strings.Builder
	sb.WriteString("digraph G {\n  rankdir=LR;\n  node [shape=circle];\n")

	// Cores especiais
	sb.WriteString(fmt.Sprintf("  \"%s\" [style=filled, fillcolor=green];\n", inicio))
	sb.WriteString(fmt.Sprintf("  \"%s\" [style=filled, fillcolor=red];\n", objetivo))

	for no := range obstaculos {
		sb.WriteString(fmt.Sprintf("  \"%s\" [style=filled, fillcolor=gray, fontcolor=white];\n", no))
	}

	caminhoArestas := make(map[string]bool)
	for i := 0; i < len(caminho)-1; i++ {
		key := string(caminho[i]) + "->" + string(caminho[i+1])
		caminhoArestas[key] = true
	}

	for origem, arestas := range grafo {
		for _, a := range arestas {
			key := string(origem) + "->" + string(a.Para)
			style := "solid"
			color := "black"

			if caminhoArestas[key] {
				style = "bold"
				color = "green"
			}

			if obstaculos[a.Para] || obstaculos[origem] {
				color = "gray"
				style = "dashed"
			}

			sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [label=\"%.2f\", style=%s, color=%s];\n",
				origem, a.Para, a.Custo, style, color))
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}
