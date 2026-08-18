# 🚀 Lista de Exercícios Práticos: Fundamentos de Go

### 📌 Instruções Gerais:
1. Resolva cada exercício criando um arquivo `.go` separado (ex: `ex01.go`, `ex02.go`).
2. Tente rodar e testar no seu terminal usando `go run nome_do_arquivo.go`.
3. Preste atenção aos tipos de dados (`int`, `float64`, `string`, `bool`) e às mensagens de erro do compilador.

---

## 🧩 Parte 1: Tipos, Variáveis e Condicionais

### Exercício 1: Classificador de Triângulos
Crie uma função que receba os três lados de um triângulo e retorne:
* `"Equilatero"` se todos os lados forem iguais.
* `"Isosceles"` se dois lados forem iguais.
* `"Escaleno"` se todos os lados forem diferentes.
* `"Invalid"` se os lados não formarem um triângulo (regra: a soma de dois lados quaisquer deve ser estritamente maior que o terceiro lado).

Assinatura: `classificarTriangulo(a, b, c float64) string`

> **Objetivo:** Praticar operadores lógicos (`&&`, `||`), operadores relacionais e desvio condicional (`if/else`).

---

### Exercício 2: Conversor e Formatador de Temperatura
Crie uma função que receba uma temperatura em graus Celsius (`float64`) e retorne dois valores: a temperatura convertida para **Fahrenheit** e para **Kelvin**.
* **Fórmula Fahrenheit:** $F = (C \times 1.8) + 32$
* **Fórmula Kelvin:** $K = C + 273.15$
* Na função `main`, imprima os resultados formatados com exatamente **2 casas decimais** (`fmt.Printf("%.2f | %.2f\n", F, K)`).

> **Objetivo:** Praticar tipos de ponto flutuante (`float64`), operações aritméticas, formatação de saída e retornos múltiplos.

---

## 🔁 Parte 2: Estruturas de Repetição (`for`) e Contadores

### Exercício 3: Estatísticas de um Intervalo
Crie uma função `analisarIntervalo(inicio, fim int) (pares int, impares int, somaTotal int)` que itere de `inicio` até `fim` (inclusive) e retorne:
1. A quantidade de números pares encontrados.
2. A quantidade de números ímpares encontrados.
3. A soma de todos os números desse intervalo.

> **Entrada:** 1 10 | **Saída:** 5 5 55

> **Objetivo:** Praticar laços `for`, o operador de resto da divisão `%` (módulo) e acumuladores/contadores.

---

## 🎯 Parte 3: Ponteiros e Funções

### Exercício 4: Aplicador de Desconto (Mutação via Ponteiro)
Crie uma função `aplicarDesconto(preco *float64, porcentagem float64) error` que:
1. Retorne um erro (`errors.New("Porcentagem Invalida")`) se a porcentagem for negativa ou maior que 100.
2. Caso a porcentagem seja válida, **altere diretamente o valor original** da variável apontada por `preco`, subtraindo a porcentagem informada, e retorne `nil` como erro.

> **Objetivo:** Compreender passagem por valor vs. passagem por referência, o operador `&` (endereço) e `*` (desreferenciação), e o padrão `error` idiomático de Go.

```go
func aplicarDesconto(preco *float64, porcentagem float64) error {

}

func main() {
	var x float64 = a
	var ptr *float64 = &x

	err = aplicarDesconto(ptr, b)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("%.2f", x)
	}
}
```

---

## 📊 Parte 4: Algoritmos Fundamentais (Busca e Ordenação)

### Exercício 5: Busca Linear em Vetor (Slice)
Crie uma função chamada `buscaLinear(slice []int, alvo int) int` que:
* Percorra o slice procurando pelo número `alvo`.
* Se encontrar o valor, retorne o **índice (posição)** da primeira ocorrência.
* Se o elemento não existir no slice, retorne `-1`.

**Exemplo de teste:**
```go
numeros := []int{14, 5, 89, 2, 7, 50}
pos := buscaLinear(numeros, 7) // deve retornar 4
posInexistente := buscaLinear(numeros, 99) // deve retornar -1
```

> **Objetivo**: Praticar iteração com for ou for range, leitura sequencial de dados em memória e saídas antecipadas com return.

---

### Exercício 6: Ordenação por Flutuação (Bubble Sort)

Implemente seu próprio algoritmo de ordenação sem usar funções prontas da linguagem. Crie a função `bubbleSort(numeros []int)` que ordene o slice em ordem crescente.

Como funciona o Bubble Sort:
Compare elementos adjacentes (lado a lado: slice[j] e slice[j+1]).

1. Se o elemento da esquerda for maior que o da direita, faça a troca (swap) de posições entre eles:
2. Repita o processo até que todo o vetor esteja ordenado.

Exemplo de teste:

```go
desordenado := []int{64, 34, 25, 12, 22, 11, 90}
ordenarCrescente(desordenado)
fmt.Println(desordenado) // Saída esperada: [11, 12, 22, 25, 34, 64, 90]
```

> **Objetivo**: Manipulação direta de índices, controle de loops aninhados (for dentro de for) e mutação de slices contíguos na memória.