+++
title = "Ponteiros em C"
date = "2017-06-09T21:00:00-07:00"
tags = ["programação"]
categories = ["linguagens"]
banner = "img/banners/languages.jpg"
+++

## Introdução

O uso de ponteiros em C é um dos aspectos mais poderosos e importantes da
linguagem, e ainda assim, um dos mais confusos para os iniciantes. O objetivo
deste documento é fornecer uma  introdução básica ao funcionamento e uso de
ponteiros em C.

Importante:  Vários conceitos foram simplificados para ajudar a compreensão do
tópico principal, tais como a representação exata de números em memória,
gerenciamento de memória e outros.

## Premissas

* Usaremos inteiros (int) para a maioria das explicações. Assuma que um inteiro
  na plataforma em questão tenha 4 bytes (32 bits).

* O computador em questão tem "memória infinita". Cada byte de memória pode ser
  acessado diretamente, pelo seu endereço (em hexa, no formato 0xNNNN).

* Para simplificar o entendimento, todas as variáveis são estáticas (static)
  por default, ou seja, alocadas diretamente em memória e não na pilha (stack).

## Variáveis e uso de memória.

Antes de entender ponteiros, é importante entender o relacionamento entre variáveis
e memória. Considere o trecho de programa abaixo:

```
#include <stdio.h>

void soma(int a, int b) {
  printf("%d + %d = %d\n", a, b, a + b);
}

int main() {
  int foo;
  int bar;

  foo = 5;
  bar = 6;

  soma(foo, bar);
}

```

A declaração `int foo` declara uma variável do tipo inteiro. Quatro bytes serão
reservados em algum lugar na memória para guardar o valor dessa variável.
Imagine que o endereço de memória 0x1000 esteja disponível (a realidade é um
pouco diferente, mas a simplificação ajuda a compreensão).  Nesse caso, após
essa linha a memória agora contém:

```
    +----------------------------------------------+
    | 0x1000      |                                | endereços
    +-------------+--------------------------------+
    | ?? ?? ?? ?? |                                | valores
    +-------------+--------------------------------+
      foo
```

Observe como a posição 0x1000 foi "reservada" para a variável "foo", tomando 4
bytes. Observe também que o conteúdo é *indefinido* (representado aqui por
"??"). Um printf nessa variável nesse momento resultará em valores indefinidos
ou zero (dependendo da implementação do compilador.)

Agora, declaramos outra variável: "bar", com `int bar`:

```
    +---------------------------+------------------+
    | 0x1000      | 0x1004      |                  | endereços
    +-------------+-------------+------------------+
    | ?? ?? ?? ?? | ?? ?? ?? ?? |                  | valores
    +-------------+-------------+------------------+
      foo           bar
```

Mais quatro bytes foram alocados na memória para a variável "bar". Como a
posição 0x1000 já guarda a variável "foo", a posição 0x1004 foi escolhida
(0x1000 + 4 bytes).

A partir desse ponto, o compilador sabe que toda vez que uma referência for
feita a variável "foo", a posição de memória 0x1000 deve ser usada.  Referências
a variável "bar" vão usar a posição 0x1004.

Em seguida, o código guarda o valor 5 na variável "foo" com `foo = 5` e o valor 6
na variável "bar" com `bar = 6`. O nosso "mapa de memória" agora é:

```
    +---------------------------+------------------+
    | 0x1000      | 0x1004      |                  | endereços
    +-------------+-------------+------------------+
    | 00 00 00 05 | 00 00 00 06 |                  | valores
    +-------------+-------------+------------------+
      foo           bar
```

Observe como as posições de memória agora contém os valores de cada uma das
variáveis.

Ao chamar a função "soma", o compilador lê os nomes das variáveis a serem
passadas para a função ("foo" e "bar"). O *conteúdo* de cada uma daquelas
posições de memória é então *copiado* para a função, e a função executada. A
função recebe uma *cópia* dos valores de "foo" e "bar" em variáveis locais
chamadas "a" e "b". As variáveis "a" e "b" estão alocadas em áreas completamente
diferentes da memória.

Isso explica porque é **impossível mudar o valor das variáveis externas** de
dentro da função. Se a função "soma" modificar o valor de "a", ela na verdade está
modificando o valor de uma outra área de memória (e não 0x1000 onde os valores
originais foram guardados).

Na prática, as variáveis declaradas dentro de funções são guardadas em uma área
temporária, e desaparecem quando a função retorna.

Claramente, é impossível modificar os valores originais de dentro da função, mas
muitas vezes, é necessário fazer justamente isso. Imagine o exemplo de uma
função que aceita 2 parâmetros, adicionando 1 ao primeiro e 2 ao segundo.  Como
fazer isso em C?

A solução é passar o *endereço* da variável para uma função e acessar os valores
originais indiretamente (via ponteiros).

