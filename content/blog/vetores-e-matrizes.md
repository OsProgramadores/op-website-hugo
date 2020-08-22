+++
title = "Introdução a Vetores e Matrizes"
date = "2020-07-30T23:49:00-03:00"
tags = ["programação","algoritmo","vetor","matriz"]
categories = ["algoritmo"]
banner = "img/banners/matrix-cube.jpg"
+++

Olá, eu sou o [Bruno Sana](https://github.com/brunosana) e vou introduzir os conceitos e aplicações de **Vetores e Matrizes**. Para este post, nos códigos a linguagem [Python](https://www.python.org/)!

---

# Conceito

Absolutamente tudo que fazemos na vida tem quase sempre o mesmo propósito: *Resolver Problemas*. Não é diferente com a programação, conforme vamos escrevendo nossas linhas de código, precisamos garantir que está escrita de forma otimizada e eficaz, não bastando apenas o resultado ser o esperado. Os Vetores e Matrizes são vitais na maioria dos sistemas e são usados com inteligência e precisão.

Normalmente, a primeira coisa que vem à cabeça quando pensamos em Vetores ou Matrizes nos remete aos campos da matemática. Não é pra menos, existe uma relação direta no conceito e aplicação tanto de vetor, como de matriz.

## Vetores

Vejamos a seguinte situação:

> A escola *Os Programadores* precisa de uma solução para calcular a média de uma turma de 5 alunos.

Você, como todo programador, imediatamente pensa na solução, que logica e convenientemente, seria:

```python
nota_aluno1 = float(input('Digite a nota do aluno 1:'))
nota_aluno2 = float(input('Digite a nota do aluno 2:'))
nota_aluno3 = float(input('Digite a nota do aluno 3:'))
nota_aluno4 = float(input('Digite a nota do aluno 4:'))
nota_aluno5 = float(input('Digite a nota do aluno 5:'))

media = (nota_aluno1 + nota_aluno2 + nota_aluno3 + nota_aluno4 + nota_aluno5)/10

print('Media dos alunos: {}'.format(media))
```

Perfeito. Resolvemos o problema!

Porém, no outro dia, recebemos um alteração:

> A escola *Os Programadores*, satisfeita com o algoritmo anterior, resolveu expandir a turma para 100 alunos. Ela precisa de uma nova solução!

Garanto que você não está afim de criar mais 90 variáveis para resolver o problema, como mencionei antes, não basta apenas ter o resultado correto. Situações como essa dão sentido aos **vetores**. Podemos então defini-lo:

> Um vetor é um conjunto sequencial de variáveis acessadas pelo seu respectivo índice.

Quando declaramos uma variável, um espaço na memória é reservado para ela, para podermos atribuir valores a ela, e por fim, armazenar. O propósito do vetor é reservar uma variável que tenha acesso à várias posições de memória.

Significa que, ao invés de ter várias variáveis cada uma para cada aluno, temos um grupo de variáveis que chamaremos de nota, e cada índice representará um aluno.

Para definir um vetor em `python` usamos o colchete:

```python
notas_alunos = []
```

Cada elemento possui um espaço na memória:

![Vetor](/img/conteudos-de-artigos/vetor-2.png)


Usando o vetor, podemos acessar cada valor armazenado por ele pelo seu respectivo índice. Para entender melhor, vamos analisar a imagem abaixo:

![Vetor](/img/conteudos-de-artigos/vetor-1.png)

Cada índice representa um espaço na memória e cada espaço pode ser acessado pelo seu respectivo índice. O primeiro índice sempre é o zero (0). Para acessar cada elemento de um vetor em python usamos a seguinte sintaxe:

```python
notas_alunos[2]
```

**Com o código acima, estou buscando o terceiro valor no vetor.**

Para inserir um elemento em um vetor, usamos a função `append`:

```python
notas_alunos = []
#Adicionando um elemento ao final do vetor, como ele está vazio, este ocupará a primeira posição, de índice 0
notas_alunos.append(8)
```

