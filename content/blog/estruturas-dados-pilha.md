+++
 categories = ["algoritmo"]
 date = "2017-09-10T19:42:00-03:00"
 tags = ["osprogramadores", "estruturas de dados", "Stack", "Pilhas"]
 title = "Estrutura de dados: pilha"
 banner = "img/banners/pilha-pratos.jpg"
+++

# Estruturas de dados: Pilha.

Pilha é uma estrutura de dados muito comum em sistemas computacionais. Dentre as várias soluções possíveis que a pilha permite podemos citar: inversão de listas; armazenar dados; implementar LIFOs.


### LIFO (Last In First Out, Último a entrar, primeiro a sair).
Um **LIFO** é um conceito computacional simples, quer dizer os mais recentes primeiro. Em uma pilha de pratos por exemplo, ao adicionarmos um prato à pilha o colocamos no topo, e ao retirar sai o do topo também.

<p><a href="https://commons.wikimedia.org/wiki/File:Lifo_stack.png#/media/File:Lifo_stack.png"><img src="https://upload.wikimedia.org/wikipedia/commons/b/b4/Lifo_stack.png" height="300" alt="Lifo stack.png"></a><br>Por Maxtremus, <span class="int-own-work" lang="pt">Obra do próprio</span>, <a href="http://creativecommons.org/publicdomain/zero/1.0/deed.en" title="Creative Commons Zero, Public Domain Dedication">CC0</a>, <a href="https://commons.wikimedia.org/w/index.php?curid=44458752">Ligação</a></p>

A utilização dessa estrutura é natural onde temos cenários que a ultima operação empilhada seja a primeira a ser processada, em outras palavras, as mais recentes primeiro.

Um exemplo prático é a utilização em logística de produtos que não tem vencimento, isso traz algumas vantagens para essas empresas pois geralmente as empresas que operam com esse modelo de negócio precificam a mercadoria com base na última compra, isso faz com que os produtos mais antigos em estoque que porventura fossem comprados com custo menor tenham margem de lucro maior, pois serão vendidos com preço de venda maior baseado na compra mais recente.

Outro uso é a inversão de outras estruturas de dados, como pilhas, listas e *arrays*. Como os mais recentes saem primeiro, se processarmos a estrutura desejada desde o primeiro elemento o último será o mais recente na pilha, só precisamos então desempilhar todos os elementos que nossa estrutura alvo estará invertida.

### Conceitos
Vamos comparar nossa estrutura a uma pilha de pratos. Cada item na pilha conhece tanto o de cima quanto o abaixo **imediatamente** a esse. Com isso em mente podemos fazer uma analogia e programaticamente definir um algoritmo que empilhe e desempilhe os pratos.

Se estivermos em um restaurante e formos preparar nosso buffet teremos que empilhar os pratos para otimizar o espaço e permitir que os clientes comam. Para empilhar precisamos saber duas coisas

1. Onde está o prato anterior 
2. Quantos pratos nossa pilha suporta

Como o segundo item é uma preocupação que pode ser relaxada, afinal, a quantidade de itens que podemos empilhar é definida pela quantidade de memória disponível não vamos nos ater a esse detalhe na implementação.

A preocupação principal que devemos ter está então no item 1, onde está o prato anterior, temos 2 cenários possíveis, não existe pratos na pilha, e já existem pratos na pilha.

### Implementação
Vamos criar 2 tipos para auxiliar na abstração da estrutura.

1.  Tipo `Stack` genérico
2.  Tipo `StackItem` genérico

A razão dos tipos serem genéricos é permitir uma implementação mais simples, pois não haverá necessidade de fazer casting entre os tipos.

O código a seguir é uma implementação simples e efetiva de uma estrutura de dados do tipo pilha.

```csharp
public class Stack<T>
    {
	//Esse é elemento responsável por linkar os itens da pilha
	private StackItem<T> topo;
        private bool isListaVazia
        {
            get { return this.Count == 0; }
        }
        public int Count;

        //Método que adiciona itens à pilha
        public void Push(T item)
        {
            if (isListaVazia)
                //Cria uma pilha com um único elemento
                this.topo = new StackItem<T>(item);
            else
            {
                var temp = topo;
                //Faz um link entre o elemento empilhado com o do topo
                var stackItem = new StackItem<T>(temp, item);
                //Define que o elemento empilhado será o novo topo
                this.topo = stackItem;
            }
            this.Count++;
        }

        //Método que desempilha, observe o retorno.
        public T Pop()
        {
            if (isListaVazia) 
                throw new IndexOutOfRangeException();
            else 
            {
                var item = this.topo.item; //Obtem o valor armazenado no topo
                this.topo = this.topo.anterior; //Define que o topo será o item anterior, ou nulo.

                this.Count--;
                return item; //Retorna um objeto do tipo genérico T
            }
        }
 
        //Inner Class para dar suporte à pilha e ajudar na abstração
        private class StackItem<T>
        {
            public T item; //Armazena o valor
            public StackItem<T> anterior; //Link com o objeto anterior

            public StackItem(T item)
            {
		//Cria uma pilha sem vínculo
                this.item = item;
            }

            public StackItem(StackItem<T> anterior, T item)
            {
                //Faz um link entre o topo e o item
                this.anterior = anterior;
                this.item = item;
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
        //Adicionando 2 objetos anônimos a pilha
        stack.Push(new
        {
            a = 1,
            b = 2
        });
        
        stack.Push(new
        {
            A = 27,
            B = 28
        });
		
	//Desempilhando itens
	Console.WriteLine(stack.Pop());
        Console.WriteLine(stack.Pop());
	}
}
```

A grande sacada está na implementação do tipo `StackItem` e no seu constructor que recebe 2 parâmetros. Como o primeiro parâmetro é do tipo `StackItem` fazemos um vinculo entre os tipos, de forma que um `StackItem` sempre tenha outro `StackItem` internamente, aqui fazemos com que o prato do topo conheça o prato abaixo de si, afinal todos os pratos são `StackItem`, além disso um `StackItem` também tem um valor, que é do tipo genérico.

Ao analisarmos a implementação dos métodos Push e Pop veremos como é feito o vinculo e desvinculo entre os `StackItem`. 

Ao acrescentar um item criamos um tipo `StackItem` e se a nossa lista estiver vazia podemos colocar o prato em qualquer lugar, se a nossa mesa já tiver pratos temos que ter o cuidado de posicionar o prato sob o anterior, para isso obtemos o `StackItem` anterior, que até então é o topo, fazemos ele se vincular ao novo prato que será empilhado e definimos que o item será o novo topo.

Para que o cliente se sirva é necessário haver pratos na pilha, então verificamos se temos algum prato, e se tivermos armazenamos o prato temporariamente até definirmos o novo topo, após isso entregamos o prato que estava em uma das mãos para o cliente possa se servir.

O código funcional pode ser baixado através do [link](https://gist.github.com/cfguimaraes/ea97ed1030ca319bb19289afe5c9b8c2)