## Ponteiros

Modificando o programa original:

```
#include <stdio.h>

void incrementa(int *a, int *b) {
    *a = *a + 1;
    *b = *b + 2;
}

int main() {
  int foo;
  int bar;

  foo = 5;
  bar = 6;

  incrementa(&foo, &bar);
  printf("Valor de foo: %d, valor de bar: %d\n", foo, bar);
}
```

A inicialização do programa é similar, com o compilador alocando uma área em
memória para "foo" e outra para "bar" (vamos assumir novamente 0x1000 e 0x1004). A
diferença é que agora a função "incrementa" é chamada com os **endereços** de
"foo" e "bar" ao invés dos valores.

Em resumo, o operador "&var" passa o endereço da variável. Exemplo:

```
foo --> 5       Valor de foo (conteúdo da posição 0x1000)
&foo -> 0x1000  Endereço de foo na memória
bar --> 6       Valor de bar (conteúdo da posição 0x1004)
&bar -> 0x1004  Endereço de bar na memória
```

A função "incrementa" não recebe dois inteiros como a função "soma", mas dois
*ponteiros* para inteiros (`int *a, int *b`). É importante lembrar que
**variáveis guardam valores, e ponteiros guardam endereços**. Exemplos abaixo:

```
int foo = 5;    /* OK, foo = 5 */
int *pfoo;      /* OK, definido um PONTEIRO para foo, mas ainda não usado */
int pfoo = 5;   /* ERRO! Não faz sentido guardar valor num ponteiro! */
int foo = &foo; /* ERRO! Não faz sentido guardar um endereço num inteiro! */
int pfoo = &foo /* Correto! pfoo agora contem o endereço de foo */
```

Novamente:

* Variáveis sempre guardam *valores*.
* Ponteiros sempre guardam *endereços*.

## Usando ponteiros

Observe a função "incrementa":

```
void incrementa(int *a, int *b) {
    *a = *a + 1;
    *b = *b + 2;
}
```
O que acontece aqui? A função recebe 0x1000 e 0x1004 (lembre-se, os *endereços*
de "foo" e "bar" foram passados para a função com &foo e &bar, não os valores).
A partir daí, os ponteiros "a" e "b" (dentro da função) recebem os valores
0x1000 e 0x1004.

Ao usar um ponteiro, um asterisco na frente da variável declarada como ponteiro
faz com que o compilador busque o *conteúdo* do endereço, como se fosse uma
variável.  Confuso(a)? Um exemplo simples explica melhor:

```
int foo = 5;   /* foo = 5, &foo = 0x1000 */
int *pfoo;     /* declarado ponteiro para int */
pfoo = &foo;   /* pfoo = 0x1000 (pfoo "aponta" para foo) */

foo   <-- 5
pfoo  <-- 0x1000
*pfoo <-- 5
```

Repare como o asterisco na frente do ponteiro força o compilador a ler o *valor*
contido no endereço do ponteiro. Isso permite que o endereço de uma variável
seja passado para uma função e modificado de dentro da função.

No nosso caso:

1. Os valores 0x1000 e 0x1004 são recebidos por "a" e "b" (int \*)
2. "a" contém o endereço de "foo" (0x1000)
3. "b" contém o endereço de "bar" (0x1004)
4. Ao executar `*a = *a + 1`, o compilador pega o valor de "a" (0x1000, já que
   "a" é um ponteiro e recebeu o endereço de "foo"). Como usamos `*a`, o
   compilador pega o *valor* naquele endereço (5, que é a área de memória onde
   "foo" foi originalmente alocado), e adiciona 1. O valor é colocado de volta
   na área de memória apontada por "a" (0x1000) que é efetivamente a localização
   onde "foo" original foi guardado. O efeito final é modificar a memória onde
   "foo" foi guardado, efetivamente modificando o valor de "foo".
5. O mesmo processo se repete com a variável "b".

Ainda confuso(a)?

Existe outra maneira de pensar em ponteiros:

Imagine o seguinte caso:

```
int foo = 5;
int *pfoo = &foo;
```

O ponteiro é usado da seguinte forma:

```
   pfoo ou &foo          *pfoo ou foo
      0x1000  --------->      5
```

Casos muito mais complexos são possíveis, como ponteiros para ponteiros. É
importante lembrar que alguns tipos (como estruturas) possuem sintaxes especiais
para acesso via ponteiros. Em qualquer dos casos, a regra básica é:

Se:

```
ponteiro_var == &var
```

Então:

```
*ponteiro_var == var
```

## Conclusão

A explicação neste texto é extremamente simplificada. O uso de ponteiros é
extremamente importante em C, não só para permitir que funções modifiquem
valores externos, mas também para manipular vetores e tipos mais complexos.
