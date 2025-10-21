package dsa

type FilaDePrioridade []*NoAEstrela

func (fp FilaDePrioridade) Len() int { return len(fp) }

func (fp FilaDePrioridade) Less(i, j int) bool {
	return fp[i].FScore < fp[j].FScore
}

func (fp FilaDePrioridade) Swap(i, j int) {
	fp[i], fp[j] = fp[j], fp[i]
	fp[i].Indice = i
	fp[j].Indice = j
}

func (fp *FilaDePrioridade) Push(x any) {
	n := len(*fp)
	item := x.(*NoAEstrela)
	item.Indice = n
	*fp = append(*fp, item)
}

func (fp *FilaDePrioridade) Pop() any {
	antigo := *fp
	n := len(antigo)
	item := antigo[n-1]
	antigo[n-1] = nil
	item.Indice = -1
	*fp = antigo[0 : n-1]
	return item
}
