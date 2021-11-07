+++
date = "2020-09-10T16:24:06-03:00"
title = "Vetores e Matrizes"
categories = ["algoritmo"]
tags = ["matrizes","vetores"]
banner = "img/banners/cube-matrix.jpg"
+++

Recomendamos a leitura do artigo sobre [Por onde Começar](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/), antes da leitura deste artigo.

## Variáveis, conceitos básicos

No exemplo abaixo criamos uma variável chamada _a_ e associamos o valor numérico 1 à variável.

```python
a = 1
print(a)
```

No exemplo abaixo criamos uma variável _b_ e associamos uma frase "Bom Dia" à variável.

```python
b = "Bom Dia"
print(b)
```

## Vetores

> Vetores são estruturas e facilitam a armazenagem e manipulação de um conjunto de dados.
No exemplo abaixo criamos um vetor chamado dias e atribuímos os dias da semana como o conteúdo do vetor.

```python
dias = ["domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sabado"]
```

Para acessar cada dia da semana, usamos a representação dias[i]  onde i representa o **índice** do elemento do vetor.
No caso da linguagem Python o primeiro elemento de um vetor tem i=0, logo dias[0] representa "domingo" e dias[1] "segunda" e assim sucessivamente.

O exemplo abaixo demonstra como um programa em Python imprimiria o valor "segunda":

```python
dias = ["domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sabado"]
print(dias[1])
```
Podemos adicionar também variáveis em cada _índice_ de um vetor como demonstrado no exempo abaixo:

```python
a = "Bom Dia"
b = "Boa tarde"
mensagens = [a, b]
```

Para adicionar um item no vetor também podemos usar o método `append` conforme o exemplo abaixo:

```python
mensagens = ["Bom Dia", "Boa Tarde"]
mensagens.append("Boa Noite")
print(mensagens)
```

O novo vetor será `["Bom Dia", "Boa Tarde", "Boa Noite"]`!

Para remover um elemento do vetor usamos o método `pop`, passando o _índice_ como parâmetro:

```python
mensagens = ["Bom Dia", "Boa Tarde", "Boa noite", "Como vai?"]
mensagens.pop(3)
print(mensagens)
```

Qual elemento seria removido? Dica: _Se o primeiro elemento de índice 0 é o "domingo"  e o segundo de índice 1 é a "segunda", o terceiro de índice 2 seria a "terça"..._

Para exercitar, quais os valores que serão exibidos em cada `print` abaixo?

```python
vetor = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
a = vetor[0]+vetor[8]
print(a)
b = vetor[9]-vetor[8]
print(b)
c = vetor[1]*vetor[4]
print(c)
```

Estude, leia e pratique muito os exemplos acima. Use a sua criatividade para pensar em desafios e problemas que envolvam vetores. Procure entender o funcionamento e quando se sentir confortável, avance com a leitura.


## Matrizes

Depois de compreendido o conceito e aplicação de Vetores, podemos continuar com _Matrizes_:

> Matrizes são conjuntos de vetores!
Para entender melhor, vamos fazer 3 vetores, e **inserir cada vetor em um _índice_ de um novo vetor**:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marceo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = [vetor1, vetor2, vetor3]
```

Cada vetor se comportará como uma coluna, e cada índice da matriz representará uma coluna. Ou seja: A primeira coluna da matriz (_índice_ 0) será o `vetor1`, a segunda (_índice_ 1) o `vetor2` e terceira (_índice_ 2) `vetor3`.

Podemos também usar o método `append` para inserir cada vetor:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marceo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = []
matriz.append(vetor1)
matriz.append(vetor2)
matriz.append(vetor3)
```

Para acessar cada elemento na matriz, **precisamos de dois colchetes!** Pois o primeiro será o _índice da coluna_ e o segundo, o _índice da linha_. Vamos chamar o elemento _"Bom Dia"_:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marcelo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = [vetor1, vetor2, vetor3]
print(matriz[0][0]) # Matriz[0] retorna todo o nosso vetor (a coluna vetor1, de índice 0), o segundo colchete representa o índice do vetor resultante
```

Se a linha `print(matriz[0][0])` fosse substituída por `print(matriz[0][1])` qual seria o elemento impresso? _Dica: `matriz[0]=vetor1` e `vetor1[0] = "Bom Dia"`._

Agora, vamos inserir uma saudação no `vetor1`. Para isso, usaremos também o método `append`. Como o vetor1 está no primeiro índice da matriz (0), escrevemos:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marcelo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = [vetor1, vetor2, vetor3]
matriz[0].append("Bom Final de Semana")
print(matriz[0]) # O que será impresso?
```

Podemos inserir o seu nome também:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marcelo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = [vetor1, vetor2, vetor3]
matriz[1].append("Escreva seu nome aqui!")
print(matriz[1])
```

Para remover um elemento ou até mesmo uma coluna inteira, usamos o método `pop`, da mesma forma que aprendemos em _Vetores_:

```python
vetor1 = ["Bom Dia", "Boa Tarde", "Boa Noite"]
vetor2 = ["Bruno", "Marcelo", "Bernardo"]
vetor3 = ["Como Vai?", "Que horas são?" "Vamos programar?"]
matriz = [vetor1, vetor2, vetor3]
#Retirando o nome Bruno
matriz[1].pop(0)
#Retirando toda a coluna 3 (de índice 2)
matriz.pop(2)
```

Agora você já sabe como inserir, ler e excluir tanto uma célula quanto uma coluna de uma matriz. Pratique bastante pois na sua jornada como programador, o uso dessas ferramentas é essencial.

## Ah, mais uma coisa!

As matrizes não tem limites de dimensões, podemos colocar vetores dentro de vetores:

```python
pessoas = []
nomes = ["Hinata", "Naruto", "Sasuke"]
sobrenomes = ["Hyuuga", "Uzumaki", "Uchiha"]
pessoas.append(nomes)
pessoas.append(sobrenomes)
lugares = []
estados = ["São Paulo", "Rio de Janeiro", "Sergipe", "Fortaleza"]
capitais = ["São Paulo", "Rio de Janeiro", "Aracaju", "Ceará"]
lugares.append(estados)
lugares.append(capitais)
conjunto = []
conjunto.append(pessoas)
conjunto.append(lugares)
print(conjunto[0][1][0])
```

Autor do Artigo: [Bruno Sana](https://github.com/brunosana)