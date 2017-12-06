+++
 categories = ["algoritmo"]
 date = "2017-11-06T18:27:00-03:00"
 tags = ["osprogramadores", "OOP"]
 title = "Introdução a orientação a objetos"
 banner = "img/banners/object-oriented-programming.jpg"
+++

# Object Oriented Programming 

## Proposito

Esse artigo foi construído para introduzir alguns tópicos de programação orientada a objetos, seu propósito é servir como material de consulta para novatos.

A divisão do artigo segue um conceito lógico, tende a avançar de conceitos e práticas simples para mais complexas, no decorrer do texto também haverá progessão sobre o conteúdo, dessa forma o leitor poderá avançar junto e assimilar o conteúdo.

### Estrutura do artigo

- Pilares de OOP
 1. Abstração
 2. Encapsulamento
 3. Herança e Polimorfismo
 4. Conclusão

- Conceitos de OOP
 1. Coesão e Acoplamento
 2. DRY
 2. KISS
 3. YAGNI
 4. BDUF
 5. SOC
 6. SOLID
 7. TDD

- Referências bibliográficas e materiais para consulta

> Algumas sintaxes podem não remeter à realidade, tem propósito apenas didático e visa facilitar o entendimento.

## Pilares da programação orientada a objetos

Programação Orientada a Objetos, ou do inglês Object Oriented Programming (OOP) é uma forma de escrever programas, são criadas estruturas para programaticamente imitar o comportamento das coisas (objetos) do mundo real, as coisas importantes para o programa

Toda linguagem orientada a objetos deve implementar 4 conceitos, é o que define uma linguagem como orientada a objetos, essas são:  

  - Abstração
  - Encapsulamento
  - Herança
  - Polimorfismo

Uma característica comum é que as instâncias de objetos podem ser tratados como objetos do mais baixo até o mais alto nível, inclusive para os tipos base da linguagem, como inteiros, pontos flutuantes e caracteres até a mais complexa classe de um sistema.

Em outras palavras, em uma linguagem orientada a objetos tudo deriva de objetos, geralmente do tipo `Object`  por **herança implícita**, e **tudo é conversível a objetos**.

### 1. Abstração

