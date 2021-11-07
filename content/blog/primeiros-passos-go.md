+++
date = "2020-12-23T00:11:00-03:00"
title = "Primeiros passos em Go"
categories = ["linguagens"]
tags = ["programação","go", "golang", "história go", "projeto go", "install go"]
banner = "img/banners/go_banner1.png"
+++

Olá sou o Jefferson Otoni Lima conhecido como [jeffotoni](https://github.com/jeffotoni)
Vou fazer um pequeno overview de uma das linguagens que mais cresce no mundo.

Objetivo deste post é simplesmente apresentar o porque a **linguagem Go**
surgiu e apresentar a sua sintaxe e demonstrar algumas áreas onde Go 
é mais aplicado. Para todos que gostariam de aumentar ainda mais seu arsenal para 
desenvolvimento web este post irá ajuda-los a esclarecer alguns pontos importantes 
quando estamos iniciando e aprendendo uma nova linguagem de programação.
Vamos conhecer mais deste universo fantástico que é Go.

Para você que está iniciando em programação ou que já conhecça de programação este
post irá ser útil para você entender um pouco mais quando o assunto é Go.

![banner Go](/img/conteudos-de-artigos/go_banner1.png)

## Primeiros passos em Go

Antes de iniciarmos, é importante salientar que linguagens de programação são ferramentas e como toda boa ferramenta temos que saber em quais momentos usá-las. Existem cenários e problemas que só são resolvidos com linguagens específicas e existem outros universos de problemas que temos centenas ou milhares de linguagens que resolvem de alguma forma o mesmo problema. Então como um bom profissional quanto é importante entender qual(quais) ferramenta(s) de adequar a quais problemas.

A linguagem Go, em seu universo de possibilidades, é uma linguagem de uso geral mas não gosto muito deste termo pois fica parecendo que a linguagem é tipo uma bala de prata e resolve todos os problemas e isso não é verdade. 

Go nasceu por um propósito que é resolver problemas do universo web e aproveitar a nova tecnologia de multicores em servidores para este propósito.

O primeiro passo é conhecer um pouco melhor o site oficial da linguagem [site oficial go](https://golang.org).

## História do projeto Go

Robert Griesemer, Rob Pike e Ken Thompson começaram a esboçar as metas para uma nova linguagem no quadro branco em 21 de setembro de 2007. Em poucos dias, as metas se estabeleceram e foi criado um plano de execução. Em janeiro de 2008, Ken começou a trabalhar em um compilador para explorar ideias; ele gerou código C como sua saída. Em meados do ano, a linguagem se tornou um projeto em tempo integral e se estabeleceu o suficiente para ser um compilador de produção. 

Em maio de 2008, Ian Taylor começou de forma independente em um front-end GCC para Go usando as especificações preliminares. Russ Cox entrou no projeto no final de 2008 e ajudou a mover a linguagem e as bibliotecas do protótipo adiante.

Go tornou-se um projeto de código aberto público em 10 de novembro de 2009. Inúmeras pessoas da comunidade contribuíram com ideias, discussões e códigos.

"_Existem agora milhões de programadores Go “Gophers” em todo o mundo, e há mais a cada dia. O sucesso de Go excedeu em muito as expectativas._"

![Engenheiros Go](/img/conteudos-de-artigos/engenheiros_go.png)


## Por que criaram uma nova linguagem?

"Go nasceu da frustração com as linguagens e ambientes existentes para o trabalho que estávamos fazendo no Google. A programação havia se tornado muito difícil e a escolha das linguagens era parcialmente culpada. Era preciso escolher a compilação eficiente e a execução eficiente ou facilidade de programação; todos os três não estavam disponíveis na mesma linguagem principal. Os programadores poderiam escolher facilidade em vez de segurança e eficiência, mudando para linguagens tipadas dinamicamente, como Python e JavaScript, em vez de C++ ou, em menor grau, Java."

"_Go tratou dessas questões tentando combinar a facilidade de programação de uma linguagem interpretada e dinâmica com a eficiência e segurança de uma linguagem compilada estaticamente. Também pretendia ser moderno, com suporte para computação em rede e multicore. Finalmente, trabalhar com Go pretende ser rápido : deve demorar no máximo alguns segundos para construir um grande executável em um único computador._"

Nasce uma nova linguagem Go, para atender às novas necessidades e resolver problemas aproveitando o máximo possível do poder computacional. 

E por que do grande sucesso de Go ?

A comunidade Go não tem dúvidas que é devido ao seu mascote e que hoje tem milhares de gophers espalhados pelo mundo.😁

![Mascote Go](/img/conteudos-de-artigos/mascote_go.jpg)


## Site Oficial

Para iniciarmos em Go vamos dar alguns passos para trás, vamos começar toda nossa trajetória conhecendo o [site oficial da linguagem](https://golang.org). Nesta página encontramos informações sobre Go e muito mais. 

Nesta página temos os:

  - [docs](https://golang.org/doc)
  - [packages](https://golang.org/pkg)
  - [blog](https://blog.golang.org)
  - [play go](play.golang.org)
  - [efetive go](https://golang.org/doc/effective_go.html)
  - [especificações Go](https://golang.org/ref/spec)
  - [download](https://golang.org/doc/install) 
  - [tour em Go](https://tour.golang.org/welcome/1) 

E não para por ai, tem muitas possibilidades no site, se aprendermos o site todo restor torna-se fácil. 

O site oficial aparentemente parece pequeno mas ele é muito completo e grande. Então temos praticamente tudo que precisamos saber de Go para iniciarmos nosso aprendizado nesta linguagem que é um fenômeno.

### Onde Inicio no Site

Vou criar uma linha de raciocínio para que possamos entender Go de forma mais prática possível.
Antes de instalar Go, ou rodar Go pelo play, vamos dar uma passada em algumas partes do doc para que possamos entender a história do Go e por que nasceu uma nova linguagem neste universo de milhares de linguagens de programação.

Para acessar o site oficial basta clicar em: [site oficial go](https://golang.org).

![Site Oficial](/img/conteudos-de-artigos/site_oficial_go.png)

## Effective Go

Este link é todo material que se precisa ler antes de tudo.
[Effective Go](https://golang.org/doc/effective_go.html#introduction) ❤️ neste site está todo material que precisa para ter uma boa noção da linguagem Go: 

   - [Estrutura de controle](https://golang.org/doc/effective_go.html#control-structures)  
   - [Funções](https://golang.org/doc/effective_go.html#functions)
   - [Programação concorrente usando Goroutines](https://golang.org/doc/effective_go.html#concurrency)
   - [Interfaces e métodos](https://golang.org/doc/effective_go.html#interface-names)
   - [Map](https://golang.org/doc/effective_go.html#maps)
   - [Test nativa da plataforma](https://golang.org/doc/tutorial/add-a-test)
   - [Profiling nativo da plataforma](https://blog.golang.org/pprof)

Existe muito mais além destes pontos colocados no site oficial, vale muito a pena dedicar e investir um tempo na leitura pois tudo irá ficar bem mais fácil quando for trabalhar na prática com a linguagem Go.
 
## Um tour de Go

Este link seria o segundo mais importante na minha hierarquia para aprendermos Go.
[Tour Go](https://tour.golang.org/welcome/1) ❤️ neste site você irá conseguir testar e dar uma passada em algumas das funcionalidades da linguagem Go 😱 isto mesmo tudo pelo browser sem precisar instalar nadinha na máquina.

Muito show de bola esta possibilidade, estudar direto no site do go, executar os códigos e visualizar o comportamento e como funciona a linguagem.

## Playground Go

No [Play Go](https://play.golang.org/p/MAohLsrz7JQ) você irá conseguir executar Go 😱 sem instalar nada nada. Todo o código Go é escrito e executado no navegador (browser) e simplifica muito o aprendizado.

Isto é Muito legal, escrever o código no browser, excelente ferramenta para criarmos nossos pequenos código para tirarmos dúvidas, estudar etc.


## Perguntas mais Frequentes

Este link seria o nosso terceiro passo e acredito que agora você tnha diversas dúvidas, geradas pelo Effective Go e pelo Tour que fez.

[Faq](https://golang.org/doc/faq) ❤️ é aqui que você irá tirar algumas dúvidas. Esta página é essencial para organizar suas ideias e entender realmente um pouco mais sobre Go🥰. O tempo gasto lendo esta página irá com certeza lhe ajudar a economizar horas de trabalho.

Selecionei alguma delas:
	
   - [O Google está usando Go internamente?](https://golang.org/doc/faq#internal_usage)
   - [Que outras empresas usam Go?](https://golang.org/doc/faq#external_usage)
   - [Os programas Go se vinculam a programas C / C ++?](https://golang.org/doc/faq#Do_Go_programs_link_with_Cpp_programs)
   - [Go é uma linguagem orientada a objetos?](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)
   - [Por que não há aritmética de ponteiro?](https://golang.org/doc/faq#no_pointer_arithmetic)
   - [Por que fazer a coleta de lixo? Não vai ser muito caro?](https://golang.org/doc/faq#garbage_collection)


## Pacotes nativos Go

Este link contem documentação para ajudar você a entender como são organizadas funções, libs, pkgs que podem ser usados na linguagem Go.
[Pkgs](https://golang.org/pkg/) ❤️ escolhi aqui o [pkg string](https://golang.org/pkg/strings) que inclui exemplos que podem ser executados diretamente no navegador facilitar o aprendizado e entendimentos🥰.

Existe centenas de pacotes nativos da linguagem, quando for construir sua api rEST por exemplo já tem a biblioteca nativa da linguagem, quando for escrever seus tests você irá utilizar a biblioteca nativa da linguagem, quando for usar criptografia você irá usar biblioteca nativa da linguagem.

Isto significa que a linguagem é robusta quando o assunto são recursos nativos para que possa desenvolver seus projetos.

## Editores e IDE

Este link é sobre as IDEs e Editores que poderá utilizar quando estiver escrevendo programas em Go. A linguagem Go não precisa de muito para que possamos editar nossos códigos, mas temos alguns plugins interessantes para facilitar ainda mais nosso dia a dia
[Editores e IDE](https://golang.org/doc/editors.html) ❤️ Neste link não irá ter o Sublime e o nvim que são dois editores que uso por aqui no dia a dia, o sublime pela pesquisa do survey somente 7,7% utilizam sublime e vim 14% o nvim é um plugin que irá usar no vim. O queridinho e adotado pelo equipe Go é o VsCode com 41% de adesão.

## Pontos Fortes da linguagem Go

Antes de fazermos nosso famigerado **"Hello World"** vamos mostrar alguns pontos fortes da Linguagem, vamos fazer uma prévia dos pilares e características da linguagem Go.

### 3 Pilares

Temos alguns pilares bem definidos em Go, isto ajuda a clarear ainda mais seu horizonte quanto o assunto é Go.

   - Existem idiomas que são um pouco mais rápidos que o Go, mas certamente não são tão simples quanto o Go.
   
   - Existem linguagens que tornam a concorrência seu objetivo mais elevado, mas não são tão legíveis nem produtivas.
   
   - Desempenho e simultaneidade são atributos importantes, mas não tão importantes quanto:

Podemos dizer que Go possui três pilares para atender estas necessidades:

   - Simplicidade

   - Legibilidade

   - Produtividade

### Características principais

Em Go temos algumas características marcantes da linguagem que a tornam ainda mais poderosa para desenvolvimento de aplicações web para servidores.

   - Somente 25 keywords
   - Curva de aprendizado baixa
   - Compilada estaticamente
   - Multiplataforma agora com suporte a RISC-V
   - Paradigma Concorrente
   - Tipagem Estática
   - Retrocompatibilidade
   - GC (Garbage collector)

### Alguns tipos de aplicações implementadas em Go
	
   - [Web backend](https://github.com/nouney/awesome-go#web-frameworks) (com diversos frameworks disponíveis)
   - [Web Assembly](https://github.com/vugu/vugu) (um dos frameworks vugu)
   - Microservices Frameworks
	- [Go Micro](https://micro.mu/)
	- [Go Kit](https://gokit.io/)
	- [Gizmo](https://github.com/NYTimes/gizmo)
	- [Kite](https://github.com/koding/kite)
   - Fragments services (Termo citado pelo @jeffotoni em um grupo de discussão de microservices)
   - Lambdas [FaaS example](https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda)
   - Client Server
   - Aplicações em terminal [utilizando a lib tview](https://github.com/rivo/tview) 
   - IoT [alguns frameworks](https://github.com/nouney/awesome-go#iot-internet-of-things)
   - Boots [alguns aqui](https://github.com/nouney/awesome-go#bot-building)
   - Aplicações Client que usam tecnologia Web
      - Desktop:
	- [Qt+QML](https://github.com/therecipe/qt)
	- [Lib Nativa Win](https://github.com/lxn/walk)
	- [Widgets Qt](https://therecipe.github.io/widgets_playground/)
	- [Qml](https://doc.qt.io/qt-5/qtqml-index.html)
   - Aplicações de Rede
   - Aplicações para protocolos
   - Aplicações rEST
   - Aplicações SOAP
   - Aplicações GraphQL
   - Aplicações RCP
   - Aplicações TCP
   - Aplicações gRPC
   - Aplicações Websocket
   - GopherJS [ompiles Go to JavaScript](https://github.com/gopherjs)

## Instalando Go

![Docker Go](/img/conteudos-de-artigos/goinstall.png)

Install Go neste link terá o passo a passo de como instalar em diversos Sistemas Operacionais mas vou deixar aqui a instalação no Linux.

## Instalação no Linux

Se você tiver uma versão anterior do Go instalada, certifique-se de removê-la antes de instalar, vamos baixar a versão go1.15.6 mas pode entrar no site e pegar a última versão se preferir.

#### Confira o passo a passo 

1. Remova o Go caso esteja instalado

```bash
$ sudo rm -rf /usr/local/go
```
1. Baixe o arquivo [clicando aqui](https://golang.org/dl/go1.15.6.linux-386.tar.gz) 

2. Extraia-o em **/usr/local** com o seguinte comando

```bash
$ sudo tar -C /usr/local -xzf go1.15.6.linux-386.tar.gz
```

3. Adicione **/usr/local/go/bin** à PATH variável de ambiente.

```bash
$ export PATH=$PATH:/usr/local/go/bin
```

4. Para deixar o que fizemos acima global, ou seja toda vez que for entrar em seu terminal "go" seja executado você pode fazer isso adicionando a seguinte linha que mostrei acima ao seu $HOME/.profile ou /etc/profile (para uma instalação em todo o sistema) ou em seu **bash** ou o que estiver usando em seu ambiente Linux, depois basta fazer:

```bash
$ source ~/.profile
```

5. Agora vamos testar se tudo deu certo 

```bash
$ go version
```

Saída:
```bash
go version go1.15.6 linux/386
```

## Hello World

Agora sim chegou a hora de nossos primeiros **"hello world"** utilizando Go.
Não tem segredo, podemos utilizar de qualquer editor que vimos acima e irmos para cima. Em Go não precisamos compilar quando estamos escrevendo programas em Go podemos simplesmente executar nossos programas sem precisar transformá-los em binários ou seja compilar.

### go run

Abaixo está nosso primeiro hello world, e basta salvar o arquivo com a extensão "_.go_", coloque o nome de "_main.go_" depois basta executar ele com o comando "_go run main.go_", em nosso exemplo chamei o arquivo de "main.go"

```go
package main

func main() {
  
  println("Meu primeiro Hello World")

}

```

Este nosso hello world é a menor estrutura que podemos montar para da saída no terminal, sem usar pacote.

Este outro exemplo vamos utilizar um pacote nativo Go.

```go
package main

import "fmt"

func main() {
  
  fmt.Println("Meu primeiro Hello World")

}
```

Para executar o nosso programa, pode usar o "_go run_" para não precisarmos compilar enquanto estamos escrevendo programas em Go nosso projeto de uma conferida como é fácil.
```bash
$ go run main.go
```

Saída
```bash
Meu primeiro Hello World
```

Para ficar um pouquinho mais claro quando usarmos a função println temos que saber que ela é uma função embutida (no tempo de execução) que pode eventualmente ser removida, enquanto o "_fmt_" pacote está na biblioteca padrão, que irá persistir.

### go build

Agora vamos transformar em **binário** vamos compilar e para isto vamos utilizar o "go build" ou "GOARCH=386 GOOS=linux go build" estamos agora informando a arquitetura e sistema operacional para o qual desejamos compilar nosso programa Go.

```bash
$ GOARCH=386 GOOS=linux go build -o myfirstprogram main.go
```

Este comando irá fazer o **build** e será gerado um "executável", também conhecido como binário, para ser executado em seu sistema operacional. Muito fácil não é? 😍 Eu diria muito **lindo**. Com este binário você irá conseguir executar o programa em sua máquina ou em qualquer servidor que tenha a mesma **arquitetura e sistema operacional**. 

O mais legal é que é gerado um binário estático que não contém dependências para instalar em seu servidor. Legal não é? Para sabermos se o arquivo binário é dinâmico ou estático basta rodarmos o seguinte commando:

```bash
$ ldd myfirstprogram
```

Saída
```bash
não é um executável dinâmico
```

Agora vamos executar nosso binário myfirstprogram
```bash
$ ./myfirstprogram
```

Saída
```bash
Meu primeiro Hello World
```

Se aparecer esta saída então o executável gerado é estático sem dependências.
_Assumindo que o sistema está configurado para uso do idioma Português._

É importante frisar sempre este detalhe: **"Não instalamos nada de Go em nossos servidores Web"** é exatamente o que leu, só precisa do binário. ❤️ Foi devido a isto que Go ficou conhecido por ser a Linguagem dos Containers. Diz que isto não é fantástico ? 😍 Para rodar uma aplicação Web em seu servidor, basta enviar o binário para ele. Uau, isto mesmo, somente do binário e sem precisar instalar mais nada referente ao seu projeto, nadinha.

## Conclusão

Este post é um simples resumo para você que gostaria de aumentar seu arsenal para programação web. 

Espero que tenham gostado, e fique a vontade em encontrar em contato para sugestões ou melhorias no post.


Todo este assunto foi discutido ao vivo em uma live no youtube:

   - [youtube live](https://youtube.com/user/jeffotoni)
   - [PDF da apresentação](https://speakerdeck.com/jeffotoni/primeiros-passos-em-go)
