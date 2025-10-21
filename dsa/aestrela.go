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

func AEstrela(grafo Grafo, inicio, objetivo No, obstaculos map[No]bool, h FuncaoHeuristica) ([]No, float64) {
	if obstaculos[inicio] || obstaculos[objetivo] {
		return nil, math.Inf(1)
	}

	fp := make(FilaDePrioridade, 0)
	heap.Init(&fp)
	noOpenSet := make(map[No]*NoAEstrela)
	veioDe := make(map[No]No)
	gScore := map[No]float64{inicio: 0.0}
	fScore := map[No]float64{inicio: h(inicio, objetivo)}

	noInicio := &NoAEstrela{Nome: inicio, FScore: fScore[inicio]}
	heap.Push(&fp, noInicio)
	noOpenSet[inicio] = noInicio

	for fp.Len() > 0 {
		noAtual := heap.Pop(&fp).(*NoAEstrela)
		atual := noAtual.Nome
		delete(noOpenSet, atual)

		if atual == objetivo {
			return reconstruirCaminho(veioDe, atual), gScore[atual]
		}

		for _, aresta := range grafo[atual] {
			vizinho := aresta.Para

			if obstaculos[vizinho] {
				continue
			}

			gScoreTentativo := gScore[atual] + aresta.Custo
			gScoreVizinho, existe := gScore[vizinho]
			if !existe {
				gScoreVizinho = math.Inf(1)
			}

			if gScoreTentativo < gScoreVizinho {
				veioDe[vizinho] = atual
				gScore[vizinho] = gScoreTentativo
				novoFScore := gScoreTentativo + h(vizinho, objetivo)
				fScore[vizinho] = novoFScore

				if noVizinho, encontrado := noOpenSet[vizinho]; !encontrado {
					noVizinho = &NoAEstrela{Nome: vizinho, FScore: novoFScore}
					heap.Push(&fp, noVizinho)
					noOpenSet[vizinho] = noVizinho
				} else {
					noVizinho.FScore = novoFScore
					heap.Fix(&fp, noVizinho.Indice)
				}
			}
		}
	}
	return nil, math.Inf(1)
}
