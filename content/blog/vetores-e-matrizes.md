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

Vamos retomar um exemplo lá do post [Por onde Começar](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/), quando criamos uma variável _a_, colocamos um valor nela e logo depois impresso no console:

```python
a = 1
print(a)
```

Nesse exemplo, criamos uma variável _a_. Conseguimos colocar qualquer coisa nela, embora no exemplo dado tenha sido o número 1, podemos também atribuir um texto:

```python
a = "Bom Dia"
print(a)
```

## Vetores

Agora vamos entender como funcionam os vetores:

> Vetores são conjuntos de variáveis declarados com colchetes [] que serve para guardar um conjunto de dados.

Vamos criar um vetor com todos os dias da semana:

```python
dias = ["domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sabado"]
```

Para acessar cara dia da semana, usamos os colchetes, começando do zero (0). Então para acessar o primeiro elemento do vetor usamos `dias[0]`. Chamamos esse número entre os colchetes de **índice**. Então, vamos imprimir a segunda:

```python
dias = ["domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sabado"]
print(dias[1])
```

Se eu inserir a linha `print(dias[6])` qual dia seria impresso? Dica: _Se o primeiro elemento de índice 0 é o "domingo"  e o segundo de índice 1 é a "segunda", o terceiro de índice 2 seria a "terça"..._

Podemos adicionar também variáveis em cada _índice_ de um vetor:

```python
a = "Bom Dia"
b = "Boa tarde"
mensagens = [a, b]
```

Agora para adicionar um item no vetor também podemos usar o método `append`:

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

Qual elemento seria removido? Dica: _Siga a mesma lógica da pergunta anterior._

Para exercitar, quais os valores que sairão em cada `print` abaixo?

```python
vetor = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
a = vetor[0]+vetor[8]
print(a)
b = vetor[9]-vetor[8]
print(b)
c = vetor[1]*vetor[4]
print(c)
```

Estude, leia e pratique muito os exemplos acima, e use a sua criatividade para pensar em desafios e problemas que envovam vetores. Procure entender o funcionamento e quando se sentir confortável, avançe com a leitura.


## Matrizes

Depois de compreendido o conceito e aplicação de Vetor, podemos continuar com _Matrizes_:

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

As matrizes não tem limites de dimensões, isto é (talvez seja bem bugado ler isso pela(s) primeira(s) vez(ez) xD), podemos colocar vetores dentro de vetores:

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
```

**Em outras linguagens a sintaxe da declaração de uma matriz muda. Nesse caso, podemos ter qualquer coisa dentro do índice de um vetor, um objeto, número, texto etc. Em C, por exemplo, é uma [lingagem fortemente tipada](http://www.ppgsc.ufrn.br/~rogerio/material_auxiliar/CLP20131_tipos_semantica.pdf) e em sua declaração precisamos definir o tipo da variável, se é inteiro (int), número real (float), de caractere (char) etc. Mas não se preocupe com isso, o conceito é o mesmo!**

As possibilidades são infinitas, porém há alguns casos específicos para que a aplicação dessa técnica seja plausível. Brinque um pouco, entenda a lógica e se divirta. Abraços e até a próxima!