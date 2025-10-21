package dsa

import (
	"container/heap"
	"math"
)

type FuncaoHeuristica func(n, objetivo No) float64

type NoAEstrela struct {
	Nome   No
	FScore float64
	Indice int
}

func reconstruirCaminho(veioDe map[No]No, atual No) []No {
	caminhoTotal := []No{atual}
	for {
		anterior, ok := veioDe[atual]
		if !ok {
			break
		}
		caminhoTotal = append([]No{anterior}, caminhoTotal...)
		atual = anterior
	}
	return caminhoTotal
}

func AEstrela(grafo Grafo, inicio, objetivo No, h FuncaoHeuristica) ([]No, float64) {
	// Inicializa a Fila de Prioridade (Open Set), que armazena os nós a serem explorados.
	fp := make(FilaDePrioridade, 0)
	heap.Init(&fp)

	// Mapa para rastrear nós na Fila de Prioridade, permitindo acesso e atualização rápidos.
	noOpenSet := make(map[No]*NoAEstrela)

	// Mapa para reconstruir o caminho: armazena de onde cada nó foi alcançado.
	veioDe := make(map[No]No)

	// gScore: Custo do caminho mais barato encontrado de 'inicio' até este nó.
	gScore := make(map[No]float64)
	gScore[inicio] = 0.0 // O custo para o nó inicial é zero.

	// fScore: Custo total estimado (gScore + custo heurístico h).
	fScore := make(map[No]float64)
	fScore[inicio] = h(inicio, objetivo) // fScore inicial = g(inicio) + h(inicio, objetivo)

	// Prepara e adiciona o nó inicial à Fila de Prioridade e ao Open Set.
	noInicio := &NoAEstrela{
		Nome:   inicio,
		FScore: fScore[inicio],
	}
	heap.Push(&fp, noInicio)
	noOpenSet[inicio] = noInicio

	// Loop principal: continua enquanto houver nós a serem explorados.
	for fp.Len() > 0 {
		// Pega e remove o nó com o menor fScore do Open Set. Este é o 'nó atual'.
		noAtual := heap.Pop(&fp).(*NoAEstrela)
		atual := noAtual.Nome
		delete(noOpenSet, atual) // Move o nó para o 'Closed Set' implicitamente.

		// Condição de parada: Se o nó atual for o objetivo, reconstrua e retorne o caminho.
		if atual == objetivo {
			return reconstruirCaminho(veioDe, atual), gScore[atual]
		}

		// Explora os vizinhos do nó atual.
		for _, aresta := range grafo[atual] {
			vizinho := aresta.Para

			// Calcula o gScore (custo do caminho) se for por este novo percurso.
			gScoreTentativo := gScore[atual] + aresta.Custo

			// Obtém o gScore atual do vizinho, tratando nós ainda não visitados (custo infinito).
			gScoreVizinho, existe := gScore[vizinho]
			if !existe {
				gScoreVizinho = math.Inf(1)
			}

			if gScoreTentativo < gScoreVizinho {
				// Este é o melhor caminho encontrado até agora: atualiza os registros.
				veioDe[vizinho] = atual           // Registra o caminho mais curto.
				gScore[vizinho] = gScoreTentativo // Atualiza o custo real do caminho (gScore).

				novoFScore := gScoreTentativo + h(vizinho, objetivo)
				fScore[vizinho] = novoFScore

				// Verifica se o vizinho já está no Open Set.
				if noVizinho, encontrado := noOpenSet[vizinho]; !encontrado {
					// Se o vizinho for novo, adiciona-o à Fila de Prioridade.
					noVizinho = &NoAEstrela{
						Nome:   vizinho,
						FScore: novoFScore,
					}
					heap.Push(&fp, noVizinho)
					noOpenSet[vizinho] = noVizinho // Adiciona ao mapa de rastreamento.
				} else {
					// Se o vizinho já estiver no Open Set, apenas atualiza seu fScore e reorganiza o heap.
					noVizinho.FScore = novoFScore
					heap.Fix(&fp, noVizinho.Indice)
				}
			}
		}
	}

	// Se o loop terminar sem encontrar o objetivo, o caminho não existe.
	return nil, math.Inf(1)
}
