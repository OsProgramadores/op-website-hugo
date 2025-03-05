+++
date = 2025-03-05T19:12:21-00:00
title = "Introdução Algoritmos"
categories = ["iniciantes","algoritmo"]
tags = ["artigos","algoritmos"]
banner = "img/banners/hand-lens.jpg"
+++

Olá, sejam bem vindos a uma série de artigos a respeito de algoritmos, essa série tem como principal objetivo explorar a respeito dos algoritmos, e será um compilado das minhas anotações e entendimento a respeito do livro [Entendendo Algoritmos](https://g.co/kgs/oA1Cfgg), então vamos lá.

## O que é um algoritmo?

Antes de mais nada, devemos entender o que é um algoritmo. Em poucas palavras um algoritmo é um conjunto de instruções para realizar uma tarefa. Um exemplo simples é a receita de um prato. Pegue um miojo, por exemplo: atrás da embalagem existe um passo a passo a ser seguido até o seu miojo estar pronto, podemos dizer que é um "algoritmo culinário" *risos*.

## Clareando ideia

Vamos utilizar um simples jogo como exemplo, será o jogo de adivinhe o número, ele funciona da seguinte maneira:

* Uma pessoa pensa em um número.
* A outra pessoa tenta adivinhar fazendo palpites.
* Se o palpite for *maior* que o número, a pessoa dirá "Muito alto", e caso seja *menor* será dito "Muito baixo"

Um jogo bem simples, para este exemplo vamos considerar de 1 a 100, e testar dois algoritmos diferentes para resolver este jogo (dois passo a passo).

### Pesquisa simples

Algo que poderíamos fazer seria tentar número por número até acertarmos o número da pessoa, se a pessoa escolher o número 50, nós teríamos feito 50 tentativas, pois teríamos ido do 1 ao 50, o algoritmo seria basicamente esse:

1. Começamos a contagem dos números no primeiro número, que no nosso caso é o 1
2. Perguntamos a pessoa se o número que escolhemos é o numero dela, caso seja nós ganhamos!
3. Caso não seja o número, nós contamos mais um número e repetimos o passo 2

Esse é um método funcional, mas se pensarmos bem iremos encontrar um grande problema, pense comigo, se o número escolhido da pessoa for o 70, nós teriamos que ter feito 70 tentativas:

```mdx
Tentativa  | Número Escolhido| Resposta
-----------|-----------------|------------
1ª         | 1               | Muito baixo
2ª         | 2               | Muito baixo
3ª         | 3               | Muito baixo
...        | ...             | ...
69ª        | 69              | Muito baixo
70ª        | 70              | Acertou! 🎯
```

Se fosse o número 100, pior ainda, só iriamos acertar na última tentativa, até que funciona mas é uma baita canseira!

### Pesquisa binária

Agora, e se nós pudéssemos reduzir esse tempo sempre tentando na metade? nós podemos seguir um algoritmo (passo a passo) mais eficiente:

1. Começamos com o número do meio, ou seja, 100÷2=50
2. Se acertarmos o palpite, podemos encerrar o jogo
3. Se o palpite estiver alto, nós podemos ignorar a metade superior e o palpite atual, e repetimos o mesmo processo na metade inferior.
4. Se o palpite estiver baixo, nós podemos ignorar a metade inferior e o palpite atual e continuarmos na metade superior.
5. Repetimos até encontrar o número da pessoa.

Esse é um método muito mais esperto, pois estamos sempre cortando pela metade ao invés de procurar um por um, se o número escolhido da pessoa for 50, iriamos encontrar na primeira tentativa, caso fosse 25, na segunda tentativa, e caso 70, como no exemplo acima seguiria este formato:

```mdx
Tentativa | Número Escolhido | Resposta     | Próximo Intervalo
----------|------------------|--------------|-------------------
1ª        | 50               | Muito baixo  | 51 a 100
2ª        | 76               | Muito alto   | 51 a 75
3ª        | 63               | Muito baixo  | 64 a 75
4ª        | 70               | Acertou! 🎯  | -
```

Foram apenas 4 Tentativas contra 70 tentativas do algoritmo anterior, muito mais eficiente.Um detalhe a se observar é que na segunda tentativa ao invés de 75 foi o número 76, isso porque eu decidi uma abordagem de arredondamento para cima, caso arredondasse para baixo também funcionaria porém teria resultados diferentes. Se fosse 1000 números, nós precisaríamos de no máximo apenas 10 tentativas.

## Codificando os Algoritmos

Irei utilizar duas linguagens como exemplo para estes algoritmos e para os futuros algoritmos, que serão Javascript e Rust (yeeeeah, Rust, mas não irei adicionar diretamente no artigo, e sim em um link, para o artigo não ficar muito extenso).

Os motivos são, tenho mais familiaridade com JavaScript e acredito que grande parte dos futuros leitores também. Já Rust é uma linguagem da qual desejo aprender, então aproveitarei está série como uma oportunidade para aprender.

 Para os primeiros artigos irei prezar por simplicidade de código então é provável que ao invés de um `switch case`, utilize apenas `if` e `else`, portanto caso você seja mais experiente tente não se desesperar, e se for iniciante, irá compreender o código com mais facilidade, agora bora lá!

### Pesquisa Simples

Em javascript esta seria uma possível implementação de uma pesquisa simples:

```ts
function pesquisa_simples(lista, item){
  let inicio = 0;
  let fim = lista.length - 1;
  let indiceCorrente = inicio
  let contador = 0;

  while (indiceCorrente <= fim){
    contador++
    let palpite = lista[indiceCorrente]

    if (palpite == item){
      console.log(`Você acertou! palpite ${palpite}\nForam necessárias: ${contador} tentativas`);
      return palpite
    } else {
      indiceCorrente++
    }
  }

  console.log(`Número não encontrado.\nTentativas realizadas: ${contador}`);
  return null
}

//Função simples para gerar uma lista
function intervalo(limite) {
  let minhaLista = [];
  let  indice = 0;

  while (indice <= limite) {
    minhaLista.push(indice);
    indice++;
  }
  return minhaLista;
}
  pesquisa_simples(intervalo(100), 70);
```

Criamos uma lista e buscamos sequencialmente através dela, a saída que teriamos seria esta:

```ts
Você acertou! palpite 70
Foram necessárias: 71 tentativas
```

Repare que foram feitas 71 tentativas, pois nesta lista estamos incluindo o número 0, caso não incluíssemos ai seriam as 70 tentativas. *Uma possível implementação em rust seria esta : [Solução Rust](https://github.com/guiribeirodev/leetcode/blob/main/articles/osprogramadores/simple-search.rs)*

### Pesquisa Binária

Em JavaScript esta seria uma possível implementação de uma pesquisa binária:

```ts
function pesquisa_binaria(lista, item) {
  let baixo = 0;
  let alto = lista.length - 1;
  let contador = 0;

  while (baixo <= alto) {
    contador++;
    let meio = Math.ceil((baixo + alto) / 2);
    let palpite = lista[meio];

    if (palpite == item) {
      console.log(`Você acertou! palpite ${palpite}\nForam necessárias: ${contador} tentativas`);
      return palpite;
    } else if (palpite > item) {
      console.log(`Muito alto! palpite ${palpite}`);
      alto = meio - 1;
    } else {
      console.log(`Muito baixo! palpite ${palpite}`);
      baixo = meio + 1;
    }
  }

  console.log(`Número não encontrado.\nTentativas realizadas: ${contador}`);

  return null;
}

//Função simples para gerar uma lista
function intervalo(limite) {
  let minhaLista = [];
  let indice = 0;

  while (indice <= limite) {
    minhaLista.push(indice);
    indice++;
  }

  return minhaLista;
}

pesquisa_binaria(intervalo(100), 70);
```

Criamos uma lista contendo os números de 0 a 100 (101 elementos, pois estamos acrescentando o 0 também) e buscamos o número 70, a saída que teríamos seria esta:

```ts
Muito baixo! palpite 50
Muito alto! palpite 76
Muito baixo! palpite 63
Você acertou! palpite 70
Foram necessárias: 4 tentativas
```

*Caso queira ver uma possível implementação em Rust: [Solução em Rust](https://github.com/guiribeirodev/leetcode/blob/main/articles/osprogramadores/binary-search.rs)*

## Big O - O que é isso?

Vamos falar um pouco de Big O e o que ele tem a ver com isso tudo, Big O é uma notação matemática usada em ciência da computação para descrever o limite superior do crescimento de um algoritmo em termos de tempo ou espaço, conforme o tamanho da entrada (n) aumenta. Uau, espera, não se desespere, vamos aos poucos, a primeira pergunta é: "Por que eu preciso saber isso?" De maneira simplista a notação big O vai te auxiliar a ter uma ideia de quão "rápido" ou "lento" um algoritmo é, e nós estamos sempre escrevendo os nossos e usando o de outras pessoas, então é bom ter uma compreensão.

Alguns exemplos de complexidade Big O são:

* O(log n): Tempo logarítmico (exemplo: pesquisa binária).

* O(n): Tempo linear (exemplo: pesquisa simples).

* O(n * log n): Tempo linear-logarítmico (exemplo: quicksort).

* O(n²): Tempo quadrático (exemplo: selection sort).

* O(n!): Tempo fatorial (exemplo: caixeiro viajante).

### Pesquisa Simples - Big O(n)

Vamos focar nos dos primeiros casos que são o dos exemplos acima.
No caso da Pesquisa Simples nós temos um Big O(n), aonde n é o tamanho da entrada, no nosso exemplo de acertar um número entre 1 a 100, n seria a quantidade de tentativas caso nós acertássemos somente na ultima tentativa, a notação Big O vai levar em consideração o pior caso (embora também exista casos aonde iremos medir o melhor caso e o caso mediano).

No nosso exemplo n seria 100, pois teríamos que tentar 100 vezes. Em uma suposição aonde para cada tentativa gastássemos 1 milissegundo, na pior situação nós iriamos levar 10 milissegundos (0,1 segundos), caso n fosse 10000 seriam 10000ms (10 segundos), como podemos observar o tempo do algoritmo cresce de forma linear (conceito matemático de crescimento constante), olha ai a matemática nos salvando.

### Pesquisa Binária - Big O(log n)

Já na pesquisa binária o crescimento é log n, você deve se perguntar "o que é esse tal de log ?", **n** nós já sabemos que é a entrada, agora **log** é um conceito matemático de logaritmo, caso você não tenha dormido na aula já vai sacar, agora se dormiu vamos lá, logaritmo é o oposto de exponencial, um logaritmo de a na base b, é representado log logₐb, o valor x tal que a elevado a x seja igual a b. Não se desespere, um exemplo: log₂8 (lê-se logaritmo de 8 na base 2), nesse logaritmo nós estamos procurando qual número devemos elevar o 2 para que seja igual a 8.

No caso é Log₂8 = 3, pois 2³ = 8 (~~e tinham me falado que matemática n serve para nada aaaaa~~)

Voltando ao nosso exemplo, em um caso aonde temos que procurar 100 itens, a pesquisa binária é O(log₂n)=100, ou seja, eu tenho que elevar 2 a qual número caso eu deseje o resultado de 100? seria aproximadamente 6,64, neste caso 7, portanto na pior hipótese eu precisaria de 7 tentativas, em 10000 seriam apenas 14 tentativas, um baita ganho em comparação a pesquisa simples, segue a tabela abaixo com uma comparação entre os algoritmos:

```
 Tamanho da Lista (n) | Pesquisa Simples (O(n)) | Pesquisa Binária (O(log₂(n)))
----------------------|-------------------------|-----------------------------
 100                  | 100 tentativas          | ~7 tentativas
 1.000                | 1.000 tentativas        | ~10 tentativas
 10.000               | 10.000 tentativas       | ~14 tentativas
 100.000              | 100.000 tentativas      | ~17 tentativas
 1.000.000            | 1.000.000 tentativas    | ~20 tentativas
 10.000.000           | 10.000.000 tentativas   | ~24 tentativas
 1.000.000.000        | 1.000.000.000 tentativas| ~30 tentativas
```

Supondo 1ms para cada tentativa, em 1 bilhão usando pesquisa simples levaria 1 bilhão de ms que são 11,5 dias contra 30 milissegundos da pesquisa binária (talvez algoritmo seja um pouquinho importante haha)

## LeetCode

Agora vamos botar em prática o algoritmo em um [leetcode](https://leetcode.com/), algo muito útil para aprender sobre algoritmos e também cai em muitas entrevistas de emprego. Irei utilizar a plataforma Leetcode, o problema em questão será [este](https://leetcode.com/problems/search-insert-position/):

```
Posição de Inserção na Busca
Nível: Fácil

Dado um array ordenado de inteiros distintos e um valor alvo, retorne o índice se o alvo for encontrado. Caso contrário, retorne o índice onde ele deveria ser inserido para manter a ordem.

Você deve escrever um algoritmo com complexidade de tempo O(log n).

Exemplo 1:
Entrada: nums = [1,3,5,6], target = 5
Saída: 2

Exemplo 2:
Entrada: nums = [1,3,5,6], target = 2
Saída: 1

Exemplo 3:
Entrada: nums = [1,3,5,6], target = 7
Saída: 4
```

Vamos entender, o problema nos diz que temos um array ordenado e um valor que devemos buscar,caso nós encontre o valor, iremos retornar o índice dele no array, ou seja, a posição que ele se encontra no array, e caso não encontremos o nosso número procurado, devemos retornar aonde ele deveria estar, os exemplos dados são :

**Exemplo 1:**

Entrada: nums = [1,3,5,6], alvo = 5 | Saída: 2

Temos como número alvo o 5, a saída retornada deve ser 2, pois está na segunda posição do nosso array (lembre-se que o array sempre começa com índice 0).

**Exemplo 2:**

Entrada: nums = [1,3,5,6], alvo = 2 | Saída: 1

Neste caso o número 2 não se encontra no array, então é retornado a posição 1, pois seguindo a lógica, de ser um array ordenado, era para ele estar entre o 1 e 3.

**Vamos ao código agora:**

Na plataforma ao escolhermos a linguagem JavaScript, recebemos o seguinte código para iniciarmos:

```typescript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {};
```

É nos dado uma função, que irá receber dois parâmetros, nums sendo a lista de números e target o nosso número alvo, e nós devemos retornar um número. Bora lá começar a resolver:

```typescript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    let bottomIndex = 0
    let topIndex = nums.length - 1

//Até Aqui nós definimos o indíce do fundo, no caso o 0 e o ultimo indice da lista, que é para pegarmos todo a lista



    while (bottomIndex <= topIndex) {
        //Pegamos em middleIndex o indice **central**, estamos utilizando o Math.ceil para arredondar o número para cima caso o resultado da divisão não seja um número inteíro, e currentTarget é o elemento, o número que estamos testando no momento
        let middleIndex = Math.ceil((bottomIndex + topIndex) / 2)
        let currentTarget = nums[middleIndex]

        // Verificamos se o elemento atual é o número alvo, se for nós já retornamos o indíce
        if (currentTarget == target) {
            return middleIndex
        }
        // Se o elemento atual for menor que o elemento alvo, o indice do fundo será igual o indice atual que estamos testando (pois como é menor, podemos ignorar a parte inferior, e adicionamos +1 pois é o proprio indice do momento)
        else if (currentTarget < target) {
            bottomIndex = middleIndex + 1
        } 
        // Aqui entrará caso o elemento atual for maior que o elemento alvo,seguiremos a lógica reversa do if acima
        else{
            topIndex = middleIndex - 1
        }
    }

    // com o while acima nós efetivamente iremos encontrar o indíce do nosso alvo, porém nós ainda precisamos retornar o indice caso não encontremos o número, para isso basta nós retornarmos o bottomIndex, pois após o while o bottomIndex irá corresponder a posição da qual o número não encontrado deveria estar

    return bottomIndex

};
```

*Possível implementação em [Rust](https://github.com/guiribeirodev/leetcode/blob/main/leetcode/problems/search-insert-position.rs)*

E com isso nós conseguimos o seguinte resultado:

<picture>
  <source media="(max-width: 768px)" srcset="/img/conteudos-de-artigos/introducao-algoritmos/leetcode-result-binary-search-pequena.png">
  <source media="(min-width: 769px)" srcset="/img/conteudos-de-artigos/introducao-algoritmos/leetcode-result-binary-search.png">
  <img src="/img/conteudos-de-artigos/introducao-algoritmos/leetcode-result-binary-search.png" alt="Resultado leetcode">
</picture>

---
Yeeeeah, desafio concluído e com um tempo de 0ms, um algoritmo de O(log n). Espero que após este artigo vocês tenham compreendido um pouco a respeito de algoritmos, complexidade e tenha te ajudado de alguma forma.

Por hoje é isso, vejo vocês no próximo artigo!

---

## Referências

* [Entendendo algoritmos](https://g.co/kgs/oA1Cfgg) - Livro
* [Algoritmo](https://pt.wikipedia.org/wiki/Algoritmo) - Wikipedia
* [Big O](https://www.freecodecamp.org/portuguese/news/o-que-e-a-notacao-big-o-complexidade-de-tempo-e-de-espaco/) - Freecodecamp
* [Big O](https://pt.wikipedia.org/wiki/Grande-O) - Wikipedia
* [Logaritmo](https://www.todamateria.com.br/logaritmo/) - Todamateria
