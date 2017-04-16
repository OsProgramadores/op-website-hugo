+++
title = "Estruturas de dados em Python"
categories = ["linguagens"]
tags = ["python", "estruturas de dados", "matrizes"]
date = "2017-04-14T10:22:25-03:00"
banner = "img/banners/math-matrix-800px.png"

+++

## Matrizes em _Python_[^fa]

As matrizes[^f1] são uma das formas de estruturas de dados utilizadas para resolver uma série de problemas. Na linguagem de programação _Python_, as matrizes podem ser representadas como um conjunto de vetores. Para uma definição formal de matriz, sugere-se consultar as fontas ao fim desse artigo[^ff].

A estrutura de dados _list_ representa um vetor, que no nosso exemplo abaixo cria um vetor sem elementos:

```sh
Python 3.4.3 (default, Nov 17 2016, 01:08:31)
[GCC 4.8.4] on linux 
Type "help", "copyright", "credits" or "license" for more information.
>>> notas = list()
>>>
```

Assim, inicializamos um vetor. Isso também poderia ser feito já com alguns valores, dependendo do problema a ser resolvido. Por exemplo, uma _list_ com as notas de matemática de uma turma de 5 (cinco) estudantes:
 

```sh
>>> notas_matematica = [8.5, 9.5, 10, 6.5, 8.0]
>>> notas_matematica
[8.5, 9.5, 10, 6.5, 8.0]
>>>
```

Podemos observar que ao chamar a _list_ notas de uma disciplina, o interpretador retorna os valores armazenados. A principal característica desse vetor é que ele é **unidimensional**, isto é, possui apenas **uma** dimensão. As matrizes possuem pelo menos 2 (duas) dimensões, isto é, elas são um conjunto de vetores, que na sua forma plana, representam as _linhas_ e _colunas_.

E se fosse necessário representar mais de uma disciplina? Para representar mais uma disciplina, por exemplo, _Lingua Portuguesa_, temos que incluir mais uma linha.

```sh
>>> notas_matematica = [8.5, 9.5, 10, 6.5, 8.0]
>>> notas_lingu_port = [9.5, 3.4, 10, 5.5, 7.5]
```

Note que agora temos dois vetores, porém eles ainda não constituem uma matriz. Por enquanto, eles são vetores distintos. Então precisamos juntar esses dois vetores de maneira a formar nossa matriz _notas_. 
 
### Métodos de uma _list_

Para juntar esses vetores, vamos utilizar o método _append_. Considere no início desse artigo que inicializamos a matriz _notas_ como uma _list_ vazia. Para consultar os métodos disponíveis podemos utilizar a função _dir(notas)_. Esse comando retorna uma _list_, representada por informações entre os sinais _[ ]_. 

```sh
>>> dir(notas)
>>> ['__add__', '__class__', '__contains__', '__delattr__', '__delitem__', '__delslice__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getslice__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__setslice__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']
```
Então, ao utilizar esse método, vamos juntar esses dois vetores.
```sh
>>> notas.append(notas_matematica)
>>> notas.append(notas_lingu_port)
>>> notas
>>> [[8.5, 9.5, 10, 6.5, 8.0], [9.5, 3.4, 10, 5.5, 7.5]]
```

Assim temos uma _list_ contendo duas _list_. Isso pode ser verificado com a função _len()_. Para verificar a quantidade de elementos em cada vetor dessa matriz _notas_, o que deve ser feito?

### Funções _len()_ e _type()_

```sh
>>> len(notas)
2
```

Mais um conceito deve ser aprendido. Nessa matriz _notas_, a primeira _list_, que representam as notas de matemática,  devem ser acessadas indicando sua posição na matriz _notas_, isto é, em _Python_, indicando o valor _zero_, na seguinte forma:

```sh
>>>
>>> notas[0]
[8.5, 9.5, 10, 6.5, 8.0] 
>>>
>>>
>>> notas[1]
[9.5, 3.4, 10, 5.5, 7.5] 
>>> 
```
Respondendo à nossa pergunta anterior, podemos verificar que cada um desses vetores é uma _list_. Confirma-se usando a função _type()_. Isto significa que sendo do mesmo tipo, possui os mesmos métodos verificados.

```sh
>>> type(notas[0])
<class 'list'> 
```

### Acesso aos elementos da matriz _notas_

Assim, depois dessa introdução finalmente vamos trabalhar efetivamente com a matriz. Vamos efetuar as seguintes operações:

- Calcular a média de notas de matemática;
- Calcular a média de notas de língua portuguesa; e
- Calcular a média de notas de todas as disciplinas.



#### Média das notas de matemática ou língua portuguesa

Consideremos para o aprendizado que vamos utilizar variáveis auxiliares. À medida que o tempo de treino se estenda, elas poderão ser deixadas de lado.

Para calcular a média das notas de matemática dentro da matriz _notas_ utilizamos o primeiro vetor _notas[0]_. A mesma operação poderá ser utizada para a disciplina de língua portuguesa, através do vetor _notas[1]_.

```sh
>>> soma_matematica=0
>>> for nota in enumerate(notas[0]):    <-- nota é diferente de notas[0],  singular e plural (conjunto).
>>>     print(nota)                     <-- código de construção.
>>>     soma_matematica=soma_matematica+nota[1]
```
O resultado dessas primeiras linhas para o código de construção _print(nota)_ e a variável **soma_matematica** retornam as seguintes informações:

```sh
(0, 8.5)
(1, 9.5)
(2, 10)
(3, 6.5)
(4, 8.0)
>>> soma_matematica
42.5
>>>
```
A média pode ser verificada com a manipulação dos resultados prévios:

```sh
>>> media_matematica=soma_matematica/len(notas[0])
>>> print(media_matematica)
8.5 
```

#### Média de todas as notas

O acesso aos elementos de ambos os vetores pode ser realizado com um laço duplo:

```sh
>>> soma_notas=0
>>> contagem=0
>>> for v in notas:                   <-- Percorre cada vetor na matriz notas.
>>>     for nota in v:	              <-- Percorre cada nota em cada vetor (v).
>>>         soma_notas=soma_notas+nota
>>>         contagem=contagem+1
>>>
```
A soma pode ser recuperada com o _print(soma_notas)_ e a média com o auxílio de outras variáveis, dessa forma:

```sh
>>> print(soma_notas)
78.4
>>> print(contagem)
10
>>> print(soma_notas/contagem)
7.84
```
Logo, conseguimos varrer todos os elementos da matriz _notas_. Essa varredura utiliza dois laços _for_, mas não obrigatoriamente devem ser utilizados. Podem ser utilizados os laços _do while_ ou _while do_, principalmente quando a dimensão dos vetores não é conhecida.

## Considerações finais

A matriz como estrutura de dados pode resolver diversos problemas. Muitas aplicações que vão desde a solução de sistemas lineares e gráficos se utlizam dessa estrutura para representar seus dados e transformações necessárias em suas representações. As propriedades das matrizes podem ser melhor estudadas na referência fornecida. O mais importante desse assunto é saber como acessar os elementos na matriz plana através dos seus índices.

[^fa]: David Gesrob: [E-mail david81br@gmail.com](mailto: david81br@gmail.com);
[^f1]: Imagem do Banner: https://openclipart.org/detail/108619/math-matrix
[^ff]: Referências de Matrizes: https://pt.wikipedia.org/wiki/Matriz_(matem%C3%A1tica)#Matriz_sim.C3.A9trica 
