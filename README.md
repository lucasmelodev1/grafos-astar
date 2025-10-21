# Projeto de Implementação do Algoritmo A*

Este projeto implementa o **algoritmo A\*** para encontrar o caminho mais curto em grafos. Ele lê grafos de arquivos de entrada, calcula o caminho mais curto entre um nó inicial e um nó objetivo, e gera relatórios em **texto** e **Graphviz (DOT)**.

---

## Estrutura do Projeto

```
├── exemplos/ # Pasta com os exemplos de grafos
│ ├── exemplo1/
│ │ ├── entrada.txt
│ │ ├── saida.txt
│ │ └── saida.dot
│ ├── exemplo2/
│ │ └── ...
├── grafos/
│ ├── dsa/ # Estruturas de dados e implementação do A*
│ │ ├── a_estrela.go
│ │ └── types.go
│ └── iolib/ # Pacote de IO para leitura e escrita de arquivos
│ ├── ler_escrever.go
├── main.go # Arquivo principal que processa os exemplos
└── README.md
```

---

## Como os arquivos de exemplo devem ser estruturados

Cada exemplo está em uma pasta própria dentro de `exemplos/`. Dentro da pasta do exemplo:

- `entrada.txt` → contém a definição do grafo, nó inicial e objetivo
- `saida.txt` → será gerado pelo programa com o caminho encontrado e custo total
- `saida.dot` → será gerado pelo programa com o grafo em formato Graphviz, destacando o caminho encontrado

### Formato de `entrada.txt`

inicio=A
objetivo=F

A B 2
A C 4
B C 1
B D 7
C E 3
D F 1
E F 5


- Linhas com três elementos definem **arestas**: `origem destino custo`
- `inicio` define o nó inicial
- `objetivo` define o nó de destino

---

## Como executar

1. Certifique-se de ter o **Go** instalado (versão 1.21 ou superior recomendada).
2. Compile ou execute diretamente:

```bash
go run main.go
```

3. O programa vai percorrer todas as pastas em exemplos/, processar cada grafo e gerar os arquivos de saída (saida.txt e saida.dot) na mesma pasta do exemplo.

## Pacotes

### 1. `grafos/dsa`

Contém as **estruturas de dados** e a **implementação do A\***.

- `No` → representa um nó do grafo (string)
- `Aresta` → representa uma aresta do grafo com destino e custo
- `Grafo` → mapa de nós para listas de arestas
- `AEstrela` → função que implementa o algoritmo A*

### 2. `grafos/iolib`

Pacote responsável por **ler e escrever arquivos**:

- `LerGrafoDoArquivo(nome string)` → lê o grafo, início e objetivo
- `SalvarSaida(caminho, conteudo string)` → salva arquivos TXT ou DOT
- `GerarRelatorio(nome string, caminho []No, custo float64)` → gera saída em texto
- `GerarGraphviz(grafo Grafo, caminho []No, inicio, objetivo No)` → gera arquivo DOT, destacando o caminho encontrado

---

## Heurística

Atualmente o projeto usa uma **heurística simples** (distância zero), equivalente ao algoritmo de Dijkstra.  
É possível substituir `h(n)` por qualquer função heurística admissível para direcionar melhor a busca.

---

## Saída esperada

### Arquivo TXT (`saida.txt`)

════════════════════════════════════════════
Exemplo: exemplo1
────────────────────────────────────────────
Caminho encontrado:
[A B C E F]
Custo total: 11.00
════════════════════════════════════════════

### Arquivo DOT (`saida.dot`)

digraph G {
  rankdir=LR;
  node [shape=circle];
  "A" [style=filled, fillcolor=green];
  "F" [style=filled, fillcolor=red];
  "A" -> "B" [label="2.00", style=bold, color=green];
  "A" -> "C" [label="4.00", style=solid, color=black];
  "B" -> "C" [label="1.00", style=bold, color=green];
  "B" -> "D" [label="7.00", style=solid, color=black];
  "C" -> "E" [label="3.00", style=bold, color=green];
  "D" -> "F" [label="1.00", style=solid, color=black];
  "E" -> "F" [label="5.00", style=bold, color=green];
}

