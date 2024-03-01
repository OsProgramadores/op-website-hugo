+++
title = "Conceitos essenciais para começar"
date = "2024-02-29T18:50:00-03:00"
tags = ["programação", "desafios", "github", "programação", "tutorial"]
categories = ["artigos", "desafios"]
banner = "img/banners/Le_Penseur_in_the_Jardin_du_Musée_Rodin,_Paris_14_June_2015.jpg"
+++
## Introdução

Por vezes, algumas pessoas acabam encontrando dificuldades para iniciar na programação e encontram-se um pouco perdidas em meio à quantidade de recursos e opções disponíveis.
Aqui nos programadores, buscamos ajudar a todos da melhor forma possível. 

Sem sombra de dúvida, uma das perguntas mais comuns [no nosso grupo no telegram](https://t.me/osprogramadores) de quem está querendo iniciar é: 

> "**Por onde eu começo?**"

Comumente, direcionamos as pessoas aos artigos introdutórios como [por onde começar](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/) e na sequência, também encaminhamo-nas para o [desafio 01](https://osprogramadores.com/desafios/d01/).

Estes artigos são ótimos para **quem já tem uma boa base**! Porém, **e se esse não for o seu caso**? 

E se você tiver dificuldades em compreender o que é um sistema operacional?
O que é um computador? Como navegar entre diretórios? Não souber o que é um terminal?

> **Será que eu realmente estou pronto para começar sem entender sobre essas coisas**?
> 
> São muitos termos que não conheço! E agora, o que fazer?

Pode parecer muito complexo a princípio, mas você está pronto para poder começar **mesmo que não saiba nada do que foi mencionado até agora!**

Todos nascemos sem saber nada e aos poucos aprendemos sobre o mundo à nossa volta.
A vontade e persistência de entender as coisas, ter a iniciativa de ir atrás do conhecimento, é muito mais importante neste momento inicial! Serão estas características que lhe ajudarão
a prosseguir nessa jornada, **a curto e longo prazo**; porque a vida do programador é fundamentada por aprendizado constante, autoaprimoramento contínuo e compartilhar suas experiências!

O seu caminho inicial pode parecer ser um pouco mais longo, pois terá que aprender algumas coisas a mais. Isto significa que, uma vez que tenha compreendido estes conceitos fundamentais,
**poderá prosseguir com confiança e tranquilidade**! Estas coisas podem parecer bobas e simples a princípio, mas são diferenciais para você conseguir entender o que está fazendo e consolidar o seu progresso.

## O Programador e o Computador

Ao programar você está, a grosso modo, instruindo o computador a realizar uma sequência de instruções da forma que você deseja. Ou seja, o código que você gerar é responsável por orientar o computador
em todos os três escopos:
* O que fazer?
  * É o seu código que vai "falar" para o computador o que deve ser feito.
* Como fazer?
  * É o seu código que vai "mostrar" para o computador como algo deve ser feito.
* Por que fazer de modo X?
  * O computador não entende que algo pode ser feito de diversas maneiras, ele não tem essa capacidade. A capacidade dele é de seguir as instruções fornecidas, então o porquê de algo ser feito de tal forma
é algo dependente do programador e não do computador.

Notavelmente, quanto maior o seu entendimento sobre o funcionamento de um computador, a tendência é que o seu código seja mais preciso e faça um melhor uso dos recursos dele! 

### O que é um computador?
Para entender o que é um computador, recomendamos a [leitura deste artigo do IFBA](http://www.ifba.edu.br/professores/antoniocarlos/index_arquivos/resumodearquiteturadecomputadores.pdf). O artigo é de nível
introdutório, com uma abordagem amigável e de fácil compreensão, mesmo para quem não é da área.

Uma vez que você tenha lido o artigo, deves ter uma melhor compreensão do que é um computador e como ele funciona. Esta parte, o como ele funciona, irá permitir a **você, futuro programador**, fazer melhor uso
dos recursos disponíveis de um computador.

### O Sistemas de Arquivo

A próxima etapa agora é compreender como os dados são armazenados no disco de um computador. Uma boa forma de fazer isso é através [deste artigo do lucas linux](https://www.lucaslinux.com/post/sistema-de-arquivos-linux).
> Mas o meu computador não é linux, é Windows.

Qualquer sistema operacional pode ser usado, quer seja [Windows](https://www.microsoft.com/en-ca/windows), [Linux](https://en.wikipedia.org/wiki/Linux) ou [macOS](https://en.wikipedia.org/wiki/MacOS).
Não se preocupe, o conhecimento que o artigo lhe fornecerá será útil para todos os sistemas! De bônus, ao final do artigo, ele ainda vai lhe ensinar um comando para exibir os arquivos e diretórios abaixo do diretório atual.

Está curioso para saber qual é o comando? Dá uma lida no artigo e depois continue daqui!

> Tá, li o artigo e agora eu sei como os sitemas de arquivos funcionam. E agora?

### O Terminal

Agora seria bom você consolidar o entendimento sobre o que é um [terminal](https://pt.wikipedia.org/wiki/Terminal_(inform%C3%A1tica)). Em nosso [artigo introdutório](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/) e,
principalmente, no [desafio 01](https://osprogramadores.com/desafios/d01/), você precisará usar um terminal para inserir os comandos necessários.

Os artigos são bem completos e possuem os comandos necessários para quase todas as operações, porém há algumas coisas que são importantes que ficaram de fora deles, como, por exemplo, a navegação entre diretórios.

Uma habilidade fundamental para um programador é **saber como navegar entre diretórios no terminal**. Isto é necessário para que você consiga navegar entre as diferentes pastas dos projetos sem dificuldades, otimizando
o tempo que você gasta para alterar os ambientes e trocar de projeto. Usuários mais avançados, podem combinar diversos comandos de terminal para localizar informações dentro do código de forma tão rápida que pode
parecer assustador a princípio! 

> Os comandos para terminal são um mundo a parte, e apesar de parecerem desafiadores, eles estão ali para lhe auxiliar! Não tenha medo deles, eles são ferramentas feitas para otimizar o seu fluxo de trabalho.
> 
> [Temos um vídeo introdutório ao terminal do linux disponível no YouTube.](https://www.youtube.com/watch?v=CFWttwWZSAQ)

#### Navegação Básica no terminal

Para uma navegação básica no terminal você vai precisar utilizar os comandos [LS](https://pt.wikipedia.org/wiki/Ls) e [CD](https://pt.wikipedia.org/wiki/Cd_(comando)).

O commando `ls` listará os arquivos e os diretórios disponíveis dentro do seu diretório atual.
Ele é útil para você saber o que está disponível para você acessar. 

Um exemplo de uso do comando seria na seguinte situação:

> Estou no meu diretório de documentos e gostaria de saber quais os arquivos disponíveis ali.

Você pode simplesmente digitar o comando e então ele deverá exibir todos os arquivos e diretórios que estão disponíveis. Vamos imaginar que você quer entrar no diretório `codigo`, que fica dentro do seu diretório `documentos`.

Ao digitar `ls` dentro do diretório `documentos` ele deverá exibir o diretório `codigo`, além dos arquivos que por ventura você tenha ali.

O comando `cd` é utilizado para você alterar o diretório atual. O exemplo abaixo alteraria o diretório atual de `documentos` para `codigo`. 
```shell
cd codigo
```

Você pode utilizar o comando `cd` de duas formas:
1. Passando como parâmetro um diretório disponível dentro do diretório atual, como no exemplo anterior; 
2. Passando um caminho completo, desde o diretório raiz até o caminho desejado.

Exemplo da outra forma de usar o comando `cd`:
```shell
cd /home/username/Documentos/
```

Neste exemplo, o seu diretório de trabalho após utilizar o comando deverá ser `Documentos`
> Para um caso de uso real, você pode substituir a palavra _username_ pelo seu nome de usuário.

#### Criando Novos Diretórios

Para criar novos diretórios dentro do terminal você pode usar o comando `MKDIR`. Este comando funciona de maneira análoga ao comando `cd`, 
no sentido que você pode passar tanto um caminho completo, como não.

Vamos seguir com nosso exemplo e criar um diretório chamado `codigo` dentro do diretório `Documentos`. Para fazer isso você deve realizar as seguintes verificações:
1. Conferir se o seu diretório atual é o diretório `Documentos`. 
   1. Você aprendeu a navegar entre os diretórios na sessão anterior.
2. Caso não seja, navegar até o diretório `Documentos`.
3. Uma vez que esteja no diretório adequado, executar o comando a seguir para criar um diretório chamado `codigo` dentro do diretório atual.

```shell
mkdir codigo
```

Neste exemplo, um novo diretório chamado `codigo` deve ter sido criado dentro do seu diretório atual.
Você pode confirmar a criação do diretório utilizando o comando `ls` e conferindo se o diretório `código` será exibido no resultado do comando.

#### Removendo Arquivos e Diretórios

Eventualmente você pode precisar remover arquivos e/ou diretórios. Nessas ocasiões você vai precisar utilizar os comandos [RM](https://pt.wikipedia.org/wiki/Rm_(Unix)) e [RMDIR](https://pt.wikipedia.org/wiki/Rmdir).

O comando `RM` é mais focado para arquivos, mas também pode ser utilizado para diretórios se você adicionar o parâmetro apropriado para isso. Já o comando `RMDIR` é **exclusivo** para diretórios.

Dando continuidade ao nosso exemplo, vamos assumir que você possui um arquivo chamado `apagar.txt` em seu diretório atual. Para confirmar a existência deste arquivo, utilize o comando `ls` e confira se ele
aparecerá na saída do comando. 

> Caso ele não seja exibido, crie um arquivo com o nome `apagar.txt` e o salve em seu diretório atual. Após criar o arquivo, execute novamente o comando `ls` para confirmar a existência do arquivo. 

Uma vez que você confirmou a existência do arquivo, você pode executar o seguinte comando para removê-lo:

```shell
rm apagar.txt
```

> Pode ser que apareça uma tela de confirmação, perguntando se você realmente deseja remover o arquivo.
> Confirme com a opção positiva para remover o arquivo. Caso você negue, o arquivo não será removido!

Após executar este comando, novamente rode o comando `ls` para confirmar a presença do arquivo. Como você acabou de remover o arquivo, é de se esperar que ele **não apareça** no resultado.

Para testar o comando `rmdir` primeiro criaremos um novo diretório que logo mais será removido. Para fazer isso, execute o seguinte comando:
```shell
mkdir me-apague
```
Após rodar este comando, rode o comando `ls` para confirmar a existência do diretório `me-apague`. (Conseguiu perceber um padrão no uso do `ls` para confirmar as operações?)

Uma vez que você tenha confirmado a existência do diretório, **como nós acabamos de o criar ele deve estar vazio**. Assim sendo, podemos remover este diretório utilizando o seguinte comando:

```shell
rmdir me-apague
```
> Pode ser que apareça uma tela de confirmação, perguntando se você realmente deseja remover o diretório, similar ao que aconteceu com o comando `rm`. 

Novamente, valide o resultado da sua operação executando o comando `ls` para confirmar que o diretório `me-apague` não existe mais.

> Okay, entendi. Agora, e se eu tiver um diretório com arquivos dentros, o que eu faço?

Nesta ocasião, você poderá combinar estes comandos que aprendeu neste artigo para:
* alterar o diretório atual para o diretório que deseja remover os arquivos
* remover os arquivos que deseja
* alterar o diretório atual para exibir o diretório que deseja remover
* remover o diretório

## Conclusão

Esperamos que agora você esteja melhor preparado para iniciar a sua jornada no mundo da programação! 

Relembrando, neste artigo, abordamos os seguintes tópicos:

* O que é a **Programação** (a grosso modo)
* O que é um **Computador**
* O que é um **Sistema de Arquivos**
* O que é um **Terminal**
* Alguns **commandos básicos** para navegação no terminal

> Acreditamos que após esta leitura, você está mais do que apto para dar prosseguimento à sua jornada!

**Para dar continuidade em seus estudos agora vá para o artigo [por onde começar](https://osprogramadores.com/blog/2019/03/12/por-onde-comecar/).** 

Também recomendamos que você faça os [desafios do grupo](https://osprogramadores.com/desafios/) após ter lido os artigos.

Caso tenha dificuldades com os exercícios, ou dúvidas, faça perguntas no [grupo no telegram](https://t.me/osprogramadores).


