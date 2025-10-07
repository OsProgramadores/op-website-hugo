+++
title = "Raízes quadradas e aritmética modular"
date = "2024-09-06T22:00:00-03:00"
tags = ["Algoritmos", "artigos"]
categories = ["artigos", "algoritmo"]
banner = "img/banners/sqrt.png"
+++

Olá, meu nome é [Elias Rodrigues Emídio](https://github.com/eremidio) sou estudante de física teórica e computacional.

<!--more-->

## Índice

-   [1. Introdução](#1-introdução)
-   [2. Aritmética modular - fundamentos matemáticos](#2-aritmética-modular-fundamentos-matemáticos)
-   [3. Aritmética modular e linguagens de programação](#3-aritmética-modular-e-linguagens-de-programação)
-   [4. De volta ao problema original: uma solução aprimorada](#4-de-volta-ao-problema-original-:-uma-solução-aprimorada)
-   [5. Operadores de manipulação de bits e representação binária](#5-operadores-de-manipulação-de-bits-e-representação-binária)
-   [6.  De volta ao problema original: uma solução aprimorada parte 2](#6-de-volta-ao-problema-original-:-uma-solução-aprimorada-parte-2)
-   [7. Conclusão](#7-conclusão)
-   [8. Bibliografia](#8-bibliografia)
  
# 1. Introdução {#1-introdução}


Vamos considerar o seguinte problema: dado um certo intervalo de valores inteiros, por exemplo, consideremos os números de 1 a 100; o objetivo é determinar entre eles quais números são quadrados perfeitos, isto é, dado dois inteiros 'm' e 'n' quais satisfazem relações do tipo m=n². Para valores pequenos tal tarefa é simples e pode ser computada mentalmente. Para os números inteiros de 1 a 100 temos: 1=1², 4=2², 9=3², 16=4², 25=5², 36=6², 49=7², 64=8², 81=9², 100=10².

Para intervalos maiores, digamos de 1 até 10000, o cálculo mental é complicado, porém do ponto de vista computacional, podemos usar uma abordagem simples para atacar o problema.

Muitas linguagens de programação disponibilizam uma biblioteca de funções matemáticas básicas que são usadas com certa frequência, entre elas uma função geralmente nomeada como sqrt (do inglês "square root"), que retorna a raiz quadrada de um número real ou inteiro, normalmente ela retorna um valor real (floating pointing number) que contempla raízes quadradas não exatas (que não resultam em valores inteiros ou com um número finito de dígitos decimais). Não estamos interessados aqui nos algoritmos usados para computar raízes quadradas. Por hora nos basta o fato de que podemos incluir esta função em nosso programa importando a biblioteca correspondente. 

Em Python, por exemplo, podemos usar a seguinte linha de comando no script:

```python
  from math import sqrt
```

Em C (e C++) devemos incluir o arquivo <math.h> (ou em C++ <cmath>) e usar uma flag -lm (para indicar para o compilador a localização das definições das funções matemáticas a serem usadas) no momento de compilar o programa (processo que gera um arquivo binário com as instruções a serem executadas pelo computador passo a passo). Para isso usamos um dos comandos (uma diretiva de preprocessamento sendo tecnicamente preciso):

```cpp
  #include<math.h>
  #include<cmath>
```


Uma vez feito isso estamos apto a usar a função sqrt() em nosso programa, passando para ela um argumento que consiste em um número, por exemplo em Python, podemos usar a sintaxe:

```python

  y=sqrt(89)
```

Neste caso a variável y aloca um valor aproximado da raiz quadrada de 89. Este é um valor não exato, de forma que y é apenas uma aproximação do valor correto.

Uma abordagem simples para resolver o problema proposto no início deste texto seria a seguinte: para todo valor a ser testado, usamos a função sqrt() para calcular um valor aproximado da raiz quadrada r de um número, arredondando o valor de r (novamente usando uma função matemática como as funções maior/menor inteiro - floor()/ceil(), geralmente disponíveis nas mesmas bibliotecas matemáticas acima citadas) podemos testar para todo inteiro n se a relação r²=n é satisfeita.

Fica como um exercício ao leitor, criar um programa para implementar tal algoritmo.

O leitor que fez a proposta de exercício acima verificará que para valores acima de 1000 o programa em questão se torna extremamente lento. Portanto abordagens mais eficazes se fazem necessárias.

Tal tarefa é extremamente importante em muitos algoritmos de teoria de números computacional e criptografia. Nosso objetivo é atacar este problema de uma forma extremamente eficaz de modo a processar valores da ordem de 10¹² em um tempo razoável. Neste ínterim se faz necessário uma revisão de alguns conceitos importantes de matemática, que são extremamente úteis no contexto da programação.


# 2. Aritmética modular - fundamentos matemáticos{#2-aritmética-modular-fundamentos-matemáticos}

Em matemática, o termo aritmética é usado para designar o estudo de certas funções matemáticas que recebem como argumentos números inteiros e retornam valores que são subconjuntos do números complexos (podendo ser inteiros, racionais, reais, ou números complexos propriamente ditos). Por exemplo, as operações aritméticas elementares (soma, subtração, multiplicação, divisão) se enquadram nessa categoria pois associam a um par de números inteiros outro número inteiro ou racional (no caso de divisões não exatas).

De particular interesse para a solução do problema proposto no início do texto é a chamada aritmética modular. Ao se efetuar uma divisão o resultado pode ser zero se a divisão for exata ou maior que zero caso contrário. Em particular, é costumeiro na matemática usar a notação a ≡ c (mod b) ou a (mod b) ≡ c, para denotar que o resto da divisão de um inteiro 'a' por um inteiro 'b' é igual a 'c'. Por exemplo, podemos escrever usando essa notação: 10 ≡ 1 (mod 3), 9 ≡ 0 (mod 3), 8 ≡ 2 (mod 3), 7 ≡ 1 (mod 3), e assim por diante. 

Usualmente em aritmética modular se trabalha com valores positivos, porém notando-se que (-1)+3=2 e (-2)+3=1, poderiamos escrever as relações dos exemplos acima como 10 ≡ -2 (mod 3), 9 ≡ 0 (mod 3), 8 ≡ -1 (mod 3), 7 ≡ -2 (mod 3) e assim por diante.

É interessante observar que a (mod b) assume valores no intervalo {0,1,..., (b-1)}. Em matemática, se diz-se formalmente que estamos trabalhando não mais com o conjunto de inteiros Z, mas sim com Z/bZ inteiros a menos de uma relação de congruência módulo 'b'. 

Uma operação de redução modular divide o conjunto de inteiros em classes (o leitor eventual pode pensar em uma gaveta que separa os números inteiros em compartimentos) de acordo com o resto da operação da divisão (tecnicamente se diz uma classe residual) por um inteiro 'b'. No caso de b=3 temos as seguintes classes de inteiros:

```
n ≡ 1 (mod 3) [o resto da divisão por 3 é 1]: ...,-5, -2, 1, 4, 7, 10, 13, 16,...
n ≡ 2 (mod 3) [o resto da divisão por 3 é 2]: ...,-4, -1, 2, 5, 8, 11, 14, 17,...
n ≡ 0 (mod 3) [o resto da divisão por 3 é 0, são os múltiplos de 3]: ...,-9, -6, -3, 0, 3, 6, 9, 12, 15, 18, ...
```

Z/3Z é o conjunto {0,1,2} ou {-2, -1, 0}, cada elemento desse conjunto define uma das classes residuais acima.

Para o propósito de se resolver o problema de determinar de forma eficiente quadrados perfeitos necessitamos da seguinte proposição matemática.


***PROPOSIÇÃO: Sejam 'n' um quadrado perfeito. Para todo 'k' inteiro n (mod k) também é um quadrado perfeito módulo 'k', isto é, existe um 'x' inteiro tal que n (mod k) = x² (mod k).***

Não nos interessa aqui provar tal proposição, vamos apenas ilustrá-la com um exemplo muito importante na resolução do problema inicial proposto.

Seja b=16=4², para os quadrados perfeitos até 100 temos que:  1 (mod 16) ≡ 1=1², 4 (mod 16) ≡ 4=2², 9 (mod 16) ≡ 9=3², 16 (mod 16) ≡ 0=0², 25 (mod 16) ≡ 9=3², 36 (mod 16) ≡ 4=2², 49 (mod 16) ≡ 1=1², 64 (mod 16) ≡ 0=0², 81 (mod 16) ≡ 1=1², 100 (mod 16) ≡ 4=2². O resto da divisão de um quadrado perfeito por 16 é sempre 0,1,4 ou 9 que por sua vez também são quadrados perfeitos.

# 3. Aritmética modular e linguagens de programação{#3-aritmética-modular-e-linguagens-de-programação}

A aritmética modular é uma operação matemática muito importante em uma enorme gama de contextos na área de programação de forma que muitas linguagens de programação dispõem de um operador matemático específico para representá-la. Linguagens como Python e C/C++ fazem uso do operador % para representar a operação de redução modular.
 
  Por exemplo, em Python, podemos escrever o exemplo anterior como:

```python
    n=1%16
    print(n)
  
    n=4%16
    print(n)

    n=9%16
    print(n)

    n=16%16
    print(n)

    n=25%16
    print(n)

    n=36%16
    print(n)

    n=49%16
    print(n)

    n=64%16
    print(n)
  
    ....................
```

Em outras linguagens de programação como Fortran e Haskell existe uma função matemática, geralmente chamada mod() que executa tal operação, esta função recebe dois argumentos inteiros 'a' e 'b'. Tal função, em geral, está disponível por padrão para uso. A sintaxe usada para empregar tal função em um programa costuma para variar para cada linguagem de programação.


# 4. De volta ao problema original: uma solução aprimorada{#4-de-volta-ao-problema-original-:-uma-solução-aprimorada}

O último exemplo da parte 2, nos dá uma ideia de como podemos otimizar o algoritmo proposto na parte 1 do texto para se resolver o problema de detectar quadrados perfeitos. Usando o fato de que o resto da divisão de um quadrado perfeito por 16 é sempre 0, 1, 4 ou 9 podemos "filtrar" possíveis candidatos. A divisão por 16 pode dar resto no intervalo {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}. Dos 16 valores possíveis temos de checar apenas 4, {0,1,4,9}, de forma que nesse processo selecionamos apenas 25% dos elementos no conjunto dos inteiros. Tal fato reduz consideravelmente o custo de se usar a função sqrt() para computar a raiz quadrada de um número inteiro, o que em geral é uma operação relativamente "pesada" do ponto de vista do esforço computacional requerido se executado repetidas vezes.

O algoritmo final então fica assim:

```
  ENTRADA: Um inteiro 'n'
  SAÍDA: Verdadeiro se 'n' for um quadrado perfeito ou falso caso contrário

  1. (Pré-teste) Computar m = n (mod 16). Se m ≠ {0,1,4,9} retorne falso.
  2. Compute r=floor(sqrt(n)).
  3. Se r²=n retorne verdadeiro, caso contrário retorne falso.
```

O primeiro passo da etapa anterior, pode ser implementado de forma extremamente eficaz usando-se um array (lista) de 16 valores 0 ou 1. O array deve ser organizado de forma que os 1's estejam nas posições 0,1,4 e 9 do array, ao passo que todos os demais elementos são 0. Basta computar n (mod 16) e checar se no array tal índice aloca um valor 0 ou 1.

  Em Python, por exemplo, podemos usar os comandos:

```python
    mod16:list=[1,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
    m=mod16[n%16]
  
    if(m==1):
      .........................................       #Resto do código
```

Usando a dica acima, fica para o leitor o exercício de implementar o algoritmo descrito.


# 5. Operadores de manipulação de bits e representação binária{#5-operadores-de-manipulação-de-bits-e-representação-binária}

Nos computadores modernos números são representados usando-se a chamada notação binária na qual números inteiros (e números reais seguindo um procedimento ligeiramente mais complexo) são expressos em termos de potências de 2.

As potências de dois são: 2⁰=1, 2¹=2, 2²=4, 2³=8, 2⁴=16, 2⁵=32, 2⁶=64, 2⁷=8, etc. Na representação binária cada potência de 2 inferior a um dado valor a ser escrito representa um algarismo '0' (se tal potência não ocorrer na expansão do número em questão) ou '1' (se tal potência ocorrer na expansão do número em questão). A ideia básica para se escrever um inteiro na notação binária consiste em expandir o número em questão em potências de 2, teoremas garantem que para qualquer inteiro tal procedimento é sempre possível e tal expansão é sempre única.

Vamos dar alguns exemplos:

```
0=0 -->(binário) 0
1=0+1 -->(binário) 1
2=0+2 -->(binário) 10
3=2+1 -->(binário) 11
4=4+0+0 -->(binário) 100
5=4+0+1 -->(binário) 101
7=4+2+1 -->(binário) 111
13=8+4+0+1 -->(binário) 1101
16=16+0+0+0+0 -->(binário) 10000
67=64+0+0+0+2+1 -->(binário) 1000011
71=64+0+0+0+4+2+1 -->(binário) 1000111
100=64+32+0+0+4+0+0 -->(binário) 1100100
```

A notação binária é extremamente útil na criação do design de um computador, pois simplifica enormemente a representação de números inteiros usando algum sinal físico. É necessário apenas que o sinal que irá representar cada digito binário tenha dois valores específicos que associamos com '0' ou '1'. No caso de computadores eletrônicos modernos são usados circuitos elétricos miniaturizados, que aplicam uma voltagem especifica para cada bit '0' ou '1' de um inteiro. Outros modelos de representação de inteiros requerem uma engenharia muito mais complexa para representação e manipulação de valores numéricos em um computador.

Toda operação aritmética em um computador que usa representação binária é executada por meio da manipulação de bits representando números. Existem um número reduzido de operações básicas de manipulação de bits que podem ser diretamente implementadas em um computador que usa algum sinal físico para representação de dígitos binários, estas operações elementares por sua vez podem ser combinadas a fim de representar operações aritméticas com maior complexidade como as operações de soma, subtração, multiplicação e divisão de números.

Entre as operações elementares de manipulação de bits, as mais básicas são as operações: NOT, AND e OR. Tais denominações advém do chamado cálculo proposicional em lógica matemática, no qual se trabalha com proposições (variáveis) que podem assumir apenas dois valores 'verdadeiro' e 'falso'. Variáveis podem ser combinadas em sentenças mais complexas que por sua vez podem ser falsas ou verdadeiras. Essas operações elementares podem ser representadas por operadores unários (NOT) ou binários (AND, OR, etc).

Algumas linguagens de programação como Python e C/C++ disponibilizam estes operadores de manipulação de bits por padrão, em outras linguagens como Haskell, tais operações requerem uma biblioteca externa. Em C/C++ e Python tais operadores de ''manipulação de bits'' são designados pelos símbolos: ~ (NOT), &(AND), |(OR).

A tabela a seguir lista os resultados de tais operações em um único bit.

```
    a     b     (a AND b)
-------------------------------
    0     1         0
    1     0         0
    0     0         0
    1     1         1
```

```
    a     b     (a OR b)
-------------------------------
    0     1         1
    1     0         1
    0     0         0
    1     1         1
```

```
    a          (NOT a)
-------------------------------
    0             1
    1             0
```


Quando um número é representado por mais de 1 bit tais operações são aplicadas a cada bit individualmente. Vamos ilustrar as operações acima usando os números 5 e 7 que em notação binária são escritos como: 

```
5=4+0+1 --> 101
7=4+2+1 --> 111
```

```
(5&7)=5, pois 

      (1  0  1)
   AND
      (1  1  1)
---------------------
       1  0  1
```

```
(5|7)=7, pois

      (1  0  1)
    OR
      (1  1  1)
---------------------
       1  1  1
```

```
(~5)=2, pois

  NOT (1  0  1)
---------------------
       0  1  0
```

Em um computador cada operação de manipulação de bit corresponde a uma instrução simples da máquina, em matemática se usa a notação O(1) para denotar que tal operação é feita em um tempo fixo (que depende do Hardware e equivale ao tempo necessário para 'processar' um único bit). Operações mais complexas como a multiplicação de dois inteiros tem maior complexidade. No caso da multiplicação de dois inteiros de 'n' bits a complexidade é O(n²), tal notação indica que neste caso são requeridos em torno de n² operações básicas de manipulação de bits para se obter o resultado final.


# 6.  De volta ao problema original: uma solução aprimorada parte 2{#6-de-volta-ao-problema-original-:-uma-solução-aprimorada-parte-2}

Vamos enunciar sem demonstração a seguinte proposição matemática.


***PROPOSIÇÃO: Seja 'n' e 'k' inteiros tais que n=(2^k). Se 'a' e 'b' forem inteiros escritos em notação binária então a = b (mod n) equivale a  equação a = {b&(n-1)}, onde & denota a operação de manipulação de bits AND.***


Para ilustrar o significado da proposição acima sejam em notação binária, os seguintes números (bits '0' no inicio de uma string binária foram acrescidos apenas para que todas strings binárias tivessem o mesmo número de bits, mas não são significativos nos valores representados):

```
37 = 100101
25 = 011001
16 = 010000
15 = 001111
9  = 001001
5  = 000101
```

Observemos que o número 16 em binário 010000 possui apenas um bit 1 seguido de 4 bits zero. Por esta razão, quando calculamos para algum 'n' inteiro o valor de n (mod 16), o que estamos fazendo na verdade é considerar apenas os 4 bits menos significativos do número em questão. Por exemplo: 25 (mod 16) = 9, em binário 25 é 011001 e 9 é 001001. Desconsiderando os dois primeiros bits (01) de 25 em notação binária vemos que os quatro bits menos significativos de 25 e 9 são exatamente os mesmos:

```
    25 = 01 1001
    9  = 00 1001
```

Ao computar (25&15) temos

```
       (01 1  0  0  1)
   AND
       (00 1  1  1  1)
---------------------------------
        00 1  0  0  1
```

O resultado obtido equivale a string binária de 9 (ela preservou os quatro bits menos significativos, todos os outros são zerados).
De fato 25 (mod 16) = 9 (25 dividido por 16 dá resto 9).

Da mesma forma, ao computar (37&15) temos

```
       (10 0  1  0  1)
   AND
       (00 1  1  1  1)
--------------------------------
        00 0  1  0  1
```

O resultado obtido equivale a string binária de 5 e de fato 37 (mod 16) = 5 (37 dividido por 16 dá resto 5). 

Tal resultado é extremamente útil, pois para valores do tipo n=2^k, com 'k' inteiro, podemos reduzir qualquer valor módulo 'n' em apenas uma única operação de manipulação de bits, ao passo que uma operação de divisão requerida para a redução módulo 'n' é consideravelmente mais dispendiosa do ponto de vista computacional.

O algoritmo em sua forma final continua o mesmo:

```
  ENTRADA: Um inteiro 'n'
  SAÍDA: Verdadeiro se 'n' for um quadrado perfeito ou falso caso contrário

  1. (Pré-teste) Computar m = n (mod 16). Se m ≠ {0,1,4,9} retorne falso.
  2. Compute r=floor(sqrt(n)).
  3. Se r²=n retorne verdadeiro, caso contrário retorne falso.
```

A novidade agora é que usando-se o artifício proposto nesta seção a primeira etapa do algoritmo pode ser executada de forma extremamente eficaz numa única operação de manipulação de bits, isto é, em tempo constante para cada valor a ser testado. Desta forma podemos eliminar ¾ (75%) dos possíveis candidatos de forma extremamente rápida.

Em Python, por exemplo, o código em sua forma final pode ser escrito como:

```python
    mod16:list=[1,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
    m=mod16[n&15]
  
    if(m==1):
      .........................................       #Resto do código
```

Como exercício final para o leitor fica a tarefa de implementar o algoritmo acima usando as otimizações sugeridas no texto e listar, por exemplo, todos os quadrados  perfeitos até 10000000000.

Como observação final destacamos que vários algoritmos similares a este existem, embora todos sejam fundamentados na proposição listada na parte 2. Usando esta proposição e uma série de listas (arrays) de valores pré-computados de classes residuais módulo inteiros pré-definidos é possível reduzir a complexidade da busca por quadrados perfeitos descartando compostos que não satisfazem a proposição citada (tal proposição é justificada rigorosamente em teoria dos números). O algoritmo aqui proposto é dos mais simples baseados no conteúdo desta proposição.


# 7. Conclusão{#7-conclusão}

O problema inicial acima e toda discussão realizada mostra a importância da matemática para a área da programação. Sobretudo a utilidade da mesma na elaboração e no aperfeiçoamento do tempo de execução de algoritmos. Tarefas computacionais extremamente dispendiosas podem ser executadas de forma mais eficaz usando-se um formalismo matemático adequado ao problema a ser tratado.

Reduzir o tempo de execução de sub-rotinas frequentemente usadas é extremamente importante, sobretudo em sistemas e aplicações de grande escala na qual várias tarefas simples (de baixa complexidade) são executadas repetidamente. Esperamos que o texto sirva como exemplo de como a matemática pode ser uma poderosa aliada neste sentido.

# 8. Bibliografia{#8-bibliografia}

O presente artigo baseou-se nas seguintes referências:

- Introduction to Analytic Number Theory, (1976) Springer-Verlag, New York. ISBN 0-387-90163-9

- Cohen, Henri (1996). A Course In Computational Algebraic Number Theory. Graduate Texts in Mathematics. Vol. 138. Springer-Verlag. ISBN 0-387-55640-0. 3rd, corrected printing; 2nd correct. print 1995; 1st printing 1993[1]

- A Course in Number Theory and Cryptography. Graduate Texts in Mathematics. Vol. 114 (Second ed.). New York: Springer-Verlag. doi:10.1007/978-1-4419-8592-7. ISBN 978-1-4419-8592-7.

