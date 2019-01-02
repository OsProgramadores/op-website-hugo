+++
 categories = ["algoritmo"]
 date = "2017-09-10T19:42:00-03:00"
 tags = ["osprogramadores", "estruturas de dados", "Stack", "Pilhas"]
 title = "Estrutura de dados: pilha"
 banner = "img/banners/pilha-pratos.jpg"
+++

## Estruturas de dados: Pilha

Pilha é uma estrutura de dados muito comum em sistemas computacionais. Dentre as várias soluções possíveis que a pilha permite podemos citar:

- Inversão de listas
- Armazenar dados
- Implementar LIFOs.

## LIFO (Last In First Out, Último a entrar, primeiro a sair)
Um **LIFO** é um conceito computacional simples, significa que os elementos adicionados mais recentemente ao LIFO serão os primeiros a serem removidos. Em outras palavras, os mais recentes primeiro.

Em uma pilha sempre adicionamos um item ao seu topo, e sempre que retiramos também é o do topo, em meios normais claro. Uma pilha é naturalmente um **LIFO**

![Exemplificação de LIFO](/img/conteudos-de-artigos/lifo_stack.png)

A utilização dessa estrutura é natural onde temos cenários que a ultima operação empilhada seja a primeira a ser processada, em outras palavras, as mais recentes primeiro.

Um exemplo prático é a utilização em logística de produtos que não tem vencimento, isso traz algumas vantagens para essas empresas pois geralmente as empresas que operam com esse modelo de negócio precificam a mercadoria com base na última compra, isso faz com que os produtos mais antigos em estoque que porventura fossem comprados com custo menor tenham margem de lucro maior, pois serão vendidos com preço de venda maior baseado na compra mais recente.

Outro uso é a inversão de outras estruturas de dados, como pilhas, listas e *arrays*. Como os mais recentes saem primeiro, se processarmos a estrutura desejada desde o primeiro elemento o último será o mais recente na pilha, só precisamos então desempilhar todos os elementos que nossa estrutura alvo estará invertida.

## Conceitos
Vamos comparar nossa estrutura a uma pilha de pratos. Cada item na pilha conhece tanto o de cima quanto o abaixo **imediatamente** a esse. Com isso em mente podemos fazer uma analogia e programaticamente definir um algoritmo que empilhe e desempilhe os pratos.

Se estivermos em um restaurante e formos preparar nosso buffet teremos que empilhar os pratos para otimizar o espaço e permitir que os clientes comam. Para empilhar precisamos saber duas coisas

1. Onde está o prato anterior
2. Quantos pratos nossa pilha suporta

Como o segundo item é uma preocupação que pode ser relaxada, afinal, a quantidade de itens que podemos empilhar é definida pela quantidade de memória disponível não vamos nos ater a esse detalhe na implementação.

A preocupação principal que devemos ter está então no item 1, onde está o prato anterior, temos 2 cenários possíveis:

- Não existe pratos na pilha
- Já existem pratos na pilha.

## Implementação
Vamos criar 2 tipos para auxiliar na abstração da estrutura.

1.  Tipo `Stack` genérico
2.  Tipo `StackItem` genérico

A razão dos tipos serem genéricos é permitir uma implementação mais simples, pois não haverá necessidade de fazer casting entre os tipos.

O código a seguir é uma implementação simples e efetiva de uma estrutura de dados do tipo pilha.

```csharp
using System;

public class Stack<T>
{
    private StackItem<T> topo; // Utilizado para vincular os elementos
    private bool isPilhaVazia
    {
        get { return this.Count == 0; } // Determina se temos itens na pilha
    }
    public int Count;

    // Método que empilha
    public void Push(T item)
    {
        if (isPilhaVazia)
            this.topo = new StackItem<T>(item); // Pilha contém um único elemento
        else
        {
            var stackItem = new StackItem<T>(topo, item); // Inserimos um novo topo vinculado ao antigo
            this.topo = stackItem; // Declaramos o novo topo
        }
        this.Count++;
    }

    // Método que desempilha
    public T Pop()
    {
        if (isPilhaVazia)
            throw new IndexOutOfRangeException(); // Não é possível remover itens de uma pilha vazia
        else
        {
            T item = this.topo.item; // Obtemos o valor armazenado no topo
            this.topo = this.topo.anterior; // Movemos o ponteiro do topo para o item anterior

            this.Count--;
            return item; //Retornamos o valor armazenado
        }
    }

    //Classe de suporte as operações da pilha, responsável por vincular os elementos e armazenar valores
    private class StackItem<T>
    {
        public T item; // Valor, genérico
        public StackItem<T> anterior; // Vinculo com o item anterior ou abaixo da pilha

        public StackItem(T item)
        { // Constructor de uma pilha vazia
            this.item = item;
        }

        public StackItem(StackItem<T> anterior, T item)
        { // Constructor de uma pilha com elementos
            this.anterior = anterior; // Vincula o atual com o anterior
            this.item = item; // Armazena o valor
        }
    }
}
```

O método Push geralmente é utilizado para acrescentar um item a pilha, e o método Pop é utilizado para remover um item da pilha.

```csharp
using System;

class Program
{
    public static void Main()
    {
        var stack = new Stack<object>();
        stack.Push(new
        {
            a = 1,
            b = 2
        });

        stack.Push(@"Como a pilha é genérica, podemos inserir quaisquer elementos,
                     até objetos heterogêneos");

        stack.Push(new
        {
            MadeIn = "Brazil"
        });

        stack.Push("Execute esse código para ver a pilha ser invertida");

        while (stack.Count > 0)
        {
            Console.WriteLine(stack.Pop());
        }
    }
}
```

A grande sacada está na implementação do tipo `StackItem` e no seu constructor que recebe 2 parâmetros. Como o primeiro parâmetro é do tipo `StackItem` fazemos um vinculo entre os tipos, de forma que um `StackItem` sempre tenha outro `StackItem` internamente, aqui fazemos com que o prato do topo conheça o prato abaixo de si, afinal todos os pratos são `StackItem`, além disso um `StackItem` também tem um valor, que é do tipo genérico.

Ao analisarmos a implementação dos métodos Push e Pop veremos como é feito o vinculo e desvinculo entre os `StackItem`. 

Ao acrescentar um item criamos um tipo `StackItem` e se a nossa lista estiver vazia podemos colocar o prato em qualquer lugar, se a nossa mesa já tiver pratos temos que ter o cuidado de posicionar o prato sob o anterior, para isso obtemos o `StackItem` anterior, que até então é o topo, fazemos ele se vincular ao novo prato que será empilhado e definimos que o item será o novo topo.

Para que o cliente se sirva é necessário haver pratos na pilha, então verificamos se temos algum prato, e se tivermos armazenamos o prato temporariamente até definirmos o novo topo, após isso entregamos o prato que estava em uma das mãos para o cliente possa se servir.

## Referências
Autor: Cleverton Fernandes Guimarães, clevertonfernandesguimaraes@gmail.com

[Faça o donwload do código através desse link](https://gist.github.com/cfguimaraes/4a286c64b0194e668b9e6fd86ae1f7a4)
