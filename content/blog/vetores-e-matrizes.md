+++
title = "Introdução a Vetores e Matrizes"
date = "2020-07-30T23:49:00-03:00"
tags = ["programação","algoritmo","vetor","matriz"]
categories = ["algoritmo"]
banner = "img/banners/matrix-cube.jpg"
+++

Olá, eu sou o [Bruno Sana](https://github.com/brunosana) e vou introduzir os conceitos e aplicações de **Vetores e Matrizes**. Para este post usaremos a linguagem [Python](https://www.python.org/)!

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
notas = [10, 8, 7, 3]
notas[2] #Irá retornar o valor 7

notas[3] # Qual será o valor que iria retornar? rs
```

**Com o código acima, estou buscando o terceiro valor no vetor.**

Para inserir um elemento em um vetor, usamos a função `append`:

```python
notas = []
#Adicionando um elemento ao final do vetor, como ele está vazio, este ocupará a primeira posição, de índice 0
notas.append(7.5)
#Após adicionar, o vetor final será [7.5]
```

**Todo primeiro elemento do vetor tem índice 0.**

Para alterar um elemento de um array, basta sobrescrever o seu valor, assim como uma variável:

```python
notas = [5.5, 7, 8.2, 6, 10]
notas[3] #Irá retornar o valor 6
notas[3] = 8 # Agora o novo valor será 8
```

Para remover um elemento de um array, usamos a função `pop()`, passando o índice como parâmetro:

```python
notas = [5.5, 7, 8.2, 6, 10]
notas.pop(0) #Remove o primeiro elemento (5.5)
# A função pop sem argumentos remove o último índice
notas.pop() # Qual o elemento que será removido ? rs
```

Agora você já sabe trabalhar com vetores!! Então, como faríamos para resolver o nosso problema da escola Os Programadores? Agora é simples!! Com a ajudinha do `for` nós podemos repetir o mesmo código quantas vezes quisermos:

```python
#Escola Os Programadores

notas = []
soma = 0
numero_de_alunos = 100

# Criamos um laço que começará do zero e irá se repetir 100 vezes (Que é o valor da variável).
for i in range(numero_de_alunos):
    #Aqui eu somo o i+1 pois embora o computador comece a contagem do 0, o primeiro valor não nulo para nós é o 1, a única diferença é na hora de imprimir, pois com o vetor com índices de 0..99, teremos os alunos 1..100.
    nota_aluno = float(input('Digite a nota do aluno {}: '.format(i+1)))
    notas.append(nota_aluno)
    # Otimizaremos o código já somando as notas!
    soma = soma + nota_aluno

# Agora é só fazer o cálculo da média e imprimir o seu valor!
media = soma/numero_de_alunos
print('Media = {}'.format(media))
```

Perfeito!! Atingimos o nosso objetivo e mais, podemos alterar o número de alunos a qualquer momento, apenas mudando o valor da variável `numero_de_alunos`.

**Agora é a sua vez! treine e refaça todos os exemplos, adicionando outros problemas para fixar o assunto! Quem sabe validar todas as notas digitadas (>=0 e <=10), separar as notas maiores que 5 das menores, calcular a média delas separada. E se o número de alunos precisasse ser lido também? tudo depende de você!**

## Matrizes

Recebemos outra consideração:

> A escola *Os Programadores*, que resolveu o seu problema do número de alunos, gostou tanto do projeto que resolveu implementar em toda a sua escola, com todas as 10 turmas, bem como calcular a sua media geral!

Como o saudoso confrade e consagrado Fausto Silva diria: **EITA!!** Situações como essa demandam o uso de **Matrizes!!**

Naturalmente, com o conhecimento que temos, precisaríamos de um vetor para cada turma, correto? Mas esse tipo de prática não é recomendada, pois cada vez que o número de turmas mudar o código precisaria ser mudado tamém.

Façamos então a seguinte definição:

> Matriz é um conjunto de vetores que são acessados pelo seu respectivo índice.

Vamos entender na prática, criando uma matriz em python e depois analisando-a:

```python
matriz = []

turma1 = [6, 6.1, 5, 2.2, 0]
turma2 = [3, 7, 7, 8, 3]
turma3 = [5, 7, 9, 8, 10]

matriz.append(turma1)
matriz.append(turma2)
matriz.append(turma3)

```

Para facilitar, vamos analisar a imagem:

![Vetor](/img/conteudos-de-artigos/vetor-3.png)

Cada coluna é acessada normalmente, como um vetor, e a linha com um segundo colchete. Podemos definir uma matriz de forma direta também:

```python
matriz = [[4, 6.3, 2, 3.5], [5, 6, 7, 8], [10, 9, 10, 6]]
#Acessando a coluna de uma matriz:
matriz[0] #Retornará o vetor [4, 6.3, 2, 3.5]
#Acessando um elemento específido de uma linha
matriz[2][1] # Coluna 2 - Linha 1. matriz[2] = [10, 9, 10, 9], portanto, o segundo elemento (de índice 1) será o 9.
```

As operações com matrizes são bem parecidas com as de vetores, portanto:

```python
matriz = []
#Adicionando uma coluna em uma matriz:
coluna = [1, 2, 3, 4, 5, 6]
matriz.append(coluna)
#Adicionando um elemendo na coluna de uma matriz:
matriz[0].append(8) #Coluna resultante: [1, 2, 3, 4, 5, 6, 8]

#Removendo um elemento de uma coluna:
matriz[0].pop(2) # Deleta o elemento de índice 2 (3)
#Removendo uma coluna de uma matriz:
matriz.pop(0) #Função pop() sem argumentos remove a última coluna
```

Certo, já sabemos todas as operações envolvendo vetores e matrizes, então podemos desenvolver nosso algoritmo:

```python
num_turmas = 10
num_alunos_por_turmas = 100
escola = []
media_geral = 0
for i in range(num_turmas):
    notas = []
    soma = 0
    media_turma = 0
    for j in range(num_alunos_por_turmas):
        nota_aluno = float(input('Nota - Turma {} Aluno {}: '.format(i+1, j+1)))
        notas.append(nota_aluno)
        soma = soma + nota_aluno
    media_turma = soma/num_alunos_por_turmas
    media_geral = media_geral + media_turma
    print('Media da turma {} : {}'.format(i+1, media_turma))
print('Media geral: {}'.format(media_geral))
```

Claro que de um primeiro momento parece um código confuso, mas vamos analisar de acordo com o que já sabemos ok?

1. Para que possamos percorrer um vetor, precisamos de um laço, nesse caso usamos um `for`! Ele percorre índice por índice.
2. Com a variável que alocamos junto com o `for` (i) podemos acessar o valor do respectivo índice (`vetor[indice]`).
3. Nessa situação, temos uma matriz, um conjunto de vetores, para cada índice que percorremos, ao invés de termos um valor, temos um outro vetor. Portanto, precisaríamos de um outro for! Enquanto um deles varre a linha, outro varre a coluna, daí a necessidade de dois `for`.
4. Para otimizar o código, toda vez que o algoritmo lê uma nota, ele soma com a média da turma, e ao final do for, imprimo a média, e somo essa média com a variável media_geral. É um pouco confuso se olhar de primeira, mas fica claro quando olhamos linha por linha.

## Conclusão

Tanto vetores como matrizes precisam ser usados com muita inteligência pois podem facilmente criar problemas de performance no seu algoritmo, se mal utilizado. Na mão de que mentende é uma ferramenta potente e eficaz.

Como prática deixarei alguns desafios para que você possam acrescentar dificuldade cada vez mais:

1. Faça com que o algoritmo leia um número genérico de turmas e alunos por turmas
2. Valide todas as notas (>=0 e <=10)
3. Imprima qual turma teve a maior média e qual teve a menor
4. Imprima qual o aluno com a maior nota e qual aluno com a menor

Com esses desafios, a ideia é que você pratique e compreenda o que abrange vetores e matrizes e consiga aplicar em seus projetos de forma otimizada.

## Ah, mais uma coisa!

As matrizes não tem limites de dimensões, isto é, até o momento fizemos um vetor de duas dimensões, como linha e coluna. Mas na verdade, dentro de cada elemento da coluna, podemos ter outro vetor, se tornando uma matriz de 3 dimensões `[x][y][z]`. E dentro desse vetor poderia ter outro, se tornando um vetor de 4 dimensões `[x][y][z][w]`.

**Em outras linguagens a sintaxe da declaração de uma matriz muda. Nesse caso, podemos ter qualquer coisa dentro do índice de um vetor, um objeto, número, string etc. Em C, por exemplo, é uma lingagem fortemente tipada e em sua declaração precisamos definir o tipo da variável, se é inteiro (int), número real (float), de caractere (char) etc. Mas o conceito é o mesmo!**


As possibilidades são infinitas, porém há alguns casos específicos para que a aplicação dessa técnica seja plausível. Brinque um pouco, entenda a lógica e se divirta. Abraços e até a próxima!