Do site [dicio](https://www.dicio.com.br/abstracao/), com grifo próprio
> Operação do espírito, que isola de uma noção um elemento, **negligenciando** os outros.

Em uma definição própria:
> Abstração é a capacidade de delimitar um objeto, suas propriedades e comportamentos naquilo que ele representa, somente.

Chegamos a importante conclusão: Para OOP devemos definir de forma muita clara o que o objeto representa, para o programa.

Um sistema desenvolvido utilizando o paradigma orientado a objetos é formado por estruturas que representam os objetos da vida real, que estarão presentes no sistema, representados através dessas abstrações.

Uma abstração define as propriedades, comportamentos e os eventos ocorridos a um objeto.

#### Exemplo

Um exemplo de abstração pode ser aplicado a um automóvel, Digamos que tenhamos uma classe Automóvel, quais são as características que o carro tem?

Podemos enumerar várias:

-  Fabricante; Marca; Modelo; Ano de fabricação
-  Distância entre eixos; Altura; Largura

É um erro comum descrever todas as propriedades que compõem o objeto, contudo a abstração é um exercício, devemos nos ater àquilo que é importante para o domínio da aplicação, e podemos e devemos progredir nesse exercício.

> O Domínio de uma aplicação é a definição do problema que ela visa resolver, e quais não.

```csharp
Class Automovel {
 DateTime DataDeFabricacao;
 string CodigoDeReferenciaDoFabricante;
 string Modelo;

 float Velocidade;
 Acelerar(){ }
 Frear(){ }

 Colidiu(){ }
 }
```
<small>Listagem 1: Classe automóvel</small>

Dessa forma escrevemos uma classe Automovel que tem como propriedades DataDeFabricacao, CodigoDeReferenciaDoFabricante e Modelo, ela também apresenta os comportamentos para Acelerar e Frear.

Classe é a forma mais comum de abstrair um objeto nas linguagens modernas, elas são as estruturas utilizadas para definir os comportamentos e propriedades dos objetos que podemos interagir em um programa.

Quando precisarmos de um objeto com aquele comportamento específico, nesse caso, quando precisarmos interagir com um `Automovel` não poderemos utilizar a classe, devemos criar um objeto que se comporta daquela forma, isso é o que chamamos de uma instância de classe.

Note que existe um comportamento representado pela assinatura Colidiu, contudo colidir não é um comportamento do carro. O que é então? Um evento, algo que acontece com o Automovel através de uma casualidade, ou imprevisto, geralmente são as exceções.

### Encapsulamento

Podemos pensar em encapsulamento na forma mais fácil de associação possível: Uma capsula de remédio.

Dentro de uma capsula gelatinosa estão todos os componentes que contribuem para a melhoria, todos os ingredientes estão contídos dentro da capsula, o remédio é a mistura de ingredientes, mas só ingerimos a capsula.

No enscapsulamento definimos as formas de manipulação da classe, através dos métodos publicos disponíveis, estes métodos são responsáveis por alterar o estado do objeto, realizar alguma ação, ou ainda responder a uma evento ocorrido.

Quando desejamos que o objeto se comporte de forma condicional de forma que seu comportamento aconteça em virtude do seu estado, encapsulamos esse comportamento, assim podemos validar programáticamente e aceitar ou não a manipulação.

Como exemplo, para frear um Automóvel só iremos reduzir a velocidade se essa for maior que zero, afinal não é possível para um veículo ter velocidades negativas, já para acelerar poderemos enquanto o motor tiver potẽncia.

Outra coisa importante do encapsulamento é abraçar as origens do paradigma, segundo Kay objetos devem se comunicar enviando e recebendo mensagens, portanto o estado de um objeto deve ser alterado através de um método definido na classe, isso é encapsulamento.

Algumas linguagens não permitem o acesso direto aos valores (propriedades) de uma classe, **Smalltalk** por exemplo só permite acesso através de troca de mensagens, **Java** te obrigará a implementar o mesmo princípio, são os *Getters* e *Setters*, C# apésar de não ser obrigatório implementar métodos getter/setter para as propriedades há implicitamente um método para esse propósito.

#### Exemplo

Tomemos como cenário um sistema para testes de aceleração, vamos utilizar a abstração da listagem 1, ela define a API que utilizaremos e quais são as possibilidades de consumo.

Vamos criar uma classe que implementa essa API para podermos utilizar seus recursos.
```csharp
public Class Palio : Automovel {
 this.DataDeFabricacao => DateTime.Parse("2017-06-01");
 this.CodigoDeReferenciaDoFabricante => "Fiat Palio Fire 1.0 8V (Flex) 4p 2017";
 this.Modelo => "Fire";

 Acelerar(){ 
 Velocidade = Velocidade + 6.52;e
 //Aceleração média por segundo tomando como base velocidade 0-100 Km/h em 15s
 }
 Frear(){ 
 /* Como o cálculo para frenagem é composto
    de variantes diversas vamos utilizar
    a mesma base de cálculo, 6.52 Km/h
    */
    Velocidade = Velocidade - 6.52;
    }
 
 }

 Colidiu(){ }
 }
```
<small>Listagem 2: Classe Palio</small>

```csharp
Class Program {
 public static void Main( ) {
 Automovel veiculo = new Palio(); // Instancia de um veículo implementando a API

 int segundosParaTestar = 5, // Por quanto tempo o programa ira rodar
     segundosPassadosDoTeste = 0; // Armazena o tempo percorrido pelo teste

 while (segundosParaTestar > segundosPassadosDoTeste) // Enquanto o tempo total não é atingido
 {
  veiculo.Acelerar(); // Acelera o veiculo
  Thread.Sleep(1000); // Pelo tempo de 1 segundo
  segundosPassadosDoTeste = segundosPassadosDoTeste + 1; // Incrementa em 1 segundo o tempo do teste
 }
 
 float velocidadeNoInstante = veiculo.Velocidade; // Armazena a velocidade do veiculo após o tempo desejado
 
 while (veiculo.Velocidade > 0)
  veiculo.Frear(); //Para o automóvel
 
 }
```
<small>Listagem 3: Programa para teste de aceleração</small>

De acordo com a especificação do próprio pálio, a classe Palio abstraí as informações do veículo em algo que faça sentido ao domínio da aplicação e encapsula o comportamento de aceleração e frenagem, assim o consumidor pode utilizar a API e não se preocupar em implementar o comportamento de aceleração e frenagem do pálio.

Para mudar o comportamento do programa para outro veículo é muito fácil, basta criar uma nova herança e através do polimorfismo fazer as chamadas a API disponível na classe Automovel e nosso programa pode se beneficiar do encapsulamento proporcionado e já se adequar a uma ferrari.

### Herança e Polimorfismo

Herança em OOP é a capacidade de uma estrutura (podem ser classes; interfaces; prototypes e até outros objetos) herdar algo de seus ancestrais, é uma forma de centralização e reuso.

Polimorfismo é a capacidade de uma estrutura descendente assumir a forma de um ancestral em um contexto, sem que isso interfira no resultado, é também a possibilidade de utilizar algo em diferentes contextos com outro significado.

Como podemos acompanhar na listagem 2 a classe Palio tem uma sintaxe que determina herança na linguagem Java `class Palio extends Automovel`, essa sintaxe torna Palio herança de Automovel, logo herda as propriedades e comportamentos de Veiculo.

Na listagem 3 podemos verificar o comportamento de polimorfismo na execução da instrução `Automovel veiculo = new Palio( )`. Nesse momento estamos atribuindo através de polimorfismo uma instância de Palio a um tipo Automovel.

O polimorfismo de Palio só é possível pois na construção da classe houve a declaração da herança. 

Em linguagens fortemente tipadas é comum haver restrição quanto a alguma marcação especial para determinar a herança, e logo, o polimorfismo. Em linguagens dinâmicas não há restrição para polimorfismo, qualquer coisa que implemente a API pode ser consumida por um cliente. Geralmente linguagens prototipadas tornam clara a implementação de uma herança, mas nem sempre é necessário para polimorfismo, logo, uma implementação a API é suficiente para o cliente fazer o consumo.

### Conclusão

Após percorrer os pilares de uma linguagem orientada a objetos podemos facilmente utilizar esse conceito em qualquer linguagem que apresente o paradigma, citemos C++, C#, Java, JavaScript, Smalltalk e Python como as linguagens mais famosas e comuns do paradigma orientado a objetos.

Aprendemos que:

- Abstração é a capacidade de delimitar o escopo de um objeto do mundo real através de uma estrutura de define propriedades, comportamentos e eventos ocorridos ao objeto
- Encapsulamento é uma forma de centralizar a implementação de algum comportamento, deixando o exterior sem conhecimento da implementação e suas dependências, seu propósito é definir um canal de mensagens para uma classe permitindo o controle de alteração e proporcionando simplicidade.
- Herança é a capacidade de reutilizar os comportamentos e propriedades previamente observados e definidos em outra estrutura e que Polimorfismo é a capacidade de uma estrutura alterar seu comportamento em outros contextos, ou assumir a responsabilidade de seus ancestrais.


## Conceitos de OOP

Do site [significados](https://www.significados.com.br/conceito/) podemos extrair o significado para conceito
>Conceito significa definição, concepção ou caracterização. É a formulação de uma ideia por meio de palavras ou recursos visuais.

E do site [dicio](https://www.dicio.com.br/conceito/):
>Noção; percepção que uma pessoa possui acerca de algo ou alguém.

Logo: conceito é o que percebemos de algo, como seres humanos, é algo que só existe abstratamente, pois ainda que seja algo concreto, pessoas distintas podem ter conceitos diferentes sobre o mesmo.

Conceito de automóvel é um veículo que tenha capacidade de locomoção sem atuação de forças externas, concretização disso são carros e motos, e aviões e helicópteros, contudo pessoas distintas podem divergir sobre esses grupos.

Entramos então no conceito de convenção, uma convenção é algo aceito amplamente, e por ser um conceito, é algo que pode ser ignorado, pois é um exercício humano da abstração.

Mas porquê isso é importante? Por ser tais convenções importantíssimas para o sucesso de um projeto orientado a objetos.

### Coesão e Acoplamento

Em termos OOP Coesão pode ser definido como o quanto o comportamento de uma classe esta correlacionado com a própria classe.

Relembrando o pilar da abstração, o quanto os métodos de acelear e frear correlacionam-se com um Automovel, quanto maior for esse nível maior será a coesão.

Acoplamento pode ser analogamente comparado a cola, ou o quanto é difícil separar ou juntar componentes, quanto mais cola necessária maior será o acoplamento.

Idealmente devemos buscar Alta Coesão e Baixo Acoplamento

### DRY: Dont Repeat Yourself
Esse conceito é amplamente utilizado no mundo OOP, em tradução livre: Não repita você mesmo.

Ele dá enfase ao que a OOP trouxe como proposta de melhoria: reuso, centralização, maior controle sobre o fluxo do programa.

#### Na prática

```javascript
let box1 = document.getElementById("box1");
box1.align = "left";

let box2 = document.getElementById("box2");
box2.align = "left";
```
<small>listagem 4: exemplo de código com repetição</small>

Podemos perceber no fragmento acima que temos o mesmo comportamento sendo aplicado a 2 elementos da página. Esse comportamento é obter os elementos identificados pelo id respectivo e alinhá-los a esquerda.

Aplicando o conceito de DRY, para não repetir eu mesmo posso extrair esse código duplicado que apresenta o mesmo propósito em uma função, centralizando o comportamento desejado

Significados
Significado de Conceito
O que é Conceito. Conceito e Significado de Conceito: Conceito significa definição, concepção ou caracterização. É a formulação de...
 de forma que facilite a escala e manutenção.

```javascript
function alinharBoxNaEsquerda(box) {
 let element = document.getElementById(box);
 element.align = "left";
}
alinharBoxNaEsquerda("box1");
alinharBoxNaEsquerda("box2");
```
<small>listagem 5: aplicação de DRY</small>

#### Importante lembrar

1.  O comportamento é mais importante que o código
 - Se o código em questão está duplicado, mas de forma que o comportamento seja diferenciado para cenários distintos não vale a pena, pois isso irá gerar acoplamento
 - Extraia código duplicado onde se espera que o mesmo comportamento seja aplicado em cenários distintos, isso facilita a escala e manutenção e aumenta coesão
2. Em OOP é comum construir classes para conter os comportamentos comuns, isso busca aumentar a coesão
 - Ao mesmo tempo que isso é uma prática comum é válido considerar interfaces para executar a ação, isso permite depender de abstrações e reduz o acoplamento, torna o código mais fácil de evoluir.

### KISS: Keep It Simple Stupid

Em tradução livre: Mantenha isso estupidamente simples

Muitos consideram OOP como retrocesso, a simplicidade que o paradigma proporciona disponibilizando para o desenvolvedor a capacidade de se espelhar em estruturas do mundo real, determinando contextos e comportamentos e, poder controlar o fluxo do programa através do relacionamentos entre objetos pode ser prejudicada pela *over engineering*

>*Over Engeneering* é tornar mais complexa que o necessário a solução para um problema

#### Na prática

Tomando como exemplo o comportamento esperado na listagem 4, alinhar box identificados por um id a esquerda poderíamos simplificar com outro código

```css
box-left-aligned {
align : left;
}
```
<small>Listagem 6: Alinhamento a esquerda com CSS</small>

```html
<div id="box1" class="box-left-aligned"></div>
<div id="box2" class="box-left-aligned"></div>
```
<small>Listagem 7: Aplicação de código mais simples para alinhar a esquerda</small>

Ora mas quem disse que isso é uma solução OOP?

Programadores OOP visam fazer sistemas manuteníveis, qualquer solução mais complexa que o necessário a não ser que por restrição deve ser considerada má prática, bons programadores selecionam a ferramenta certa para o trabalho.

### YAGNI: You A'int Gonna Need It
Em tradução livre: Você ainda não irá precisar disso.

Esse conceito é algo que muitos projetos não aplicam, acabam lidando com problemas de acoplamento exatamente pela criação de estruturas que suportarão futuras mudanças, contudo, essas podem nem chegar.

#### Na prática
 
Exemplo claro é a listagem 1, ela define um evento para determinar uma colisão, mas em nenhum momento foi utilizado por nenhuma das aplicações clientes,  e pior, as classes que herdam dela acabam ficando com a assinatura de algo que nunca usaram ou sequer implementaram.

### BDUF: Big Design Up Front
Em tradução livre: Toda arquitetura antes

Quando você não aplica o conceito KISS e também YAGNI você acaba com uma solução que antecipou etapas do processo, tudo porquê é previsto que o sistema atenda futuramente a demanda XPTO.

Isso é um erro, algo que já deveríamos ter aprendido é que um pouco de procrastinação é boa quando lidamos com TI, nesse caso, porquê se antecipar a algo que nem sequer temos definição? Com isso acabamos abrindo o escopo do projeto e isso é extremamento perigoso.

Se já tivéssemos a definição não seria uma previsão, e idealmente devemos levar a solução de forma simples até o momento de *stress*, um lugar onde não podemos avançar sem alterar algo, e esse algo que seja o que gera menor acoplamento com maior coesão.

Isso rotineiramente é atingido sempre que protelamos a solução ao momento que devemos aplicar, é o momento onde todos na linha do tempo terão o amadurecimento possível que outrora não poderia ter sido adquirido.

#### Na Prática
O mesmo exemplo se aplica do cenário anterior, quando fizemos *over engennering* da solução e não fizemos YAGNI acabamos nos preocupando com um cenário que descreve de forma mais completa um Automovel, contudo, gastamos tempo e esforços fazen
 do algo que não precisamos até o momento.

Quando fossemos realmente implementar essa solução em um sistema de verificação de segurança, poderíamos ter um melhor entendimento do domínio, poderíamos ter criado outro nível de herança, definindo classe AutomovelDeTeste que ao invés de um evento teria um comportamento de colisão, algo que faz total sentido para o domínio da aplicação e que nesse contexto é muito mais coeso

### SOC: Separation of Concerns
Em tradução livre: Separação de conceitos.

Isso nada mais é que a aplicação do pilar 1 descrito nesse texto, abstração.

É a definição do domínio do sistema, e quais são os componentes ou módulos que compõem o sistema.

Esse conceito também defende a responsabilidade única, onde cada componente deve ser preocupar apenas com o seu próprio contexto, sem interferir nos demais, logo, sem poluir outros conceitos.

#### Na prática

Quando definimos no conceito anterior BDUF, que poderíamos ter outra classe, que herda de Automovel e defina ainda mais comportamentos e propriedades para o domínio de um módulo de testes de segurança aplicamos o SoC, pois deixamos os clientes enxutos sem a necessidade do código desnecessário para gerenciar colisão.

Ao mesmo tempo, criamos uma estrutura mais robusta, pois temos agora mais um nível de implementações permitindo mais polimorfismos, cenários que atendam outros domínios de aplicação e também menos acoplamento, contudo, também aumentos o *codebase* que temos que gerenciar, logo devemos considerar KISS e YAGNI.

### SOLID

SOLID é um acrônimo, pode ser observado a seguir
![Imagem que descreve SOLID](https://rubygarage.org/uploads/article_image/file/578/SOLID.jpg)

<small>Figura 1: Representação de SOLID</small>

Apresenta boas práticas aplicadas para softwares orientado a objetos, é um convite a aplicação de boas práticas e de OOP real.

Por ser um assunto extenso e nada mais que OOP aplicada, faço um convite ao leitor a extensão ao assunto.

### TDD: Test Driven Development
Em tradução livre: Desenvolvimento Dirigido Por Testes

O grifo aqui é por conta da casa. Pelo entendimento da aplicação do processo pelo autor do artigo, percebemos que o intuito do conceito é isso, que o desenvolvimento seja com base na aplicação de testes de unidades.

Os teste de unidade asseguram a resposta esperada para um modulo do programa, cada componente idealmente deve ter o seu teste escrito antes da implementação do código, esse deve ser com base na assinatura esperada pelo componente e a sua especificação.

Como esse é também um assunto extenso não vamos abordá-lo em mais detalhes, fica também o convite ao leitor a extensão do conhecimento.

## Referências bibliográficas e materiais para consulta

1. [Acrônimos que todo desenvolvedor deveria conhecer e aplicar](https://thefullstack.xyz/dry-yagni-kiss-tdd-soc-bdfu/)

2. [Introdução a orientação a objetos](https://www.codeproject.com/Articles/22769/Introduction-to-Object-Oriented-Programming-Concep#Encapsulation )

3. [Página da wikipedia sobre OOP](https://en.wikipedia.org/wiki/Object-oriented_programming)

4. [SOLID (em inglês)](https://scotch.io/bar-talk/s-o-l-i-d-the-first-five-principles-of-object-oriented-design)

5. [SOLID (em português)](http://www.eduardopires.net.br/2013/04/orientacao-a-objeto-solid/)

6. [Guia prático para TDD](https://mva.microsoft.com/en-us/training-courses/testdriven-development-16458)

7. [Definições para OOP por Alan Kay](http://wiki.c2.com/?AlanKaysDefinitionOfObjectOriented)

### Autor
Cleverton Fernandes Guimarães, clevertonfernandesguimaraes@gmail.com
